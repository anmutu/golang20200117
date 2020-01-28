/*
  author='du'
  date='2020/1/27 18:32'
*/
package persist

import (
	"context"
	"github.com/olivere/elastic/v7"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("ItemSaver得到#%d,%v", itemCount, item)
			save(item)
		}
	}()
	return out
}

func save(item interface{}) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	//这里的index就是增加， 不是add，也不是create。
	resp, err := client.Index().
		Index("du_profile").
		Type("cnblogs").
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
