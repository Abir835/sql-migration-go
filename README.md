# Go Project with Database Migrations

This project demonstrates how to implement database migrations in a Go application using the `sql-migrate` library. It connects to a PostgreSQL database, applies migrations, and creates the necessary tables.

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Database Migrations](#database-migrations)
- [License](#license)

## Introduction

This project shows how to set up and run database migrations in a Go application. It uses the `sql-migrate` package to handle migrations for a PostgreSQL database. The migration process will automatically apply changes to your database schema, such as creating tables or modifying their structure.

## Features

- Connects to a PostgreSQL database
- Applies database migrations using `sql-migrate`
- Creates a `users` table with an `id` and `name` column
- Supports both `Up` and `Down` migrations to apply or rollback changes

## Installation

Follow these steps to set up and run the project locally:

### Prerequisites
- Go (1.x or later)
- PostgreSQL running locally or on a server
- `gopkg.in/rubenv/sql-migrate.v0` package for migrations

### Steps
1. Clone the repository:
    ```bash
    git clone https://github.com/your-username/project-name.git
    cd project-name
    ```
2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Set up your PostgreSQL database by replacing the connection string in the `ConnectDB` function:
    ```go
    db, err = sql.Open("postgres", "postgres://username:password@localhost:5432/dbname?sslmode=disable")
    ```
    - Replace `username`, `password`, `localhost`, and `dbname` with your actual PostgreSQL credentials.

4. Create the migrations folder and add your migration SQL file (e.g., `migrations/0001_create_users_table.up.sql`):
    - **Up Migration (`0001_create_users_table.up.sql`)**
        ```sql
        CREATE TABLE users (
           id SERIAL PRIMARY KEY,
           name TEXT NOT NULL
        );
        ```

    - **Down Migration (`0001_create_users_table.down.sql`)**
        ```sql
        DROP TABLE users;
        ```

5. Run the application to apply migrations:
    ```bash
    go run main.go
    ```

## Usage

The program will connect to the PostgreSQL database and apply migrations. If you run the application successfully, it will output a message indicating how many migrations were applied.

```bash
  Applied 1 migrations!
```