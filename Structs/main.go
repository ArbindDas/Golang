package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City    string
	Phone   string
	HouseNo int
}

type Product struct {
	ProductId    int
	ProductName  string
	ProductPrice string
	Address      Address
}

func main() {

	ProductDetails := Product{

		ProductId:    90,
		ProductName:  "jeans",
		ProductPrice: "1289",

		Address: Address{
			City:    "jeetpur",
			Phone:   "+977 9821161214",
			HouseNo: 1001,
		},
	}

	data, _ := json.MarshalIndent(ProductDetails, "", " ")
	fmt.Println(string(data))
	// fmt.Println(ProductDetails)

	// fmt.Println("jai shree ram")
}
