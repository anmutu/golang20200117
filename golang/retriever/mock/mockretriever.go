/*
  author='du'
  date='2020/1/21 16:28'
*/
package mock

type Retriever struct {
	Contents string
}

func (r Retriever) Get(url string) string {
	return r.Contents
}
