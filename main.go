package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

	"bytes"
	"encoding/json"
	"log"
	"net/http"
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
	Id				int		`json:"id"`
	OrderName		string	`json:"order_name"`
	NumberOfOrders	int		`json:"number_of_orders"`
}

func main() {
	handler1 := func (w http.ResponseWriter, r *http.Request)  {
		var  buf bytes.Buffer
		enc := json.NewEncoder(&buf)

		orders := []Order{}

		rows, err := DB.Query("SELECT * FROM order_list ORDER BY id")
		if err != nil {
			return
		}
		for rows.Next() {
			m := Order{}
			rows.Scan(&m.Id, &m.OrderName, &m.NumberOfOrders)
			fmt.Println(m)
			orders = append(orders, m)
		}

		if err2 := enc.Encode(&orders); err2 != nil {
			log.Fatal(err)
		}

		w.Header().Set("Access-Control-Allow-Origin","http://localhost:3000")

		_, err3 := fmt.Fprint(w, buf.String())
		if err3 != nil {
			return
		}
	}

	http.HandleFunc("/order", handler1)
	log.Fatal(http.ListenAndServe(":8000", nil))
}