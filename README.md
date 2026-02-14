# 🎵 Go Albums REST API

A simple RESTful API built with Go using the Gin framework and MySQL.
This project demonstrates basic CRUD operations, database integration, and REST API design.

---

## 🚀 Features

* Get all albums
* Get album by ID
* Create a new album
* Delete an album
* JSON request/response handling
* MySQL database integration

---

## 🛠 Tech Stack

* Go
* Gin Web Framework
* MySQL
* database/sql

---

## 📂 Project Structure

```
.
├── main.go
├── go.mod
├── go.sum
└── README.md
```

---

## ⚙️ Setup & Installation

### 1️⃣ Clone the repository

```bash
git clone https://github.com/YOUR_USERNAME/go-albums-api.git
cd go-albums-api
```

### 2️⃣ Install dependencies

```bash
go mod tidy
```

### 3️⃣ Setup MySQL Database

Create a database named:

```
recordings
```

Then create the `albums` table:

```sql
CREATE TABLE albums (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    artist VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL
);
```

---

### 4️⃣Run the Server

```bash
go run main.go
```

Server runs at:

```
http://localhost:8080
```

---

## 📌 API Endpoints

### 🔹 Get All Albums

```
GET /albums
```

Example:

```bash
curl http://localhost:8080/albums
```

---

### 🔹 Get Album by ID

```
GET /albums/:id
```

Example:

```bash
curl http://localhost:8080/albums/1
```

---

### 🔹 Create Album

```
POST /albums
```

Example:

```bash
curl -X POST http://localhost:8080/albums \
-H "Content-Type: application/json" \
-d '{
  "title": "Blue Train",
  "artist": "John Coltrane",
  "price": 56.99
}'
```

---

### 🔹 Delete Album

```
DELETE /albums/:id
```

Example:

```bash
curl -X DELETE http://localhost:8080/albums/1
```

---

## 🧠 What This Project Demonstrates

* REST API design
* CRUD operations
* Database connectivity with Go
* SQL queries & prepared statements
* JSON validation
* HTTP status codes
* Basic input validation

---

## 📈 Future Improvements

* Add PUT/PATCH update endpoint
* Add request logging middleware
* Add unit tests
* Dockerize the application
* Use environment configuration package
* Add authentication (JWT)
* Implement repository pattern

---

## 📜 License

This project is open-source and available under the MIT License.

---

## 👤 Author

Omar Elshaikh

GitHub: [https://github.com/OmarMohamedElsheikh](https://github.com/OmarMohamedElsheikh)

