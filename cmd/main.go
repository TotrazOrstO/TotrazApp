package main

import (
	"context"
	"log"
	"totraz_store/internal/delivery/http"
	"totraz_store/internal/repository/images"
	"totraz_store/internal/repository/product"
	"totraz_store/internal/repository/store"
	productUsecase "totraz_store/internal/usecase/product"
	storeUsecase "totraz_store/internal/usecase/store"
	"totraz_store/pkg/config"
	"totraz_store/pkg/store/postgres"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("new configs: %s", err.Error())
	}

	db, err := postgres.New(cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}

	productRepo := product.NewProductRepository(db)
	storeRepo := store.NewStoreRepository(db)
	imageRepo := images.NewImageRepository(db)

	productUseCase := productUsecase.NewProductManager(productRepo, imageRepo)
	storeUseCase := storeUsecase.NewStoreManager(storeRepo, imageRepo)

	server := http.NewServer(cfg.HTTP, productUseCase, storeUseCase)

	log.Fatal(server.Start(context.Background()))
}
