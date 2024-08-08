# Go Backend Services

## Overview

This project consists of two backend services:
- **Account Service**: Manages user accounts, payment accounts, and payment histories.
- **Payment Service**: Handles transactions such as send and withdraw operations.

The services are built using the following tech stack:
- **Golang**: Version 1.21
- **Gin**: Web framework for Go
- **Supabase**: Authentication
- **Prisma**: ORM for PostgreSQL
- **PostgreSQL**: Database
- **Nginx**: Reverse proxy
- **Docker**: Containerization
- **Docker Compose**: Orchestration

## Prerequisites

- **Docker**: Ensure Docker is installed on your system.
- **Docker Compose**: Ensure Docker Compose is installed on your system.
- **GVM**: Optional, if managing Go versions locally.

## Setup

### Install Dependencies

1. **Clone the repository:**

    ```sh
    git clone https://github.com/your-repo/go-backend-services.git
    cd go-backend-services
    ```

2. **Install Go Version Manager (GVM):**

    ```sh
    bash < <(curl -sSL https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
    source ~/.gvm/scripts/gvm
    ```

3. **Install Go 1.21:**

    ```sh
    gvm install go1.21
    gvm use go1.21 --default
    ```

### Configuration

1. **Create a `.env` file** in the root directory:

    ```env
    DATABASE_URL=postgres://postgres:password@postgres:5432/mydb?sslmode=disable
    SUPABASE_URL=<your_supabase_url>
    SUPABASE_KEY=<your_supabase_key>
    ```

### Docker Setup

1. **Build Docker images:**

    ```sh
    docker-compose build
    ```

2. **Start the services:**

    ```sh
    docker-compose up
    ```

    This command will start the PostgreSQL database, the account service, and the payment service.

### How to Run

1. **To start the services**, use Docker Compose:

    ```sh
    docker-compose up
    ```

2. **To stop the services**, use:

    ```sh
    docker-compose down
    ```

### API Endpoints

#### Account Service

- **Register/Login**

    Use Supabase authentication to register and log in users.

- **Send Transaction**

    ```http
    POST /send
    ```

    Request body:

    ```json
    {
      "amount": 100,
      "currency": "USD",
      "toAddress": "recipient_address",
      "accountId": 1
    }
    ```

- **Withdraw Transaction**

    ```http
    POST /withdraw
    ```

    Request body:

    ```json
    {
      "amount": 100,
      "currency": "USD",
      "toAddress": "recipient_address",
      "accountId": 1
    }
    ```

#### Payment Service

- **Get Accounts for a User**

    ```http
    GET /users/:userId/accounts
    ```

    Example:

    ```sh
    curl -X GET http://localhost:8080/users/1/accounts
    ```

- **Get Transactions for an Account**

    ```http
    GET /accounts/:accountId/transactions
    ```

    Example:

    ```sh
    curl -X GET http://localhost:8080/accounts/1/transactions
    ```

### Development

1. **Install Go modules:**

    ```sh
    go mod tidy
    ```

2. **Run tests:**

    ```sh
    go test ./...
    ```


