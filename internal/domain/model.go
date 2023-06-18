package domain

type ProductCategory struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Product []Product `json:"product"`
}

type Product struct {
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	Images            []Image           `json:"image"`
	ProductCategories []ProductCategory `json:"product_categories"`
}

type StoreCategory struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Stores []Store `json:"stores"`
}

type Store struct {
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	Images          []Image         `json:"images"`
	StoreCategories []StoreCategory `json:"store_categories"`
}

type Image struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Ext  string `json:"ext"`
	Body []byte `json:"body"`
}
