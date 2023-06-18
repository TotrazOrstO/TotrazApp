package usecase

import (
	"context"
	"totraz_store/internal/domain"
)

type ProductManager interface {
	AllProducts(ctx context.Context, limit int, offset int) ([]domain.Product, error)
	ProductById(ctx context.Context, id string) (domain.Product, error)
	CreateProduct(ctx context.Context, product domain.CreateProduct) (domain.Product, error)
	DeleteProduct(ctx context.Context, id string) error

	//Product Category

	AllProductCategory(ctx context.Context, limit int, offset int) ([]domain.ProductCategory, error)
	ProductCategoryById(ctx context.Context, id string) (domain.ProductCategory, error)
	CreateProductCategory(ctx context.Context, productCategory domain.CreateProductCategory) (domain.ProductCategory, error)
	DeleteProductCategory(ctx context.Context, id string) error
}

type StoreManager interface {
	AllStore(ctx context.Context, limit int, offset int) ([]domain.Store, error)
	StoreById(ctx context.Context, id string) (domain.Store, error)
	CreateStore(ctx context.Context, store domain.CreateStore) (domain.Store, error)
	DeleteStore(ctx context.Context, id string) error

	//Store Category

	AllStoreCategory(ctx context.Context, limit int, offset int) ([]domain.StoreCategory, error)
	StoreCategoryById(ctx context.Context, id string) (domain.StoreCategory, error)
	CreateStoreCategory(ctx context.Context, storeCategory domain.CreateStoreCategory) (domain.StoreCategory, error)
	DeleteStoreCategory(ctx context.Context, id string) error
}
