package http

import (
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func (s *Server) initRouter() {
	s.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOriginFunc:  func(origin string) (bool, error) { return true, nil },
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
	}))

	product := s.echo.Group("/product")
	store := s.echo.Group("/store")

	productCategory := s.echo.Group("/product_category")
	storeCategory := s.echo.Group("/store_category")

	product.GET("/", s.AllProducts)
	product.GET("/:id", s.ProductById)
	product.POST("/create", s.CreateProduct)
	product.DELETE("/delete", s.DeleteProduct)

	productCategory.GET("/", s.AllProductCategory)
	productCategory.GET("/:id_category", s.ProductCategoryById)
	productCategory.POST("/create_category", s.CreateProductCategory)
	productCategory.DELETE("/delete_category", s.DeleteProductCategory)

	store.GET("/", s.AllStore)
	store.GET("/:id", s.StoreById)
	store.POST("/create", s.CreateStore)
	store.DELETE("/delete", s.DeleteStore)

	storeCategory.GET("/category", s.AllStoreCategory)
	storeCategory.GET("/:id_category", s.StoreCategoryById)
	storeCategory.POST("/create_category", s.CreateStoreCategory)
	storeCategory.DELETE("/delete_category", s.DeleteStoreCategory)
	//s.echo.GET("/", s.GetAllProducts)
}
