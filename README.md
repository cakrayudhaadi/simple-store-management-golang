# Simple Store Management - API Documentation

## 📌 Overview

Simple Store Management is a REST API-based web application for managing stores that have branches in many places, record employees, items, and sales data. It can also see branches and employees who make the most sales within a month, year, or all time.

## 📖 API Documentation

The API documentation can be accessed at:

🔗 [API Documentation](https://simple-store-management-golang-production.up.railway.app/swagger/index.html#/)

## 🔑 Authentication

To use the API, you need to sign up or log in with an existing account.

### Test Account:

```json
{
  "username" : "admin",
  "password" : "password"
}
```

You can also register a new account through the API if needed.

## 🚀 Getting Started

### 1. Clone the Repository

```sh
git clone <repository_url>
cd <project_directory>
```

### 2. Install Dependencies

```sh
go mod tidy
```

### 3. Set Environment Variables

In main.go file, change **simple-store-management-golang-production.up.railway.app**
in line 18 to **localhost:8080**

And then init swagger to open swagger locally

```sh
swag init
```

Also change the database configuration in configs/config.json to use the database of your satisfy

### 4. Run the Server

```sh
go run main.go
```

### 5. Access the API

Once the server is running, open your browser and go to:

```
http://localhost:8080/swagger/index.html#/
```

## 🛠 Technologies Used

- Go (Golang)
- Gin Framework
- GORM
- PostgreSQL (Database)
- JWT (Authentication)
- Swaggo (Documentation API)
