# Telkomsel Technical Test API

Telkomsel's technical Assesment API implementation 
The app runs in the port :3000
Don't forget to bind your desired Port to the docker image when initializing it.

## Features

- **Create Product**: Adding a new product to the database.
- **Fetch Products**: Retrieve a list of products, can be filtered with brand & variety as well as product_id for single product.
- **Update Product**: Modify an existing product detail by pproduct_id.
- **Delete Product**: Remove a product from the database by product_id.

## Technologies

- **Go**: Programming language
- **Fiber**: Web framework for building REST APIs
- **SQLx**: Go package for working with databases
- **PostgreSQL**: Database management system

## How to Use the APIs

- **Create Product**:
    - Method: POST
    - URL: localhost:3000/v1/product/create
    - Body: Product object
```
curl --location 'localhost:3000/v1/product/create' \
--header 'Content-Type: application/json' \
--data '{
    "product_name": "Royko kaldu ayam",
    "product_description": "Essence kaldu ayam instant cocok untuk masak di rumah atau menjadi bumbu",
    "product_price": 15000,
    "product_variety": "Food",
    "product_rating": 4.5,
    "product_stock": 150,
    "product_url": "",
    "product_brand": "Mayora"

}'
```
- **Fetch Products**: 
    - Method: GET
    - URL: localhost:3000/v1/product/fetch
    - Query: [brand], [variety], [id]
```
curl --location 'localhost:3000/v1/product/fetch?id=&brand=samsung&variety=electronics'
```
- **Update Product**: Modify an existing product's details.
    - Method: PUT
    - URL: localhost:3000/v1/product/update
    - Body: Product Object
```
curl --location --request PUT 'localhost:3000/v1/product/update' \
--header 'Content-Type: application/json' \
--data '{
    "product_id": "bcad92da-f783-4e20-97fd-6fdfc421c48e",
    "product_name": "Royko kaldu ayam 250gr",
    "product_description": "Essence kaldu ayam instant cocok untuk masak di rumah atau menjadi bumbu",
    "product_price": 25000,
    "product_variety": "Food",
    "product_rating": 4.9,
    "product_stock": 150,
    "product_url": "testestestest.com",
    "product_brand": "Mayora"
}'
```
- **Delete Product**:
    - Method: DELETE
    - URL: localhost:3000/v1/product/delete
    - Params: product_id
```
curl --location --request DELETE 'localhost:3000/v1/product/delete/ef3f1280-a790-4a30-80fb-85a74f7159b7'
```
