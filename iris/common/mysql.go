/*
  author='du'
  date='2020/2/6 21:03'
*/
package common

import "database/sql"

//创建mysql 连接
func NewMysqlConn() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:imooc@tcp(127.0.0.1:3306)/imooc?charset=utf8")
	return
}
