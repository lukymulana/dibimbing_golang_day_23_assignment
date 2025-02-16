-- Sample data insertion
INSERT INTO products (name, description, price, category) VALUES
('Product A', 'Description A', 100, 'Category 1'),
('Product B', 'Description B', 200, 'Category 2');

INSERT INTO inventories (product_id, quantity, location) VALUES
(1, 50, 'Warehouse A'),
(2, 30, 'Warehouse B');

INSERT INTO orders (product_id, quantity, order_date) VALUES
(1, 10, '2023-10-01'),
(2, 5, '2023-10-02');