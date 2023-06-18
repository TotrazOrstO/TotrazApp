package store

import (
	"context"
	"fmt"
	"totraz_store/internal/domain"
)

func (r *StoreRepository) AllStores(ctx context.Context, limit int, offset int) ([]domain.Store, error) {
	q := fmt.Sprintf(`
SELECT id, name
FROM store
LIMIT $1
OFFSET $2`)

	rows, err := r.db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	var stores []domain.Store
	for rows.Next() {
		var store domain.Store
		err := rows.Scan(
			&store.ID,
			&store.Name,
		)

		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		q = fmt.Sprintf(`
SELECT i.id img_id, i.file_name img_name, i.file_ext img_ext, i.file_body img_body
FROM image_to_store its
         JOIN images i on i.id = its.image_id
WHERE its.store_id = $1`)

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
		}

		q = `SELECT sc.id, sc.name
FROM store_to_category stc
JOIN store_category sc on stc.store_category_id = sc.id
WHERE stc.store_id=$1`

		categoryRows, err := r.db.QueryContext(ctx, q, store.ID)
		if err != nil {
			return nil, fmt.Errorf("query: %w", err)
		}
		defer categoryRows.Close()

		for categoryRows.Next() {
			var category domain.StoreCategory
			err := rows.Scan(
				&category.ID,
				&category.Name,
			)

			if err != nil {
				return nil, fmt.Errorf("category scan: %w", err)
			}
			store.StoreCategories = append(store.StoreCategories, category)
		}

		stores = append(stores, store)
	}

	return stores, nil
}

func (r *StoreRepository) StoreById(ctx context.Context, id string) (domain.Store, error) {

	q := fmt.Sprintf(`
SELECT *
FROM store
WHERE id = $1`)

	rows, err := r.db.QueryContext(ctx, q, id)
	if err != nil {
		return domain.Store{}, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	var Store domain.Store
	for rows.Next() {
		var store domain.Store
		err := rows.Scan(
			&store.ID,
			&store.Name,
			&store.Images,
			&store.StoreCategories,
		)

		if err != nil {
			return domain.Store{}, fmt.Errorf("scan: %w", err)
		}
		Store.ID = store.ID
		Store.Name = store.Name
		Store.Images = store.Images
		Store.StoreCategories = store.StoreCategories
	}

	return Store, nil
}

//	q := fmt.Sprintf(`
//
// SELECT id, name
// FROM store
// WHERE id = $1`)
//
//	var store domain.Store
//
//	err := r.db.QueryRowContext(ctx, q, id).Scan(
//		&store.ID,
//		&store.Name,
//	)
//
//	if err != nil {
//		return domain.Store{}, fmt.Errorf("scan: %w", err)
//	}
//
//	q = fmt.Sprintf(`SELECT i.id img_id, i.file_name img_name, i.file_ext img_ext, i.file_body img_body
//
// FROM image_to_store its
//
//	JOIN images i on i.id = its.image_id
//
// WHERE its.store_id = $1
// `)
//
//	imageRows, err := r.db.QueryContext(ctx, q, store.ID)
//	if err != nil {
//		return domain.Store{}, fmt.Errorf("get images: %w", err)
//	}
//	defer imageRows.Close()
//
//	for imageRows.Next() {
//		var img domain.Image
//		err := imageRows.Scan(
//			&img.Id,
//			&img.Name,
//			&img.Ext,
//			&img.Body,
//		)
//
//		if err != nil {
//			return domain.Store{}, fmt.Errorf("scan imageRows: %w", err)
//		}
//
//		store.Images = append(store.Images, img)
//	}
//
//	q = `SELECT sc.id, sc.name
//
// FROM store_to_category stc
// JOIN store_category sc on stc.store_category_id = sc.id
// WHERE stc.store_id=$1`
//
//		categoryRows, err := r.db.QueryContext(ctx, q, store.ID)
//		if err != nil {
//			return domain.Store{}, fmt.Errorf("query: %w", err)
//		}
//		defer categoryRows.Close()
//
//		for categoryRows.Next() {
//			var category domain.StoreCategory
//			err := categoryRows.Scan(
//				&category.ID,
//				&category.Name,
//			)
//
//			if err != nil {
//				return domain.Store{}, fmt.Errorf("category scan: %w", err)
//			}
//			store.StoreCategories = append(store.StoreCategories, category)
//		}
//
//		return store, nil
//	}
func (r *StoreRepository) Create(ctx context.Context, store domain.Store) (domain.Store, error) {
	q := fmt.Sprintf(`
INSERT INTO store(name)
VALUES ($1) RETURNING id, name`)

	var newStore domain.Store
	err := r.db.QueryRowContext(ctx, q, store.Name).Scan(&newStore.ID, &newStore.Name)
	if err != nil {
		return domain.Store{}, fmt.Errorf("query: %w", err)
	}

	return newStore, nil
}

func (r *StoreRepository) AddImagesToStore(ctx context.Context, storeId string, imageId string) error {
	q := fmt.Sprintf(`
INSERT INTO image_to_store(image_id, store_id)
VALUES ($1, $2)
`)

	_, err := r.db.ExecContext(ctx, q, imageId, storeId)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	return nil
}

func (r *StoreRepository) AddStoreToCategory(ctx context.Context, storeCategoryId string, storeId string) error {
	q := fmt.Sprintf(`
INSERT INTO store_to_category(store_category_id, store_id)
VALUES ($1, $2) 
`)

	_, err := r.db.ExecContext(ctx, q, storeCategoryId, storeId)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	return nil
}

func (r *StoreRepository) AddProductCategoryToStore(ctx context.Context, productCategoryId string, storeId string) error {
	q := fmt.Sprintf(`
INSERT INTO product_category_to_store(product_category_id, store_id)
VALUES ($1, $2)
`)

	_, err := r.db.ExecContext(ctx, q, productCategoryId, storeId)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	return nil
}

func (r *StoreRepository) Delete(ctx context.Context, id string) error {
	q := fmt.Sprintf(`
DELETE FROM store
WHERE id = $1`)

	_, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
