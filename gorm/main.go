/*
  author='du'
  date='2020/2/6 21:15'
*/
package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golangtestdb?charset=utf8")
	db.SingularTable(true)
	defer db.Close()
	if err != nil {
		log.Printf("连接db出错，error是:%s", err)
	}

	db.AutoMigrate(&User{})

	//创建
	db.Create(&User{Name: "du", Age: 18})

	var user User
	db.First(&user)
	fmt.Printf("查出的结果为%v\n", user)

	db.Delete(&user, 1)

}

type User struct {
	gorm.Model
	Name string
	Age  int
}
