# Event Booking API (Go + Gin)

REST API for managing events and user registrations with JWT-based authentication.

The project allows users to:
- register and log in
- create and manage events
- register for events
- cancel registrations

Built as a backend learning project to practice Go web development and REST API design.

## Tech Stack

- Go
- Gin Web Framework
- SQLite
- JWT Authentication

## Features

- User registration and login
- JWT authentication
- Create, update, delete events
- Register for events
- Cancel registration for event

## Project Structure

```
.
├── db/            # Database initialization and queries
├── handlers/      # HTTP handlers
├── middlewares/   # Auth middleware
├── models/        # Data models
├── routes/        # API routes
├── utils/         # Helpers (JWT, hashing, etc.)
├── main.go
└── go.mod
```

## How to Run

1. Clone repository
```
git clone https://github.com/y-udov/event-booking.git
cd event-booking
git checkout dev
```
2. Install dependencies
```
go mod tidy
```
3. Run server
```
go run main.go
```

Server will start on:
```
http://localhost:8080
```

SQLite database file will be created automatically on first run.

## Authentication Flow
1. Register

`POST /signup`

```
{
  "email": "test@example.com",
  "password": "password123"
}
```

2. Login

`POST /login`

```
{
  "email": "test@example.com",
  "password": "password123"
}
```

Response contains JWT token:

```
{
  "token": "YOUR_JWT_TOKEN"
}
```

Use this token in Authorization header for protected endpoints:

Authorization: `YOUR_JWT_TOKEN`

## API Endpoints

**Public**

| Method | Endpoint     | Description       |
|--------|-------------|-----------------|
| POST   | /signup     | User registration |
| POST   | /login      | User login        |
| GET    | /events     | Get all events    |
| GET    | /events/:id | Get event by ID   |


**Protected (JWT required)**

| Method | Endpoint                | Description            |
|--------|------------------------|-----------------------|
| POST   | /events                | Create event           |
| PUT    | /events/:id            | Update event           |
| DELETE | /events/:id            | Delete event           |
| POST   | /events/:id/register   | Register for event     |
| DELETE | /events/:id/register   | Cancel registration    |

## API Testing

Project includes HTTP request examples for testing the API (e.g. via VS Code REST Client or Postman).

Typical flow:

1. Register user
2. Login and copy JWT token
3. Paste token into Authorization header
4. Call  the protected endpoints

## Notes

- This project is intended as a learning and portfolio project:
- no frontend
- no production deployment
- focuses on backend API and authentication logic
