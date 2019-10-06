package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	//func Open(driverName , dataSourceName string)(*DB , error)
	db, err := sql.Open("postgres", "postgres://bond:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping() //it checks if the connection is well established or not else throws an error
	if err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}
