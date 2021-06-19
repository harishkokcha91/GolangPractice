package main

import (
	"GolangPractice/dbutils"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	first := []string{"a", "b", "c", "d", "a", "b", "c", "d"}
	fmt.Println("first items ", first)
	getProductPrice(first)

}

func getProductPrice(list []string) {
	itemTotalQty := make(map[string]int)

	for i := 0; i < len(list); i++ {
		itemTotalQty[list[i]] = itemTotalQty[list[i]] + 1
	}
	totalSumOfItemPrice := 0

	for k, v := range itemTotalQty {
		// fmt.Print(k, v)
		totalSumOfItemPrice = totalSumOfItemPrice + fetchItemOffer(k, v)
		// fmt.Println(totalSumOfItemPrice)
	}

	fmt.Println("Total checkout price : ", calculateTatalPrice(totalSumOfItemPrice))

}

func calculateTatalPrice(totalSumOfItemPrice int) float32 {

	var itemOfferList []ItemOffer
	// perform a db.Query insert
	res1, err := dbutils.DbConn().Query("Select * from itemsOffer where item_total_price != '' and item_total_price <= ? order by item_total_price DESC limit 1", totalSumOfItemPrice)

	if err != nil {
		panic(err.Error())
	}
	for res1.Next() {
		var itemOfferw ItemOffer
		res1.Scan(&itemOfferw.ID, &itemOfferw.ItemName, &itemOfferw.ItemQty, &itemOfferw.ItemPrice, &itemOfferw.ItemTotalPrice, &itemOfferw.ItemDiscount)
		if err != nil {
			fmt.Printf(err.Error())
		}

		itemOfferList = append(itemOfferList, itemOfferw)
	}
	fmt.Println("len(res1) ", itemOfferList)
	defer res1.Close()
	var newPrice float32 = float32(totalSumOfItemPrice)
	// fmt.Println("len(itemOfferList) ", len(itemOfferList))
	if len(itemOfferList) > 0 {
		newPrice = float32(totalSumOfItemPrice) - float32(totalSumOfItemPrice*itemOfferList[0].ItemDiscount/100)
		// fmt.Println("len(itemOfferList) newPrice ", itemOfferList[0].ItemDiscount)
	}

	return newPrice
}

func fetchItemOffer(itemName string, itemqty int) int {
	price := 0
	var itemOffer ItemOffer
	var itemOfferList []ItemOffer
	// perform a db.Query insert
	res1, err := dbutils.DbConn().Query("Select * from itemsOffer where item_name =? and item_qty <=? order by item_qty DESC limit 1", itemName, itemqty)
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	for res1.Next() {
		res1.Scan(&itemOffer.ID, &itemOffer.ItemName, &itemOffer.ItemQty, &itemOffer.ItemPrice, &itemOffer.ItemTotalPrice, &itemOffer.ItemDiscount)
		if err != nil {
			fmt.Printf(err.Error())
		}
		itemOfferList = append(itemOfferList, itemOffer)
	}

	defer res1.Close()

	var item Item
	var itemList []Item
	res, err := dbutils.DbConn().Query("Select * from itemDetails where item_name =?", itemName)
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	for res.Next() {
		res.Scan(&item.ItemId, &item.ItemName, &item.ItemPrice)
		if err != nil {
			fmt.Printf(err.Error())
		}
		itemList = append(itemList, item)
	}

	defer res.Close()

	if len(itemOfferList) > 0 {
		// fmt.Println("len ", len(itemOfferList))
		for i := 0; i < len(itemOfferList); i++ {
			// fmt.Println("itemOfferList[i].ItemQty ", itemOfferList[i].ItemQty)
			if itemqty == itemOfferList[i].ItemQty {
				price += itemOfferList[i].ItemPrice
			} else if itemqty > itemOfferList[i].ItemQty {
				price += (itemqty%itemOfferList[i].ItemQty)*itemList[0].ItemPrice + (itemqty/itemOfferList[i].ItemQty)*itemOfferList[i].ItemPrice
			} else {
				price += itemqty * itemList[0].ItemPrice
			}
			// fmt.Println("price ", price)

		}
	} else {
		price += itemqty * itemList[0].ItemPrice
	}

	return price
}

type Item struct {
	ItemId    string `json:"item_id"`
	ItemName  string `json:"item_name"`
	ItemPrice int    `json:"item_price"`
}

type ItemOffer struct {
	ID             string `json:"id"`
	ItemName       string `json:"item_name"`
	ItemPrice      int    `json:"item_price"`
	ItemQty        int    `json:"item_qty"`
	ItemTotalPrice int    `json:"item_total_price"`
	ItemDiscount   int    `json:"item_discount"`
}
