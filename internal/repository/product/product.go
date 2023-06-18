package product

import (
	"context"
	"fmt"
	"totraz_store/internal/domain"
)

func (m *Repository) AllProducts(ctx context.Context, limit int, offset int) ([]domain.Product, error) {
	q := fmt.Sprintf(`
SELECT id, name
FROM product
LIMIT $1
OFFSET $2`)

	rows, err := m.db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
		)

		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		q = fmt.Sprintf(`SELECT i.id img_id, i.file_name img_name, i.file_ext img_ext, i.file_body img_body
FROM image_to_product itp
         JOIN images i on i.id = itp.image_id
WHERE itp.product_id = $1
`)

		imageRows, err := m.db.QueryContext(ctx, q, product.ID)
		if err != nil {
			return nil, fmt.Errorf("get images: %w", err)
		}
		if err != nil {
			return nil, fmt.Errorf("image rows: %w", err)
		}
		defer imageRows.Close()

		for imageRows.Next() {
			var img domain.Image
			err := imageRows.Scan(
				&img.Id,
				&img.Name,
				&img.Ext,
				&img.Body,
			)

			if err != nil {
				return nil, fmt.Errorf("scan imageRows: %w", err)
			}

			product.Images = append(product.Images, img)
		}

		q = `SELECT pc.id, pc.name
FROM product_to_category ptc
JOIN product_category pc on ptc.product_category_id = pc.id
WHERE ptc.product_id=$1`

		categoryRows, err := m.db.QueryContext(ctx, q, product.ID)
		if err != nil {
			return nil, fmt.Errorf("query: %w", err)
		}
		defer categoryRows.Close()

		for categoryRows.Next() {
			var category domain.ProductCategory
			err := rows.Scan(
				&category.ID,
				&category.Name,
			)

			if err != nil {
				return nil, fmt.Errorf("category scan: %w", err)
			}
			product.ProductCategories = append(product.ProductCategories, category)
		}

		products = append(products, product)
	}

	return products, nil
}

func (m *Repository) ProductById(ctx context.Context, id string) (domain.Product, error) {
	q := fmt.Sprintf(`
SELECT id, name
FROM product
WHERE id = $1`)

	rows, err := m.db.QueryContext(ctx, q, id)
	if err != nil {
		return domain.Product{}, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	var product domain.Product
	for rows.Next() {
		err := rows.Scan(
			&product.ID,
			&product.Name,
		)

		if err != nil {
			return domain.Product{}, fmt.Errorf("scan: %w", err)

		}

		q = fmt.Sprintf(`SELECT i.id img_id, i.file_name img_name, i.file_ext img_ext, i.file_body img_body
FROM image_to_product itp
         JOIN images i on i.id = itp.image_id
WHERE itp.product_id = $1
`)

		imageRows, err := m.db.QueryContext(ctx, q, product.ID)
		if err != nil {
			return domain.Product{}, fmt.Errorf("get images: %w", err)
		}
		defer imageRows.Close()

		for imageRows.Next() {
			var img domain.Image
			err := imageRows.Scan(
				&img.Id,
				&img.Name,
				&img.Ext,
				&img.Body,
			)

			if err != nil {
				return domain.Product{}, fmt.Errorf("scan imageRows: %w", err)
			}

			product.Images = append(product.Images, img)
		}

		q = `SELECT pc.id, pc.name
FROM product_to_category ptc
JOIN product_category pc on ptc.product_category_id = pc.id
WHERE ptc.product_id=$1`

		categoryRows, err := m.db.QueryContext(ctx, q, product.ID)
		if err != nil {
			return domain.Product{}, fmt.Errorf("query: %w", err)
		}
		defer categoryRows.Close()

		for categoryRows.Next() {
			var category domain.ProductCategory
			err := rows.Scan(
				&category.ID,
				&category.Name,
			)

			if err != nil {
				return domain.Product{}, fmt.Errorf("category scan: %w", err)
			}
			product.ProductCategories = append(product.ProductCategories, category)
		}
	}

	return product, nil
}

func (m *Repository) Create(ctx context.Context, product domain.Product) (domain.Product, error) {
	q := fmt.Sprintf(`
	INSERT INTO product(name)
	VALUES ($1) RETURNING id, name
	`)

	var newProduct domain.Product
	err := m.db.QueryRowContext(ctx, q, product.Name).Scan(&newProduct.ID, &newProduct.Name)
	if err != nil {
		return domain.Product{}, fmt.Errorf("query: %w", err)
	}

	return newProduct, nil
}

func (m *Repository) AddImagesToProduct(ctx context.Context, productId string, imageId string) error {
	q := fmt.Sprintf(`
INSERT INTO image_to_product(image_id, product_id)
VALUES ($1, $2)
`)

	_, err := m.db.ExecContext(ctx, q, imageId, productId)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	return nil
}

func (m *Repository) AddProductToCategory(ctx context.Context, productCategoryId string, productId string) error {
	q := fmt.Sprintf(`
INSERT INTO product_to_category(product_category_id, product_id)
VALUES ($1, $2)
`)

	_, err := m.db.ExecContext(ctx, q, productCategoryId, productId)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	return nil
}

func (m *Repository) Delete(ctx context.Context, id string) error {
	q := fmt.Sprintf(`
DELETE FROM product
WHERE id = $1`)

	_, err := m.db.ExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
