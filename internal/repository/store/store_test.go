package store

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

func TestAllStores(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := NewStoreRepository(db)

	stores, err := r.AllStores(context.Background(), 10, 0)
	require.NoError(t, err)

	prettyP, _ := json.MarshalIndent(stores, "", "  ")
	t.Log(string(prettyP))
}

func TestStoresById(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := NewStoreRepository(db)

	stores, err := r.StoreById(context.Background(), "2e5493a9-7dd8-4445-922c-84bbacf52907")
	require.NoError(t, err)

	prettyP, _ := json.MarshalIndent(stores, "", "  ")
	t.Log(string(prettyP))
}

func TestCreated(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := NewStoreRepository(db)

	store := domain.Store{Name: "test"}
	newStore, err := r.Create(context.Background(), store)
	require.NoError(t, err)

	assert.Equal(t, store.Name, newStore.Name)
	assert.NotEmpty(t, newStore.ID)
}

func TestDeleted(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := NewStoreRepository(db)

	store := domain.Store{Name: "test"}
	testStore, err := r.Create(context.Background(), store)
	require.NoError(t, err)
	err = r.Delete(context.Background(), testStore.ID)
	require.NoError(t, err)
}

func TestImageToStore(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := NewStoreRepository(db)

	store := domain.Store{Name: "test"}
	testStore, err := r.Create(context.Background(), store)
	require.NoError(t, err)

	i := images.NewImageRepository(db)

	image := &domain.Image{Name: "test"}
	testImage, err := i.Create(context.Background(), image)
	require.NoError(t, err)
	err = r.AddImagesToStore(context.Background(), testStore.ID, testImage.Id)
	require.NoError(t, err)
}

func TestStoreToCategory(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := NewStoreRepository(db)

	store := domain.Store{Name: "test"}
	testStore, err := r.Create(context.Background(), store)
	require.NoError(t, err)

	storeCategory := domain.StoreCategory{Name: "test"}
	testStoreCategory, err := r.CreateStoreCategory(context.Background(), storeCategory)
	require.NoError(t, err)

	err = r.AddStoreToCategory(context.Background(), testStoreCategory.ID, testStore.ID)
	require.NoError(t, err)
}
