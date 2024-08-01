Food Item Management System
Overview
This project is a Food Item Management System built using Go (Golang) with Fiber as the web framework and PostgreSQL as the database. It allows users to create and manage food items along with their associated ingredients and groceries.

Features
Create Food Items: Add new food items with their details and associated ingredients.
Manage Ingredients: Add and manage ingredients required for food items.
Manage Groceries: Add and manage groceries that can be used as ingredients.
Place Orders: Place orders for food items and manage associated inventory.
Installation
Prerequisites
Go (Golang) 1.18 or higher
PostgreSQL
Fiber (Go web framework)
GORM (Go ORM library)
Setup
Clone the Repository

bash
Copy code
git clone https://github.com/Frhnmj2004/go-restaurant-admin.git
cd yourrepository
Install Dependencies

Install the necessary Go modules:

bash
Copy code
go mod tidy
Configure Database

Update the config.yaml or equivalent configuration file with your PostgreSQL database credentials.

Example:

yaml
Copy code
database:
  host: localhost
  port: 5432
  user: yourusername
  password: yourpassword
  name: yourdatabase
Run Migrations

Ensure your database schema is up-to-date by running the migrations. You may use a migration tool or execute schema SQL manually.

Start the Server

Run the server using:

bash
Copy code
go run main.go
By default, the server will start on http://localhost:3000.

API Endpoints
Food Items
Create Food Item

POST /food-items
Request Body: JSON object with name, price, and ingredients.
Get Food Item

GET /food-items/:id
Response: JSON object with food item details and associated ingredients.
Ingredients
Create Ingredient

POST /ingredients
Request Body: JSON object with foodItemID, groceryID, and quantity.
Groceries
Create Grocery

POST /groceries
Request Body: JSON object with name.
Orders
Place Order

POST /orders
Request Body: JSON object with foodItemID and quantity.
Example JSON
Create Food Item Request
json
Copy code
{
  "name": "Green Coffee",
  "price": 10,
  "ingredients": [
    {
      "groceryName": "Coffee Beans",
      "quantity": 5
    },
    {
      "groceryName": "Sugar",
      "quantity": 2
    }
  ]
}
Create Food Item Response
json
Copy code
{
  "data": {
    "ID": 1,
    "CreatedAt": "2024-08-01T22:50:35.0270088+05:30",
    "UpdatedAt": "2024-08-01T22:50:35.0408024+05:30",
    "DeletedAt": null,
    "name": "Green Coffee",
    "price": 10,
    "ingredients": [
      {
        "ID": 211,
        "CreatedAt": "2024-08-01T22:50:35.0299905+05:30",
        "UpdatedAt": "2024-08-01T22:50:35.0299905+05:30",
        "DeletedAt": null,
        "fooditemid": 1,
        "groceryid": 2,
        "quantity": 5
      }
    ]
  },
  "message": "Food item created successfully"
}
Contributing
If you want to contribute to this project, please fork the repository and submit a pull request with your changes. Make sure to follow the existing code style and include tests for your changes.

License
This project is licensed under the MIT License. See the LICENSE file for details.

Contact
For any questions or issues, please contact jeejofarhan@gmail.com.