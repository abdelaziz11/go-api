```maenv t (requires authentication)

## Testing

You can use any HTTP client like curl or Postman to test the API endpoints.

## License

This projecthe root` file inrkdown
 of the project:

```env
POSTGRES_URL=your_postgres_connection_string
```

## Running the Application

To run the application, use the following command:

```bash
go run main.go
```

The server will start on `localhost:8000`.

## API Endpoints

- POST `/login`: Login a user
- GET `/auth/user`: Get the authenticated user's details (requires authentication)
- PUT `/auth/user`: Update the authenticated user's details (requires authentication)
- DELETE `/auth/user`: Delete the authenticated user's account# GO Authentication API

This project is a simple authentication API built with Go, Echo framework, and PostgreSQL.

## Project Structure

```
.
├── main.go
├── config
│   └── config.go
├── models
│   └── user.go
├── database
│   └── db.go
├── routes
│   └── routes.go
├── controllers
│   └── authController.go
├── middlewares
│   └── authMiddleware.go
└── README.md
```

## Prerequisites

- Go 1.16 or later
- PostgreSQL

## Setup

1. Clone the repository:

```bash
git clone https://github.com/yourusername/yourprojectname.git
```

2. Navigate to the project directory:

```bash
cd yourprojectname
```

3. Install the dependencies:

```bash
go mod download
```

4. Set up your PostgreSQL database and add the connection string to a `.