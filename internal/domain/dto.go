package domain

type CreateProduct struct {
	ID                 string            `json:"id"`
	Name               string            `json:"name"`
	Images             []Image           `json:"images"`
	ProductCategoryIds []ProductCategory `json:"product_category_ids"`
}

type DeleteProduct struct {
	ID string `json:"id"`
}

type CreateProductCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DeleteProductCategory struct {
	ID string `json:"id"`
}

type CreateStore struct {
	ID               string          `json:"id"`
	Name             string          `json:"name"`
	Images           []Image         `json:"images"`
	StoreCategoryIds []StoreCategory `json:"storeCategoryIds"`
}

type DeleteStore struct {
	ID string `json:"id"`
}

type CreateStoreCategory struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image Image  `json:"image"`
}

type DeleteStoreCategory struct {
	ID string `json:"id"`
}
