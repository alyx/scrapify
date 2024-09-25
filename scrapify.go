package scrapify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

var USER_AGENT string = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"

var HTTP_RETRIES int = 3

// doRequest performs an HTTP request using our custom User-Agent and returns a
// byte array
func doRequest(url string) ([]byte, error) {
	var resp *http.Response
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", USER_AGENT)
	for i := 0; i <= HTTP_RETRIES; i++ {
		resp, err = client.Do(req)
		if err != nil {
			return nil, err
		}
		switch resp.StatusCode {
		case 200:
			break
		case 404:
			return nil, fmt.Errorf("404 Not Found")
		default:
			time.Sleep(time.Second * 60)
			continue
		}
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP status code %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// GetCollections returns a list of collections from a Shopify store.
func GetCollections(site string) ([]Collection, error) {
	page := 1
	var collections []Collection
	for {
		body, err := doRequest(site + "/collections.json?page=" + strconv.Itoa(page))
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		var container ShopifyCollections
		err = json.Unmarshal(body, &container)
		if err != nil {
			return nil, err
		}
		if len(container.Collections) != 0 {
			collections = append(collections, container.Collections...)
			page += 1
		} else {
			break
		}
	}

	return collections, nil
}

// GetProducts returns a list of products from the provided collections in a
// Shopify store.
func GetProducts(site string, collections []string) ([]Product, error) {
	var products []Product
	for _, c := range collections {
		page := 1
		for {
			body, err := doRequest(site + "/collections/" + c + "/products.json?page=" + strconv.Itoa(page))
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			var container ShopifyCollectionProducts
			err = json.Unmarshal(body, &container)
			if err != nil {
				return nil, err
			}
			if len(container.Products) != 0 {
				products = append(products, container.Products...)
				page += 1
			} else {
				break
			}
		}
	}
	return products, nil
}
