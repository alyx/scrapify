package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/alyx/scrapify"
)

func openFileAndReadProducts(file string) ([]scrapify.Product, error) {
	var products []scrapify.Product
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(&products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func containsProduct(products []scrapify.Product, p scrapify.Product) bool {
	for _, product := range products {
		if product.ID == p.ID {
			return true
		}
	}
	return false
}

func main() {
	var missingProducts []scrapify.Product
	if len(os.Args) != 3 {
		fmt.Println("Usage: productdiff <old.json> <new.json>")
		return
	}
	oldProducts, err := openFileAndReadProducts(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	newProducts, err := openFileAndReadProducts(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, p := range newProducts {
		if !containsProduct(oldProducts, p) {
			missingProducts = append(missingProducts, p)
		}
	}

	for _, q := range missingProducts {
		fmt.Printf("%d: %s\n", q.ID, q.Title)
	}
}
