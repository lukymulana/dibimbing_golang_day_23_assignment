CREATE TABLE IF NOT EXISTS products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    category VARCHAR(100),
    image_path VARCHAR(255) -- add product image
);

CREATE TABLE IF NOT EXISTS inventory (
    product_id INT PRIMARY KEY,
    quantity INT NOT NULL,
    location VARCHAR(100),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE IF NOT EXISTS orders (
    order_id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT,
    quantity INT NOT NULL,
    order_date DATE,
    FOREIGN KEY (product_id) REFERENCES products(id)
);

INSERT INTO products (name, description, price, category) VALUES
('Product A', 'Description A', 100, 'Category 1'),
('Product B', 'Description B', 200, 'Category 2');

INSERT INTO inventories (product_id, quantity, location) VALUES
(1, 50, 'Warehouse A'),
(2, 30, 'Warehouse B');

INSERT INTO orders (product_id, quantity, order_date) VALUES
(1, 10, '2023-10-01'),
(2, 5, '2023-10-02');