# scrapify

A library and CLI tools for scraping products from Shopify stores.

## Installation

```bash
go install github.com/alyx/scrapify/cmd/scrapify@latest
```

## Usage

```
scrapify -h
Usage of scrapify:
  -c string
        Download products only from the given collections
  -l    List all collections
```

## Example

```bash
scrapify -l
Listing collections

scrapify -c "collection-handle"
Wrote 100 products to products.json
```

## Library

```bash
go get github.com/alyx/scrapify
```

```go
package main

import (
	"github.com/alyx/scrapify"
)

func main() {
	collections, err := scrapify.GetCollections("https://example.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	products, err := scrapify.GetProducts("https://example.com", collections)
	if err != nil {
		fmt.Println(err)
		return
	}
}