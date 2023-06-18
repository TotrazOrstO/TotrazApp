package store

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

func TestAllStoresCategory(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := NewStoreRepository(db)

	storesCategory, err := r.AllStoreCategory(context.Background(), 10, 0)
	require.NoError(t, err)

	prettyP, _ := json.MarshalIndent(storesCategory, "", "  ")
	t.Log(string(prettyP))
}

func TestStoresCategoryById(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := NewStoreRepository(db)

	stores, err := r.StoreCategoryById(context.Background(), "2e5493a9-7dd8-4445-922c-84bbacf52907")
	require.NoError(t, err)

	prettyP, _ := json.MarshalIndent(stores, "", "  ")
	t.Log(string(prettyP))
}

func TestCreatedStoreCategory(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := NewStoreRepository(db)

	store := domain.StoreCategory{Name: "test"}
	newStoreCategory, err := r.CreateStoreCategory(context.Background(), store)
	require.NoError(t, err)

	assert.Equal(t, store.Name, newStoreCategory.Name)
	assert.NotEmpty(t, newStoreCategory.ID)
}

func TestDeletedStoreCategory(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := NewStoreRepository(db)

	storeCategory := domain.StoreCategory{Name: "test"}
	testStoreCategory, err := r.CreateStoreCategory(context.Background(), storeCategory)
	require.NoError(t, err)
	err = r.DeleteStoreCategory(context.Background(), testStoreCategory.ID)
	require.NoError(t, err)
}
