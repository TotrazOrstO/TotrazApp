CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE product (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL
);

CREATE TABLE store (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL
);

CREATE TABLE images (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    file_name TEXT,
    file_ext VARCHAR(50),
    file_body BYTEA NOT NULL
);

CREATE TABLE image_to_product (
    image_id uuid REFERENCES images (id) ON DELETE CASCADE,
    product_id uuid REFERENCES product (id) ON DELETE CASCADE
);

CREATE TABLE image_to_store (
    image_id uuid REFERENCES images (id) ON DELETE CASCADE,
    store_id uuid REFERENCES store (id) ON DELETE CASCADE
);

CREATE TABLE store_category (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL
);

CREATE TABLE product_category (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL
);

CREATE TABLE store_to_category (
                                   store_category_id uuid REFERENCES store_category (id) ON DELETE CASCADE,
                                   store_id uuid REFERENCES store (id) ON DELETE CASCADE
);

CREATE TABLE product_to_category (
                                     product_category_id uuid REFERENCES product_category (id) ON DELETE CASCADE,
                                     product_id uuid REFERENCES product (id) ON DELETE CASCADE
);
