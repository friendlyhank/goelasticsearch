package common

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"os"
)


type Elasticclient struct{
	*elastic.Client
}

//NewClient -
func NewClient(host string)(client *elastic.Client,err error){
	errorlog := log.New(os.Stdout,"APP",log.LstdFlags)
	client,err = elastic.NewClient(elastic.SetErrorLog(errorlog),elastic.SetURL(host))
	if err != nil{
		panic(err)
	}
	return
}

//ElasticGetVersion -
func (ec *Elasticclient)ElasticGetVersion(host string)string{
	esversion,err := ec.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)
	return esversion
}

//CreateElasticTable -Create an index
func (ec *Elasticclient)CreateElasticTable(tableName string)error{
	//Create an index
	var err error
	_,err = ec.CreateIndex(tableName).Do(context.Background())
	return err
}

//InsertElasticData -Add a document to the index
func (ec *Elasticclient)InsertElasticData(tableName string,tpe string,id string,data interface{})error{
	//Insert
	var err error
	_,err=ec.Index().
		Index(tableName).
		Type(tpe).
		Id(id).
		BodyJson(data).
		Refresh("wait_for").
		Do(context.Background())
	return err
}

//UpdateElasticData - update a document to the index
func (ec *Elasticclient)UpdateElasticData(tableName string,tpe string,id string,data interface{})error{
	//Update
	var err error
	_,err = ec.Update().
		Index(tableName).
		Type(tpe).
		Id(id).
		Doc(data).   //map[string]interface{}{"age": 88} or struct
		Do(context.Background())
	return err
}


//DeleteElasticData -delete on data
func (ec *Elasticclient)DeleteElasticData(tableName string,tpe string,id string)error{
	var err error
	_,err = ec.Delete().
		Index(tableName).
		Type(tpe).
		Id(id).
		Do(context.Background())
	return err
}

//DropTable - DropTable
func (ec *Elasticclient)DropTable(tableName string)error{
	var err error
	_,err = ec.DeleteIndex(tableName).Do(context.Background())
	return err
}

//SearchElasticQuery - Search Search with a term query
func (ec *Elasticclient)SearchElasticQuery(tablename string,key string,value string)(searchResult *elastic.SearchResult,err error){
	termQuery := elastic.NewTermQuery(key,value)
	searchResult,err = ec.Search().
		Index(tablename). //search in index "tweets"
		Query(termQuery).
		Sort(key+".keyword",true). //sort by "user" field,ascending
		From(0).Size(10).    //take docuemnt 0-9
		Pretty(true).   //output json
		Do(context.Background())
	return
}




// searchResult is of type SearchResult and returns hits, suggestions,
// and all kinds of other information from Elasticsearch.

/**
fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

**/

// Each is a convenience function that iterates over hits in a search result.
// It makes sure you don't need to check for nil values in the response.
// However, it ignores errors in serialization. If you want full control
// over iterating the hits, see below.

/**
var ttyp Tweet
for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
if t, ok := item.(Tweet); ok {
fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
}
}

**/

// TotalHits is another convenience function that works even when something goes wrong.
/**
fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())
**/





