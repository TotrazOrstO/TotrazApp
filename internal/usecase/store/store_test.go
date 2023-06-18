package store

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"totraz_store/internal/domain"
	"totraz_store/internal/repository/images"
	"totraz_store/internal/repository/store"
	"totraz_store/pkg/config"
	"totraz_store/pkg/store/postgres"
)

func TestCreated(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := store.NewStoreRepository(db)

	i := images.NewImageRepository(db)

	m := NewStoreManager(r, i)

	sc := domain.StoreCategory{ID: "642826c5-96ec-4477-a8c4-57f48ed9f287"}

	createStore := domain.CreateStore{
		ID:               "123",
		Name:             "test_from_usecase",
		StoreCategoryIds: append([]domain.StoreCategory{}, sc),
		Images: []domain.Image{
			{
				Name: "t1",
				Ext:  "jpg",
				Body: []byte("image"),
			},
		},
	}

	store, err := m.CreateStore(context.Background(), createStore)
	assert.NoError(t, err)
	t.Log(store)
}

func TestCreatedCategory(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := store.NewStoreRepository(db)

	i := images.NewImageRepository(db)

	m := NewStoreManager(r, i)

	createStoreCategory := domain.CreateStoreCategory{
		Name: "test_from_usecase",
	}

	storeCategory, err := m.CreateStoreCategory(context.Background(), createStoreCategory)
	assert.NoError(t, err)
	t.Log(storeCategory)
}
