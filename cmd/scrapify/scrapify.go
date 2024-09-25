package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/alyx/scrapify"
)

func writeProducts(products []scrapify.Product, file string) (int, error) {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	err = enc.Encode(products)
	if err != nil {
		return 0, err
	}
	return len(products), nil
}

func main() {
	flagListCollections := flag.Bool("l", false, "List all collections")
	flagCollections := flag.String("c", "", "Download products only from the given collections")
	flagOutput := flag.String("o", "products.json", "Output file")
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("No site provided")
		return
	}
	site := flag.Args()[0]
	if *flagListCollections {
		fmt.Println("Listing collections")
		collections, err := scrapify.GetCollections(site)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, c := range collections {
			fmt.Println(c.Title)
		}
		return
	}

	var allProducts []scrapify.Product
	if *flagCollections != "" {
		collections := strings.Split(*flagCollections, ",")
		products, err := scrapify.GetProducts(site, collections)
		if err != nil {
			fmt.Println(err)
			return
		}
		allProducts = products
	} else {
		collections, err := scrapify.GetCollections(site)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, c := range collections {
			products, err := scrapify.GetProducts(site, []string{c.Handle})
			if err != nil {
				fmt.Println(err)
				return
			}
			allProducts = append(allProducts, products...)
		}
	}

	numProducts, err := writeProducts(allProducts, *flagOutput)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Wrote %d products to %s\n", numProducts, *flagOutput)
}
