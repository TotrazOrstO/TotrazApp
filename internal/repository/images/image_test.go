package images

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
	"totraz_store/internal/domain"
	"totraz_store/pkg/config"
	"totraz_store/pkg/store/postgres"
)

func TestCreateImage(t *testing.T) {
	cfg := config.TestConfigs(t)

	db, err := postgres.New(cfg.Postgres)
	require.NoError(t, err)

	r := NewImageRepository(db)

	imageTest := domain.Image{Name: "test", Ext: "png", Body: []byte("dsdsfsfdsfsfd")}
	image, err := r.Create(context.Background(), &imageTest)
	require.NoError(t, err)

	prettyP, _ := json.MarshalIndent(image, "", "  ")
	t.Log(string(prettyP))
}
