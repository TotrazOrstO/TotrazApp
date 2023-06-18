package product

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"totraz_store/internal/domain"
	"totraz_store/internal/repository/images"
	"totraz_store/pkg/config"
	"totraz_store/pkg/store/postgres"
)

func TestAllProducts(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := NewProductRepository(db)

	products, err := r.AllProducts(context.Background(), 10, 0)
	require.NoError(t, err)

	prettyP, _ := json.MarshalIndent(products, "", "  ")
	t.Log(string(prettyP))
}

func TestProductsById(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := NewProductRepository(db)

	products, err := r.ProductById(context.Background(), "2e5493a9-7dd8-4445-922c-84bbacf52907")
	require.NoError(t, err)

	prettyP, _ := json.MarshalIndent(products, "", "  ")
	t.Log(string(prettyP))
}

func TestCreated(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := NewProductRepository(db)

	product := domain.Product{Name: "test"}
	newProduct, err := r.Create(context.Background(), product)
	require.NoError(t, err)

	assert.Equal(t, product.Name, newProduct.Name)
	assert.NotEmpty(t, newProduct.ID)
}

func TestDeleted(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := NewProductRepository(db)

	product := domain.Product{Name: "test"}
	testProduct, err := r.Create(context.Background(), product)
	require.NoError(t, err)
	err = r.Delete(context.Background(), testProduct.ID)
	require.NoError(t, err)
}

func TestImageToProduct(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := NewProductRepository(db)

	product := domain.Product{Name: "test"}
	testProduct, err := r.Create(context.Background(), product)
	require.NoError(t, err)

	i := images.NewImageRepository(db)

	image := &domain.Image{Name: "test"}
	testImage, err := i.Create(context.Background(), image)
	require.NoError(t, err)
	err = r.AddImagesToProduct(context.Background(), testProduct.ID, testImage.Id)
	require.NoError(t, err)
}

func TestProductToCategory(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := NewProductRepository(db)

	product := domain.Product{Name: "test"}
	testProduct, err := r.Create(context.Background(), product)
	require.NoError(t, err)

	productCategory := domain.ProductCategory{Name: "test"}
	testProductCategory, err := r.CreateProductCategory(context.Background(), productCategory)
	require.NoError(t, err)
	err = r.AddProductToCategory(context.Background(), testProductCategory.ID, testProduct.ID)
	require.NoError(t, err)
}
