package main

import (
	"4-order-api/internal/config"
	"4-order-api/internal/database"
	"4-order-api/internal/product"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&product.Product{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("migration completed")
}