package repository

import (
	"context"
	"totraz_store/internal/domain"
)

type ProductRepository interface {
	AllProducts(ctx context.Context, limit int, offset int) ([]domain.Product, error)
	ProductById(ctx context.Context, id string) (domain.Product, error)
	Create(ctx context.Context, product domain.Product) (domain.Product, error)
	Delete(ctx context.Context, id string) error

	AllProductCategory(ctx context.Context, limit int, offset int) ([]domain.ProductCategory, error)
	ProductCategoryById(ctx context.Context, id string) (domain.ProductCategory, error)
	CreateProductCategory(ctx context.Context, product domain.ProductCategory) (domain.ProductCategory, error)
	DeleteProductCategory(ctx context.Context, id string) error

	AddImagesToProduct(ctx context.Context, productId string, imageId string) error
	AddProductToCategory(ctx context.Context, productCategoryId string, productId string) error
}

type StoreRepository interface {
	AllStores(ctx context.Context, limit int, offset int) ([]domain.Store, error)
	StoreById(ctx context.Context, id string) (domain.Store, error)
	Create(ctx context.Context, store domain.Store) (domain.Store, error)
	Delete(ctx context.Context, id string) error

	//Store Category

	AllStoreCategory(ctx context.Context, limit int, offset int) ([]domain.StoreCategory, error)
	StoreCategoryById(ctx context.Context, id string) (domain.StoreCategory, error)
	CreateStoreCategory(ctx context.Context, store domain.StoreCategory) (domain.StoreCategory, error)
	DeleteStoreCategory(ctx context.Context, id string) error

	AddImagesToStore(ctx context.Context, storeId string, imageId string) error
	AddStoreToCategory(ctx context.Context, storeCategoryId string, storeId string) error

	AddProductCategoryToStore(ctx context.Context, productCategoryId string, storeId string) error
}

type ImageRepository interface {
	Create(ctx context.Context, image *domain.Image) (*domain.Image, error)
}
