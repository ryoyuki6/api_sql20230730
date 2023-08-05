package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "user=root dbname=testdb password=root sslmode=disable")
	if err != nil {
		panic(err)
	}
}

type Order struct {
	Id				int
	OrderName		string
	NumberOfOrders	int
}

func main() {
	rows, err := DB.Query("SELECT * FROM order_list ORDER BY id")
	if err != nil {
		return
	}
	for rows.Next() {
		m := Order{}
		rows.Scan(&m.Id, &m.OrderName, &m.NumberOfOrders)
		fmt.Println(m)
	}
}