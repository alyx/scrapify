package scrapify

// ShopifyCollections is a struct that represents the response from the
// collections endpoint.
type ShopifyCollections struct {
	Collections []Collection `json:"collections"`
}

// ShopifyCollectionProducts is a struct that represents the response from the
// products endpoint.
type ShopifyCollectionProducts struct {
	Products []Product `json:"products"`
}

// Collection is a struct that represents a collection in a Shopify store.
type Collection struct {
	ID            int64       `json:"id"`
	Title         string      `json:"title"`
	Handle        string      `json:"handle"`
	Description   string      `json:"description"`
	PublishedAt   string      `json:"published_at"`
	UpdatedAt     string      `json:"updated_at"`
	Image         interface{} `json:"image"`
	ProductsCount int         `json:"products_count"`
}

// Product is a struct that represents a product in a Shopify store.
type Product struct {
	ID          int64    `json:"id"`
	Title       string   `json:"title"`
	Handle      string   `json:"handle"`
	BodyHTML    string   `json:"body_html"`
	PublishedAt string   `json:"published_at"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	Vendor      string   `json:"vendor"`
	ProductType string   `json:"product_type"`
	Tags        []string `json:"tags"`
	Variants    []struct {
		ID               int64       `json:"id"`
		Title            string      `json:"title"`
		Option1          string      `json:"option1"`
		Option2          interface{} `json:"option2"`
		Option3          interface{} `json:"option3"`
		Sku              string      `json:"sku"`
		RequiresShipping bool        `json:"requires_shipping"`
		Taxable          bool        `json:"taxable"`
		FeaturedImage    interface{} `json:"featured_image"`
		Available        bool        `json:"available"`
		Price            string      `json:"price"`
		Grams            int         `json:"grams"`
		CompareAtPrice   interface{} `json:"compare_at_price"`
		Position         int         `json:"position"`
		ProductID        int64       `json:"product_id"`
		CreatedAt        string      `json:"created_at"`
		UpdatedAt        string      `json:"updated_at"`
	} `json:"variants"`
	Images []struct {
		ID         int64         `json:"id"`
		CreatedAt  string        `json:"created_at"`
		Position   int           `json:"position"`
		UpdatedAt  string        `json:"updated_at"`
		ProductID  int64         `json:"product_id"`
		VariantIds []interface{} `json:"variant_ids"`
		Src        string        `json:"src"`
		Width      int           `json:"width"`
		Height     int           `json:"height"`
	} `json:"images"`
	Options []struct {
		Name     string   `json:"name"`
		Position int      `json:"position"`
		Values   []string `json:"values"`
	} `json:"options"`
}
