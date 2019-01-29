package main

var host = "http://192.168.66.130:9200"

type Tweet struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

func main(){
	/**

	client,err := common.NewClient(host)

	var elasticclient = common.Elasticclient{client}

	if nil != err{
		panic(err)
	}

	//Create an index
	_,err = client.CreateIndex("tweets").Do(context.Background())

	if nil != err{
		panic(err)
	}

	//Add a docuemnt to the index
	tweet := Tweet{User:"oliverre",Message:"Take Five"}
	err = elasticclient.InsertElasticData("tweets","doc","1",tweet)

	if nil != err{
		panic(err)
	}

	//修改
	tweet := Tweet{User:"oliverre",Message:"hello world"}
	err = elasticclient.UpdateElasticData("tweets","doc","1",tweet)

	if nil != err{
		panic(err)
	}

	//删除
	err = elasticclient.DeleteElasticData("tweets","doc","1")

	if nil != err{
		panic(err)
	}

	searchResult,err := elasticclient.SearchElasticQuery("tweets","user","oliverre")

	if nil != err{
		panic(err)
	}

	fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())

	var ttyp Tweet
	for _,item := range searchResult.Each(reflect.TypeOf(ttyp)){
		if t, ok := item.(Tweet); ok {
			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
		}
	}

	**/
}