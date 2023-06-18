package http

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"totraz_store/internal/domain"
)

func (s *Server) AllStore(c echo.Context) error {
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

	stores, err := s.storeManager.AllStore(ctx, limit, offset)
	if err != nil {
		log.WithError(err).Error("failed to get all stores")
		return c.String(http.StatusBadRequest, "failed to get all stores")
	}

	return c.JSON(http.StatusOK, stores)
}
func (s *Server) StoreById(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	log := logrus.WithFields(logrus.Fields{
		"id": id,
	})

	store, err := s.storeManager.StoreById(ctx, id)
	if err != nil {
		log.WithError(err).Error("failed to parse store")
		return c.String(http.StatusBadRequest, "failed to parse store: "+err.Error())
	}

	return c.JSON(http.StatusOK, store)
}
func (s *Server) CreateStore(c echo.Context) error {
	ctx := c.Request().Context()

	var newStore domain.CreateStore

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

	err := c.Bind(&newStore)
	if err != nil {
		logrus.WithError(err).Error("failed to bind store")
		return c.String(http.StatusBadRequest, "failed to bind store")
	}

	store, err := s.storeManager.CreateStore(ctx, newStore)
	if err != nil {
		logrus.WithError(err).Error("failed to create store")
		return c.String(http.StatusBadRequest, "failed to create store")
	}

	return c.JSON(http.StatusCreated, store)
}
func (s *Server) DeleteStore(c echo.Context) error {
	ctx := c.Request().Context()
	ds := domain.DeleteStore{}

	err := c.Bind(&ds)
	if err != nil {
		logrus.WithError(err).Error("failed to bind store")
		return c.String(http.StatusBadRequest, "failed to bind store")
	}

	err = s.storeManager.DeleteStore(ctx, ds.ID)
	if err != nil {
		logrus.WithError(err).Error("failed to delete store")
		return c.String(http.StatusBadRequest, "failed to delete store")
	}

	return c.String(http.StatusOK, "deleted")
}

func (s *Server) AllStoreCategory(c echo.Context) error {
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

	allStoreCategories, err := s.storeManager.AllStoreCategory(ctx, limit, offset)
	if err != nil {
		log.WithError(err).Error("failed to get all store categories")
		return c.String(http.StatusBadRequest, "failed to get all store categories: "+err.Error())
	}

	return c.JSON(http.StatusCreated, allStoreCategories)
}

func (s *Server) StoreCategoryById(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	log := logrus.WithFields(logrus.Fields{
		"id": id,
	})

	storeCategory, err := s.storeManager.StoreCategoryById(ctx, id)
	if err != nil {
		log.WithError(err).Error("failed to parse store category")
		return c.String(http.StatusBadRequest, "failed to parse store category: "+err.Error())
	}

	return c.JSON(http.StatusOK, storeCategory)
}
func (s *Server) CreateStoreCategory(c echo.Context) error {
	ctx := c.Request().Context()

	var newStoreCategory domain.CreateStoreCategory

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

	err := c.Bind(&newStoreCategory)
	if err != nil {
		logrus.WithError(err).Error("failed to bind store category")
		return c.String(http.StatusBadRequest, "failed to bind store category")
	}

	storeCategory, err := s.storeManager.CreateStoreCategory(ctx, newStoreCategory)
	if err != nil {
		logrus.WithError(err).Error("failed to create store category")
		return c.String(http.StatusBadRequest, "failed to create store category")
	}

	return c.JSON(http.StatusCreated, storeCategory)
}
func (s *Server) DeleteStoreCategory(c echo.Context) error {
	ctx := c.Request().Context()
	dsc := domain.DeleteStoreCategory{}

	err := c.Bind(&dsc)
	if err != nil {
		logrus.WithError(err).Error("failed to bind store category")
		return c.String(http.StatusBadRequest, "failed to bind store category")
	}

	err = s.storeManager.DeleteStoreCategory(ctx, dsc.ID)
	if err != nil {
		logrus.WithError(err).Error("failed to delete store category")
		return c.String(http.StatusBadRequest, "failed to delete store category")
	}

	return c.String(http.StatusOK, "deleted")
}
