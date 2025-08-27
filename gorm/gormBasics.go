package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("There is some error in opening the database")
	}
	db.AutoMigrate(&Product{})
	db.Create(&Product{Code: "ayush", Price: 25})
	var product Product
	db.First(&product, "Code=?", "ayush") //db.First gets the first value that satisfies the condition
	fmt.Println(product)

	// update the price of the product
	db.Model(&product).Update("Price", 200) // it will update the price(one field) of the product to 200
	log.Println(product)

	// to update more than one field we use
	db.Model(&product).Updates(Product{Price: 540, Code: "hello"})
	fmt.Println(product)
	var product2 Product
	db.First(&product2, 1)
	fmt.Println(product2)
}
