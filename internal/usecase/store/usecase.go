package store

import (
	"context"
	"fmt"
	"totraz_store/internal/domain"
	"totraz_store/internal/repository"
)

type manager struct {
	imagesRepository repository.ImageRepository
	storeRepository  repository.StoreRepository
}

func NewStoreManager(storage repository.StoreRepository, imagesRepository repository.ImageRepository) *manager {
	return &manager{
		storeRepository:  storage,
		imagesRepository: imagesRepository,
	}
}

func (s *manager) AllStore(ctx context.Context, limit int, offset int) ([]domain.Store, error) {
	return s.storeRepository.AllStores(ctx, limit, offset)
}

func (s *manager) StoreById(ctx context.Context, id string) (domain.Store, error) {
	return s.storeRepository.StoreById(ctx, id)
}

func (s *manager) CreateStore(ctx context.Context, createStore domain.CreateStore) (domain.Store, error) {
	store := domain.Store{
		Name: createStore.Name,
	}

	newStore, err := s.storeRepository.Create(ctx, store)
	if err != nil {
		return domain.Store{}, fmt.Errorf("reppository: create store")
	}

	return newStore, nil
}

func (s *manager) DeleteStore(ctx context.Context, id string) error {
	return s.storeRepository.Delete(ctx, id)
}

//Store Category

func (s *manager) AllStoreCategory(ctx context.Context, limit int, offset int) ([]domain.StoreCategory, error) {
	return s.storeRepository.AllStoreCategory(ctx, limit, offset)
}

func (s *manager) StoreCategoryById(ctx context.Context, id string) (domain.StoreCategory, error) {
	return s.storeRepository.StoreCategoryById(ctx, id)
}

func (s *manager) CreateStoreCategory(ctx context.Context, createCategory domain.CreateStoreCategory) (domain.StoreCategory, error) {
	category := domain.StoreCategory{
		Name: createCategory.Name,
	}

	newCategory, err := s.storeRepository.CreateStoreCategory(ctx, category)
	if err != nil {
		return domain.StoreCategory{}, fmt.Errorf("reppository: create store")
	}

	return newCategory, nil
}

func (s *manager) DeleteStoreCategory(ctx context.Context, id string) error {
	return s.storeRepository.Delete(ctx, id)
}
