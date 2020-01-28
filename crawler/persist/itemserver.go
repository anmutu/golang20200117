/*
  author='du'
  date='2020/1/27 18:32'
*/
package persist

import (
	"context"
	"github.com/olivere/elastic/v7"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("ItemSaver得到#%d,%v", itemCount, item)
			_, err := save(item)
			if err != nil {
				log.Printf("ItemSaver:存储%v时出错，错误为%v", item, err)
			}
		}
	}()
	return out
}

func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return "", err
	}
	//这里的index就是增加， 不是add，也不是create。
	resp, err := client.Index().
		Index("du_profile").
		Type("cnblogs").
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}
