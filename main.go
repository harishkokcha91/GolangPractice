package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "sammy:Password@123@tcp(127.0.0.1:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// perform a db.Query insert
	res, err := db.Query("Select * from itemDetails")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	var item Item
	var itemList []Item

	for res.Next() {
		fmt.Println(res)
		res.Scan(&item.ItemId, &item.ItemName, &item.ItemPrice)
		row, err := res.Columns()

		if err != nil {
			fmt.Printf(err.Error())
		}

		fmt.Printf("%v\n", item)
		fmt.Printf("%v\n", row)
		itemList = append(itemList, item)
	}
	// be careful deferring Queries if you are using transactions
	defer res.Close()

}

type Item struct {
	ItemId    string `json:"item_id"`
	ItemName  string `json:"item_name"`
	ItemPrice string `json:"item_price"`
}
