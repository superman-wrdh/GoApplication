package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GetDB(dbName string) *sql.DB {
	db, err := sql.Open("mysql", "root:hcissuperman@tcp(192.168.199.199:3306)/"+dbName+"?charset=utf8")
	checkErr(err)
	return db
}

func query() {
	s := `
		select id,no,staff_id,name from user
	`
	db := GetDB("cgs")
	rows, err := db.Query(s)

	checkErr(err)

	fmt.Println("all data")
	columns, _ := rows.Columns()
	fmt.Println(columns)

	for rows.Next() {

		var id string

		var no string

		var staff_id string

		var name string

		err = rows.Scan(&id, &no, &staff_id, &name)

		checkErr(err)

		fmt.Println(id, " ", no, " ", staff_id, " ", name)

	}

}
func main() {
	query()
}
