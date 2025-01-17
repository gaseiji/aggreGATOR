# aggreGATOR CLI

aggreGATOR is a command-line tool designed to interact with your database, handle migrations, and perform other necessary tasks related to your project. This tool is written in Go and uses Postgres as the backend database.

> **Note**: This project is part of an exercise from Boot.dev.

## Requirements

Before you can run the program, you need to have the following installed:

- **Go** (1.23.3 or higher): [Install Go](https://go.dev/doc/install)
- **PostgreSQL**: [Install PostgreSQL](https://www.postgresql.org/download/)

## Installation

To install the aggreGATOR CLI, run the following command:

```bash
go install github.com/gaseiji/aggreGATOR@latest
```

## Database Migration

Before running the application, you need to migrate the database schemas using **Goose**.

1. **Install Goose**: First, make sure that Goose is installed. If it's not, you can install it globally by running the following command:

   ```bash
   go install github.com/pressly/goose/v3@latest
    ```
2. **Run Migrations**: Navigate to the /sql/schema folder in your project. This folder contains the 5 SQL schema files that need to be migrated. Use the following command to apply the migrations:

```bash
goose -dir ./sql/schema postgres "<your-database-url>" up
```
Replace `<your-database-url>` with the appropriate connection URL to your Postgres database.

3. **Generated Code**: The application uses sqlc to generate the database access code. Since this code is already generated and committed to the repository, you do not need to regenerate it. The generated code is located in the internal/database directory, and it will be used automatically when you run the application.

4. **Verify**: After running the migration, you should see that the database schemas have been successfully updated. You can verify the changes by checking the database or using any appropriate database management tool.

## Configuration

Before running the application, you need to create a `.gatorconfig.json` file in the root directory of the project. This file should contain the following structure:

```json
{
  "db_url": "postgres://<username>:<password>@localhost:<port>/<dbname>?sslmode=disable",
  "current_user_name": "<username>"
}
```

Replace the placeholders with your actual database connection details:

`<username>`: Your database username.
`<password>`: Your database password.
`<port>`: The port your PostgreSQL instance is running on (usually 5432).
`<dbname>`: The name of your database.

## Available Commands

Once the setup is complete, you can use the following commands:

- **login**: Log in to the application. Example usage: `aggregator login <username> <password>`
- **register**: Register a new user. Example usage: `aggregator register <username> <password>`
- **reset**: Reset the database to its initial state. Example usage: `aggregator reset`
- **users**: List all registered users. Example usage: `aggregator users`
- **agg**: Perform aggregation. Example usage: `aggregator agg`
- **addfeed**: Add a feed to the application (requires login). Example usage: `aggregator addfeed <feed_url>`
- **feeds**: List all the feeds (requires login). Example usage: `aggregator feeds`
- **follow**: Follow a feed (requires login). Example usage: `aggregator follow <feed_id>`
- **following**: View the feeds you're following (requires login). Example usage: `aggregator following`
- **unfollow**: Unfollow a feed (requires login). Example usage: `aggregator unfollow <feed_id>`
- **browser**: Open the browser to view the aggregated content (requires login). Example usage: `aggregator browser`

## Running a Command
Once installed, you can run commands through the CLI. For example, to log in with a username

```bash
aggregator login <username>
```
Replace `<username>` with your desired credentials.

