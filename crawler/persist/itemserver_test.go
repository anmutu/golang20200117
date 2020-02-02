/*
  author='du'
  date='2020/2/2 15:42'
*/
package persist

import (
	"golang20200117/crawler/model"
	"testing"
)

func TestSave(t *testing.T) {
	expected := model.Blog{
		Name: "dudu",
	}
	Save(expected)
}
