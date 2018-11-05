package main

import (
	"database/sql"
	"fmt"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mytest?charset=utf8")
	checkErr(err)
	//插入数据

	fmt.Println("insert data info into table, userinfo")

	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")

	checkErr(err)

	res, err := stmt.Exec("Test", "people", "2017-10-27")
	fmt.Print(res)

}
