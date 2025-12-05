# Go Base Project

A base project template for building RESTful APIs in Go using Clean Architecture.

## Features

-   **Clean Architecture**: `cmd`, `internal`, `pkg` structure.
-   **Routing**: `Gin` HTTP framework.
-   **Database**: `GORM` with support for `MySQL` and `Postgres`.
-   **Authentication**: JWT-based authentication (`login`/`register`).
-   **Configuration**: Environment-based configuration using `.env`.
-   **Logging**: Structured logging with `Zerolog`.
-   **Live Reload**: Hot-reload with `Air`.
-   **Standardized Response**: Consistent JSON response format.

## Getting Started

1.  **Clone and extract** this project.
2.  **Navigate to the project directory**:
    ```bash
    cd go-base-project
    ```
3.  **Install dependencies**:
    ```bash
    go mod download
    ```
4.  **Setup your environment**:
    -   Copy `.env.example` to `.env`.
    -   Update the `.env` file with your database credentials and JWT secret.
5.  **Run the server** (with Air for hot-reload):
    ```bash
    air
    ```
    Or run it manually:
    ```bash
    go run cmd/api/main.go
    ```

## API Endpoints

-   `GET /api/v1/ping`: Health check endpoint.
-   `POST /api/v1/register`: Register a new user.
-   `POST /api/v1/login`: Log in a user and get a JWT token.
-   `GET /api/v1/profile`: Get the authenticated user's profile (protected).

## Project Structure

