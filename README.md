# gin-user-api
This is a RESTful User Management API built with Golang and Gin also featuring CRUD operations and SQLite database integration using GORM.
It allows you to create, view, and delete users.
All the data is stored using SQLite and GORM.

Features

-  Create a user
-  Get all users
-  Get a user by ID
-  Delete a user
-  Store data in a database


Tech Stack

* Go (Golang)
* Gin
* GORM
* SQLite


Project Structure

ginbasics/
├── main.go
├── database/
├── handlers/
├── models/
└── routes/


 Setup

1. Clone the repository
git clone https://github.com/your-username/go-user-api.git


2. Go into the project folder
cd go-user-api

3. Install dependencies
go mod tidy

4. Run the server
go run main.go

The server will run on:
http://localhost:8080


API Endpoints

Get all users - GET /users


Get user by ID - GET /users/:id


Create a user - POST /users



Request body:

{
  "username": "Hannah",
  "email": "hannah@email.com"
}

To Delete a user

DELETE /users/:id

 
 NB:
* Data is stored in a local SQLite file (`users.db`)
* The database file is created automatically when the app run

  Other steps:
* Add update user endpoint
* Add login and authentication
* Connect to a frontend app



Hannah Bamgbelu Gbemisola
