package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// mysql
func test2() {
	db, err := sql.Open("mysql", "root:tango9896@/ebean?charset=utf8")
	checkErr(err)

	//table := "customer"

	// insert
	stmt, err := db.Prepare("INSERT into customer SET name=?,start_date=?,comments=?,version=?")
	checkErr(err)

	res, err := stmt.Exec("Tom", "2017-09-12", "一些文字表存在感", 2)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	// update
	stmt, err = db.Prepare("update customer set name=? where id=?")
	checkErr(err)

	res, err = stmt.Exec("Tom-update", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// query
	rows, err := db.Query("SELECT id, name, start_date, comments, version FROM customer")
	checkErr(err)

	for rows.Next() {
		var id int64
		var name string
		var start_date string
		var comments string
		var version int64
		err = rows.Scan(&id, &name, &start_date, &comments, &version)
		checkErr(err)
		fmt.Println("id: ", id, ", name: ", name, ", start_date: ", start_date, ", comments: ", comments, ", version: ", version)
	}

	// delete
	/*
		stmt, err = db.Prepare("delete from ? where uid=?")
		checkErr(err)

		res, err = stmt.Exec(table, id)
		checkErr(err)

		affect, err = res.RowsAffected()
		checkErr(err)

		fmt.Println(affect)
	*/

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	test2()
}
