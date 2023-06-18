package product

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"totraz_store/internal/domain"
	"totraz_store/internal/repository/images"
	"totraz_store/internal/repository/product"
	"totraz_store/pkg/config"
	"totraz_store/pkg/store/postgres"
)

//func AllProduct(t *testing.T) {
//	cfg := config.TestConfigs(t)
//
//	db, err := postgres.New(cfg.Postgres)
//	require.NoError(t, err)
//
//	r := product.NewProductRepository(db)
//
//	i := images.NewImageRepository(db)
//
//	m := NewProductManager(r, i)
//
//	products, err := m.AllProducts(context.Background(), 0, 10)
//	assert.NoError(t, err)
//	t.Log(products)
//}
//
//func ProductById(t *testing.T) {
//	cfg := config.TestConfigs(t)
//
//	db, err := postgres.New(cfg.Postgres)
//	require.NoError(t, err)
//
//	r := product.NewProductRepository(db)
//
//	i := images.NewImageRepository(db)
//
//	m := NewProductManager(r, i)
//
//	product, err := m.ProductById(context.Background(), "3dc5a788-44cf-447a-a05f-7e551257e9ac")
//	assert.NoError(t, err)
//	t.Log(product)
//}

func TestCreated(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := product.NewProductRepository(db)

	i := images.NewImageRepository(db)

	m := NewProductManager(r, i)

	pc := domain.ProductCategory{ID: "642826c5-96ec-4477-a8c4-57f48ed9f287"}

	createProduct := domain.CreateProduct{
		Name:               "test_from_usecase",
		ProductCategoryIds: append([]domain.ProductCategory{}, pc),
		Images: []domain.Image{
			{
				Name: "t1",
				Ext:  "jpg",
				Body: []byte("image"),
			},
		},
	}

	product, err := m.CreateProduct(context.Background(), createProduct)
	assert.NoError(t, err)
	t.Log(product)
}

func TestCreateProductCategory(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := product.NewProductRepository(db)

	i := images.NewImageRepository(db)

	m := NewProductManager(r, i)

	createProductCategory := domain.CreateProductCategory{
		Name: "test_from_usecase",
	}

	productCategory, err := m.CreateProductCategory(context.Background(), createProductCategory)
	assert.NoError(t, err)
	t.Log(productCategory)
}
