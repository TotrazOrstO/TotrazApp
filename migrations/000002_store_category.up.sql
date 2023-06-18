CREATE TABLE product_category_to_store (
                                           product_category_id uuid REFERENCES product_category (id) ON DELETE CASCADE,
                                           store_id uuid REFERENCES store (id) ON DELETE CASCADE
);