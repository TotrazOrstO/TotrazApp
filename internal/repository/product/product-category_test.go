package product

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"totraz_store/internal/domain"
	"totraz_store/pkg/config"
	"totraz_store/pkg/store/postgres"
)

func TestAllProductsCategory(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := NewProductRepository(db)

	productsCategory, err := r.AllProductCategory(context.Background(), 10, 0)
	require.NoError(t, err)

	prettyP, _ := json.MarshalIndent(productsCategory, "", "  ")
	t.Log(string(prettyP))
}

func TestProductsCategoryById(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := NewProductRepository(db)

	products, err := r.ProductCategoryById(context.Background(), "2e5493a9-7dd8-4445-922c-84bbacf52907")
	require.NoError(t, err)

	prettyP, _ := json.MarshalIndent(products, "", "  ")
	t.Log(string(prettyP))
}

func TestCreatedProductCategory(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := NewProductRepository(db)

	product := domain.ProductCategory{Name: "test"}
	newProductCategory, err := r.CreateProductCategory(context.Background(), product)
	require.NoError(t, err)

	assert.Equal(t, product.Name, newProductCategory.Name)
	assert.NotEmpty(t, newProductCategory.ID)
}

func TestDeletedProductCategory(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := NewProductRepository(db)

	productCategory := domain.ProductCategory{Name: "test"}
	testProductCategory, err := r.CreateProductCategory(context.Background(), productCategory)
	require.NoError(t, err)
	err = r.DeleteProductCategory(context.Background(), testProductCategory.ID)
	require.NoError(t, err)
}
