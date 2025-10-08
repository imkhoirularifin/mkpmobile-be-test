# MKP Mobile Backend Test

## Overview

### High Level Architecture Diagram

![High Level Architecture Diagram](./diagram/high-level-arch.png)

### Booking Flow Diagram

![Booking Flow Diagram](./diagram/booking-flow.png)

### Refund Flow Diagram

![Refund Flow Diagram](./diagram/refund-flow.png)

### DB Schema Diagram

![DB Schema Diagram](./diagram/db-schema.png)

> **Note:**  
> The database schema in this repository is intentionally simplified for demonstration purposes.  
> In a real-world application, the schema would be more complex to support additional features and requirements.

## Environment Variables

Create a `.env` file in the root directory and add the following environment variables:

| Variable Name       | Description                       | Example Value                              |
| ------------------- | --------------------------------- | ------------------------------------------ |
| GOOSE_DRIVER        | Database driver (e.g., postgres)  | postgres                                   |
| GOOSE_DBSTRING      | Database connection string        | postgres://user:password@host:port/db-name |
| GOOSE_MIGRATION_DIR | Directory for database migrations | ./migrations                               |

## Running the Application

1. Install dependencies:

   ```bash
   go mod download
   ```

2. Migrate the database with [Goose](https://pressly.github.io/goose/)

   ```bash
   goose up
   ```

3. Run the application:

   ```bash
    go run main.go
   ```

   or if you have [`air`](https://github.com/air-verse/air) installed:

   ```
   air
   ```

4. The application docs will be accessible at `http://localhost:3000/api/v1/docs` (or the port you specified).
