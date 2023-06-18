package images

import (
	"context"
	"fmt"
	"totraz_store/internal/domain"
)

func (m *Repository) Create(ctx context.Context, image *domain.Image) (*domain.Image, error) {
	q := fmt.Sprintf(`
INSERT INTO images(file_name, file_ext, file_body)
VALUES ($1, $2, $3) RETURNING id`)

	err := m.db.QueryRowContext(ctx, q, image.Name, image.Ext, image.Body).Scan(&image.Id)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return image, nil
}
