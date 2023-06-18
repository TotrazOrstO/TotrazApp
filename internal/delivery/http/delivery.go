package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"totraz_store/internal/usecase"
	"totraz_store/pkg/config"
)

type Server struct {
	cfg  config.HTTP
	echo *echo.Echo

	productManager usecase.ProductManager
	storeManager   usecase.StoreManager
}

func NewServer(cfg config.HTTP,
	productManager usecase.ProductManager,
	storeManager usecase.StoreManager,
) *Server {
	d := &Server{
		cfg:            cfg,
		echo:           echo.New(),
		productManager: productManager,
		storeManager:   storeManager,
	}

	d.initRouter()

	return d
}

func (d *Server) Start(_ context.Context) error {
	err := d.echo.Start(fmt.Sprintf("%s:%d", d.cfg.Host, d.cfg.Port))
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (d *Server) Stop(ctx context.Context) error {
	return d.echo.Shutdown(ctx)
}
