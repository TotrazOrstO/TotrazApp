package http

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"totraz_store/internal/domain"
)

func (s *Server) AllProducts(c echo.Context) error {
	ctx := c.Request().Context()
	limitParam := c.QueryParam("limit")
	offsetParam := c.QueryParam("offset")

	log := logrus.WithFields(logrus.Fields{
		"limit":  limitParam,
		"offset": offsetParam,
	})

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		log.WithError(err).Error("failed to parse limit")
		return c.String(http.StatusBadRequest, "failed to parse limit: "+err.Error())
	}
	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		log.WithError(err).Error("failed to parse offset")
		return c.String(http.StatusBadRequest, "failed to parse offset: "+err.Error())
	}

	products, err := s.productManager.AllProducts(ctx, limit, offset)
	if err != nil {
		log.WithError(err).Error("failed to get all products")
		return c.String(http.StatusBadRequest, "failed to get all products")
	}

	return c.JSON(http.StatusOK, products)
}

func (s *Server) ProductById(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	log := logrus.WithFields(logrus.Fields{
		"id": id,
	})

	product, err := s.productManager.ProductById(ctx, id)
	if err != nil {
		log.WithError(err).Error("failed to parse product")
		return c.String(http.StatusBadRequest, "failed to parse product: "+err.Error())
	}

	return c.JSON(http.StatusOK, product)
}
func (s *Server) CreateProduct(c echo.Context) error {
	ctx := c.Request().Context()

	var newProduct domain.CreateProduct

	/*
		{
			"name": fff,
			"images":[
						{
							"name": fff,
							"ext":jpg,
							"body":"lkkhjhjggh"
						}
					]
		}
	*/

	err := c.Bind(&newProduct)
	if err != nil {
		logrus.WithError(err).Error("failed to bind product")
		return c.String(http.StatusBadRequest, "failed to bind product")
	}

	product, err := s.productManager.CreateProduct(ctx, newProduct)
	if err != nil {
		logrus.WithError(err).Error("failed to create product")
		return c.String(http.StatusBadRequest, "failed to create product")
	}

	return c.JSON(http.StatusCreated, product)
}
func (s *Server) DeleteProduct(c echo.Context) error {
	ctx := c.Request().Context()
	dp := domain.DeleteProduct{}

	err := c.Bind(&dp)
	if err != nil {
		logrus.WithError(err).Error("failed to bind product")
		return c.String(http.StatusBadRequest, "failed to bind product")
	}

	err = s.productManager.DeleteProduct(ctx, dp.ID)
	if err != nil {
		logrus.WithError(err).Error("failed to delete product")
		return c.String(http.StatusBadRequest, "failed to delete product")
	}

	return c.String(http.StatusOK, "deleted")
}
func (s *Server) AllProductCategory(c echo.Context) error {
	ctx := c.Request().Context()
	limitParam := c.QueryParam("limit")
	offsetParam := c.QueryParam("offset")

	log := logrus.WithFields(logrus.Fields{
		"limit":  limitParam,
		"offset": offsetParam,
	})

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		log.WithError(err).Error("failed to parse limit")
		return c.String(http.StatusBadRequest, "failed to parse limit: "+err.Error())
	}
	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		log.WithError(err).Error("failed to parse offset")
		return c.String(http.StatusBadRequest, "failed to parse offset: "+err.Error())
	}

	allProductCategories, err := s.productManager.AllProductCategory(ctx, limit, offset)
	if err != nil {
		log.WithError(err).Error("failed to get all product categories")
		return c.String(http.StatusBadRequest, "failed to get all product categories: "+err.Error())
	}

	return c.JSON(http.StatusCreated, allProductCategories)
}
func (s *Server) ProductCategoryById(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	log := logrus.WithFields(logrus.Fields{
		"id": id,
	})

	productCategory, err := s.productManager.ProductCategoryById(ctx, id)
	if err != nil {
		log.WithError(err).Error("failed to parse product category")
		return c.String(http.StatusBadRequest, "failed to parse product category: "+err.Error())
	}

	return c.JSON(http.StatusOK, productCategory)
}
func (s *Server) CreateProductCategory(c echo.Context) error {
	ctx := c.Request().Context()

	var newProductCategory domain.CreateProductCategory

	/*
		{
			"name": fff,
			"images":[
						{
							"name": fff,
							"ext":jpg,
							"body":"lkkhjhjggh"
						}
					]
		}
	*/

	err := c.Bind(&newProductCategory)
	if err != nil {
		logrus.WithError(err).Error("failed to bind product category")
		return c.String(http.StatusBadRequest, "failed to bind product category")
	}

	productCategory, err := s.productManager.CreateProductCategory(ctx, newProductCategory)
	if err != nil {
		logrus.WithError(err).Error("failed to create product category")
		return c.String(http.StatusBadRequest, "failed to create product category")
	}

	return c.JSON(http.StatusCreated, productCategory)
}
func (s *Server) DeleteProductCategory(c echo.Context) error {
	ctx := c.Request().Context()
	dpc := domain.DeleteProductCategory{}

	err := c.Bind(&dpc)
	if err != nil {
		logrus.WithError(err).Error("failed to bind product category")
		return c.String(http.StatusBadRequest, "failed to bind product category")
	}

	err = s.productManager.DeleteProductCategory(ctx, dpc.ID)
	if err != nil {
		logrus.WithError(err).Error("failed to delete product category")
		return c.String(http.StatusBadRequest, "failed to delete product category")
	}

	return c.String(http.StatusOK, "deleted")
}
