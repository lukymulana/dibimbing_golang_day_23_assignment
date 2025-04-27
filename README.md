# dibimbing_golang_day_23_assignment (Golang Simple API Inventory Management)

## Project setup

Clone the project

```bash
  git clone https://github.com/lukymulana/dibimbing_golang_day_23_assignment.git
```

Go to the project directory

```bash
  cd dibimbing_golang_day_23_assignment
```

Install dependencies

```bash
  go mod tidy
```
## Setup Database

Create Database
```bash
    CREATE DATABASE golang_inventory_management;
```

Run migration 
```bash
    mysql -u root golang_inventory_management < migrations/migrations.sql
```


## Run Project

```bash
  go run main.go
```


## Running Test

### Products

Create a product

```bash
  curl -X POST http://localhost:8081/products -H "Content-Type: application/json" -d '{"Name": "Product C", "Description": "Description C", "Price": 300, "Category": "Category 3"}'
```

Get all products

```bash
    curl http://localhost:8081/products
```

Get product by ID

```bash
    curl http://localhost:8081/products/1
```

Update a product

```bash
    curl -X PUT http://localhost:8081/products/1 -H "Content-Type: application/json" -d '{"Name": "Updated Product A", "Description": "Updated Description A", "Price": 150, "Category": "Category 1"}'
```

Delete a product

```bash
    curl http://localhost:8081/inventory/1
```

### Inventory

Get inventory by product ID

```bash
curl http://localhost:8081/inventory/1
```

Update inventory

```bash
curl -X PUT http://localhost:8081/inventory/1 -H "Content-Type: application/json" -d '{"Quantity": 60, "Location": "Warehouse A"}'
```

### Orders

Create an order

```bash
curl -X POST http://localhost:8081/orders -H "Content-Type: application/json" -d '{"ProductID": 2, "Quantity": 10, "OrderDate": "2023-10-03"}'
```

Get order by ID 

```bash
curl http://localhost:8081/orders/1
```

