# Test-Trafilea-API

API for creating a cart, adding products to the cart, getting cart data, updating the quantity of a product in a cart, and creating an order from the cart

## API test

1. Clone or download the repository

2. Open the terminal and install the project dependencies with 'go mod tidy' to be able to test the endpoints.

3. Run the 'go run cmd/main.go' command to start the server, you should expect a message like this:
    ""Welcome to my API
    Server running on the port 3000""

4. Use a tool that allows us to test the endpoints with their respective methods, such as POSTMAN or INSOMNIA.



##  API Endpoints

1.   Create a cart

#### POST /api/v1/create_cart

Request body example

~~~
{
    "CartID": "Cart1",
    "UserID": "User1",
    "Products": [
        {
            "ProductID": "1",
            "Name": "coffee",
            "Category": "Coffee",
            "Price": 4,
            "Quantity": 2
        }
    ]
}
~~~

Example of request response with status code 200

~~~
{
    "CartID": "Cart1",
    "UserID": "User1",
    "Products": [
        {
            "ProductID": "1",
            "Name": "coffee",
            "Category": "Coffee",
            "Price": 4,
            "Quantity": 2
        }
    ]
}
~~~


2. Get cart 

#### GET /api/v1/cart

Example of request response with status code 200
~~~
{
    "CartID": "Cart1",
    "UserID": "User1",
    "Products": [
        {
            "ProductID": "1",
            "Name": "coffee",
            "Category": "Coffee",
            "Price": 4,
            "Quantity": 2
        }
    ]
}
~~~

3. Add product to cart

#### POST /api/v1/add_product_cart

Request body example

~~~
{
    "ProductID": "2",
    "Name": "Table",
    "Category": "Equipment",
    "Price": 20,
    "Quantity": 2
}
~~~

Example of request response with status code 200

~~~
{
    "ProductID": "2",
    "Name": "Table",
    "Category": "Equipment",
    "Price": 20,
    "Quantity": 2
}
~~~

4. Update the quantity of a specific product

#### PUT /api/v1/update_quantity_product

Request body example

~~~
{
    "ProductID": "2",
    "Name": "Table",
    "Category": "Equipment",
    "Price": 40,
    "Quantity": 4
}
~~~

Example of request response with status code 200

~~~
{
    "ProductID": "2",
    "Name": "Table",
    "Category": "Equipment",
    "Price": 40,
    "Quantity": 4
}
~~~

5. Create cart order

#### POST /api/v1/create_order

Request body example

~~~
{
    "CartID": "Cart1"
}
~~~

Example of request response with status code 200

~~~
{
    "CartID": "Cart1",
    "Totals": {
        "Products": [
            {
                "ProductID": "1",
                "Name": "coffee",
                "Category": "Coffee",
                "Price": 4,
                "Quantity": 2
            },
            {
                "ProductID": "2",
                "Name": "Table",
                "Category": "Equipment",
                "Price": 40,
                "Quantity": 4
            },
            {
                "ProductID": "CoffeeFree1",
                "Name": "coffee",
                "Category": "Coffee",
                "Price": 0,
                "Quantity": 1
            }
        ],
        "Amount_shipping": 0,
        "Percent_discount": 0,
        "Amount_total_order": 44
    }
}
~~~
