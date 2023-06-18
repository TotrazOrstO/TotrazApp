package product

import (
	"context"
	"fmt"
	"totraz_store/internal/domain"
)

func (m *Repository) AllProductCategory(ctx context.Context, limit int, offset int) ([]domain.ProductCategory, error) {
	q := fmt.Sprintf(`
SELECT id, name
FROM product_category
LIMIT $1
OFFSET $2`)

	rows, err := m.db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	var ProductCategory []domain.ProductCategory
	for rows.Next() {
		var productCategory domain.ProductCategory
		err := rows.Scan(
			&productCategory.ID,
			&productCategory.Name,
			//&productCategory.Product,
		)

		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		q = `SELECT ptc.product_category_id, ptc.product_id
FROM product_to_category ptc JOIN product_category pc
ON ptc.product_category_id = pc.id`

		productRows, err := m.db.QueryContext(ctx, q)
		if err != nil {
			return nil, fmt.Errorf("query: %w", err)
		}
		defer productRows.Close()

		var Products []domain.Product
		for productRows.Next() {
			var product domain.Product
			err := productRows.Scan(
				&product.ID,
				&product.Name,
			)

			if err != nil {
				return nil, fmt.Errorf("category scan: %w", err)
			}

			q = fmt.Sprintf(`
SELECT i.id img_id, i.file_name img_name, i.file_ext img_ext, i.file_body img_body
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
				Products = append(Products, product)
				productCategory.Product = Products
			}

		}
		ProductCategory = append(ProductCategory, productCategory)
	}

	return ProductCategory, nil
}

func (m *Repository) ProductCategoryById(ctx context.Context, id string) (domain.ProductCategory, error) {
	q := fmt.Sprintf(`
SELECT id
FROM product_category pc
WHERE pc.id = $1`)

	rows, err := m.db.QueryContext(ctx, q, id)
	if err != nil {
		return domain.ProductCategory{}, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	var ProductCategory domain.ProductCategory
	for rows.Next() {
		var productCategory domain.ProductCategory
		err := rows.Scan(
			&productCategory.ID,
			&productCategory.Name,
		)

		if err != nil {
			return domain.ProductCategory{}, fmt.Errorf("scan imageRows: %w", err)
		}

		q = `SELECT ptc.product_category_id, ptc.product_id
FROM product_to_category ptc JOIN product_category pc
ON ptc.product_category_id = pc.id`

		productRows, err := m.db.QueryContext(ctx, q)
		if err != nil {
			return domain.ProductCategory{}, fmt.Errorf("query: %w", err)
		}
		defer productRows.Close()

		var Products []domain.Product
		for productRows.Next() {
			var product domain.Product
			err := productRows.Scan(
				&product.ID,
				&product.Name,
			)

			if err != nil {
				return domain.ProductCategory{}, fmt.Errorf("category scan: %w", err)
			}

			q = fmt.Sprintf(`
SELECT i.id img_id, i.file_name img_name, i.file_ext img_ext, i.file_body img_body
FROM image_to_product itp
         JOIN images i on i.id = itp.image_id
WHERE itp.product_id = $1
`)

			imageRows, err := m.db.QueryContext(ctx, q, product.ID)
			if err != nil {
				return domain.ProductCategory{}, fmt.Errorf("get images: %w", err)
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
					return domain.ProductCategory{}, fmt.Errorf("scan imageRows: %w", err)
				}

				product.Images = append(product.Images, img)
				Products = append(Products, product)
				productCategory.Product = Products
				ProductCategory = productCategory
			}

		}
	}

	return ProductCategory, nil
}
func (m *Repository) CreateProductCategory(ctx context.Context, productCategory domain.ProductCategory) (domain.ProductCategory, error) {
	q := fmt.Sprintf(`
INSERT INTO product_category(name)
VALUES ($1) RETURNING id`)

	err := m.db.QueryRowContext(ctx, q, productCategory.Name).Scan(&productCategory.ID)
	if err != nil {
		return domain.ProductCategory{}, fmt.Errorf("scan: %w", err)
	}

	return productCategory, nil
}
func (m *Repository) DeleteProductCategory(ctx context.Context, id string) error {
	q := fmt.Sprintf(`
DELETE FROM product_category
WHERE id = $1`)

	_, err := m.db.ExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
