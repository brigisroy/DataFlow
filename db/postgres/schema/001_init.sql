-- +goose Up
-- +goose StatementBegin
-- Create schema
CREATE SCHEMA dataflow;

-- Create orders table
CREATE TABLE dataflow.orders (
    id SERIAL PRIMARY KEY, 
    order_id INT NOT NULL,
    date_of_sale DATE NOT NULL,
    region VARCHAR(100) NOT NULL,
    discount NUMERIC(5, 2) DEFAULT 0.00,
    shipping_cost NUMERIC(10, 2) NOT NULL,
    payment_method VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_order_id UNIQUE (order_id)  -- Ensure order_id is unique
);

-- Create products table
CREATE TABLE dataflow.products (
      id SERIAL PRIMARY KEY, 
      product_id VARCHAR(50) NOT NULL,
      unit_price NUMERIC(10, 2) NOT NULL,
      product_name VARCHAR(255) NOT NULL,
      category VARCHAR(100) NOT NULL,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      CONSTRAINT unique_product_id UNIQUE (product_id)  -- Ensure product_id is unique
);

-- Create customers table
CREATE TABLE dataflow.customers (
       id SERIAL PRIMARY KEY, 
       customer_id VARCHAR(50) NOT NULL,
       name VARCHAR(255) NOT NULL,
       email VARCHAR(255) UNIQUE NOT NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       address TEXT NOT NULL,
       CONSTRAINT unique_customer_id UNIQUE (customer_id)  -- Ensure customer_id is unique
);

-- Create order_mapping table
CREATE TABLE dataflow.order_mapping (
           id SERIAL PRIMARY KEY, 
           order_id INT REFERENCES dataflow.orders(order_id) ON DELETE CASCADE,
           product_id VARCHAR(50) REFERENCES dataflow.products(product_id) ON DELETE CASCADE,
           customer_id VARCHAR(50) REFERENCES dataflow.customers(customer_id) ON DELETE CASCADE,
           no_of_units_ordered INT NOT NULL,
           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
           CONSTRAINT unique_order_mapping UNIQUE (order_id, product_id, customer_id) -- Ensure unique mapping
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS dataflow.order_mapping;
DROP TYPE IF EXISTS dataflow.customers;
DROP TYPE IF EXISTS dataflow.products;
DROP TABLE IF EXISTS dataflow.orders ;
DROP SCHEMA dataflow;

-- +goose StatementEnd
