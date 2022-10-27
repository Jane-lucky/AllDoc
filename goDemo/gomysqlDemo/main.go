package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

func main() {
	db, err = sql.Open("mysql", "root:123456@(172.17.0.2:3306)/testdb?&parseTime=true")
	checkErr(err)
	defer db.Close()
	err = db.Ping()
	checkErr(err)

	//路由

	//列表
	http.HandleFunc("/", listHandle)
	http.HandleFunc("/list", listHandle)

	http.HandleFunc("/add", addHandle)
	http.HandleFunc("/update", updateHandle)
	http.HandleFunc("/delete", deleteHandle)

	http.ListenAndServe(":8080", nil)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
