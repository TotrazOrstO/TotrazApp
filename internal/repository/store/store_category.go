package store

import (
	"context"
	"fmt"
	"totraz_store/internal/domain"
)

func (r *StoreRepository) AllStoreCategory(ctx context.Context, limit int, offset int) ([]domain.StoreCategory, error) {
	q := fmt.Sprintf(`
SELECT id, name
FROM store_category
LIMIT $1
OFFSET $2`)

	rows, err := r.db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	var StoreCategory []domain.StoreCategory
	for rows.Next() {
		var storeCategory domain.StoreCategory
		err := rows.Scan(
			&storeCategory.ID,
			&storeCategory.Name,
			//&productCategory.Product,
		)

		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		q = `SELECT stc.store_category_id, stc.store_id
FROM store_to_category stc JOIN store_category sc
ON stc.store_category_id = sc.id`

		storeRows, err := r.db.QueryContext(ctx, q)
		if err != nil {
			return nil, fmt.Errorf("query: %w", err)
		}
		defer storeRows.Close()

		var Stores []domain.Store
		for storeRows.Next() {
			var store domain.Store
			err := storeRows.Scan(
				&store.ID,
				&store.Name,
			)

			if err != nil {
				return nil, fmt.Errorf("category scan: %w", err)
			}

			q = fmt.Sprintf(`
SELECT i.id img_id, i.file_name img_name, i.file_ext img_ext, i.file_body img_body
FROM image_to_store its
         JOIN images i on i.id = its.image_id
WHERE its.store_id = $1
`)

			imageRows, err := r.db.QueryContext(ctx, q, store.ID)
			if err != nil {
				return nil, fmt.Errorf("get images: %w", err)
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

				store.Images = append(store.Images, img)
				Stores = append(Stores, store)
				storeCategory.Stores = Stores
			}

		}
		StoreCategory = append(StoreCategory, storeCategory)
	}

	return StoreCategory, nil
}
func (r *StoreRepository) StoreCategoryById(ctx context.Context, id string) (domain.StoreCategory, error) {
	q := fmt.Sprintf(`
SELECT id
FROM store_category sc
WHERE sc.id = $1`)

	rows, err := r.db.QueryContext(ctx, q, id)
	if err != nil {
		return domain.StoreCategory{}, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	var StoreCategory domain.StoreCategory
	for rows.Next() {
		var storeCategory domain.StoreCategory
		err := rows.Scan(
			&storeCategory.ID,
			&storeCategory.Name,
		)

		if err != nil {
			return domain.StoreCategory{}, fmt.Errorf("scan imageRows: %w", err)
		}

		q = `SELECT stc.store_category_id, stc.store_id
FROM store_to_category stc JOIN store_category sc
ON stc.store_category_id = sc.id`

		storeRows, err := r.db.QueryContext(ctx, q)
		if err != nil {
			return domain.StoreCategory{}, fmt.Errorf("query: %w", err)
		}
		defer storeRows.Close()

		var Stores []domain.Store
		for storeRows.Next() {
			var store domain.Store
			err := storeRows.Scan(
				&store.ID,
				&store.Name,
			)

			if err != nil {
				return domain.StoreCategory{}, fmt.Errorf("category scan: %w", err)
			}

			q = fmt.Sprintf(`
SELECT i.id img_id, i.file_name img_name, i.file_ext img_ext, i.file_body img_body
FROM image_to_store its
         JOIN images i on i.id = its.image_id
WHERE its.store_id = $1
`)

			imageRows, err := r.db.QueryContext(ctx, q, store.ID)
			if err != nil {
				return domain.StoreCategory{}, fmt.Errorf("get images: %w", err)
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
					return domain.StoreCategory{}, fmt.Errorf("scan imageRows: %w", err)
				}

				store.Images = append(store.Images, img)
				Stores = append(Stores, store)
				storeCategory.Stores = Stores
				StoreCategory = storeCategory
			}

		}
	}

	return StoreCategory, nil
}
func (r *StoreRepository) CreateStoreCategory(ctx context.Context, storeCategory domain.StoreCategory) (domain.StoreCategory, error) {
	q := fmt.Sprintf(`
INSERT INTO store_category(name)
VALUES ($1) RETURNING id`)

	err := r.db.QueryRowContext(ctx, q, storeCategory.Name).Scan(&storeCategory.ID)
	if err != nil {
		return domain.StoreCategory{}, fmt.Errorf("scan: %w", err)
	}

	return storeCategory, nil
}
func (r *StoreRepository) DeleteStoreCategory(ctx context.Context, id string) error {
	q := fmt.Sprintf(`
DELETE FROM store_category
WHERE id = $1`)

	_, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
