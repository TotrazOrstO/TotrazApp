package product

import (
	"context"
	"fmt"
	"totraz_store/internal/domain"
	"totraz_store/internal/repository"
)

type manager struct {
	imagesRepository  repository.ImageRepository
	productRepository repository.ProductRepository
}

func NewProductManager(productRepository repository.ProductRepository, imageRepository repository.ImageRepository) *manager {
	return &manager{
		productRepository: productRepository,
		imagesRepository:  imageRepository,
	}
}

func (s *manager) AllProducts(ctx context.Context, limit int, offset int) ([]domain.Product, error) {
	return s.productRepository.AllProducts(ctx, limit, offset)
}

func (s *manager) ProductById(ctx context.Context, id string) (domain.Product, error) {
	return s.productRepository.ProductById(ctx, id)
}

func (s *manager) CreateProduct(ctx context.Context, createProduct domain.CreateProduct) (domain.Product, error) {

	product := domain.Product{
		Name:              createProduct.Name,
		Images:            createProduct.Images,
		ProductCategories: createProduct.ProductCategoryIds,
	}

	newProduct, err := s.productRepository.Create(ctx, product)
	if err != nil {
		return domain.Product{}, fmt.Errorf("reppository: create product")
	}

	return newProduct, nil
}

func (s *manager) DeleteProduct(ctx context.Context, id string) error {
	return s.productRepository.Delete(ctx, id)
}

func (s *manager) AllProductCategory(ctx context.Context, limit int, offset int) ([]domain.ProductCategory, error) {
	return s.productRepository.AllProductCategory(ctx, limit, offset)
}

func (s *manager) ProductCategoryById(ctx context.Context, id string) (domain.ProductCategory, error) {
	return s.productRepository.ProductCategoryById(ctx, id)
}

func (s *manager) CreateProductCategory(ctx context.Context, createCategory domain.CreateProductCategory) (domain.ProductCategory, error) {
	category := domain.ProductCategory{
		Name: createCategory.Name,
	}

	newCategory, err := s.productRepository.CreateProductCategory(ctx, category)
	if err != nil {
		return domain.ProductCategory{}, fmt.Errorf("reppository: create product")
	}

	return newCategory, nil
}

func (s *manager) DeleteProductCategory(ctx context.Context, id string) error {
	return s.productRepository.DeleteProductCategory(ctx, id)
}
