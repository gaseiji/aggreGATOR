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


## Configuration

Before running the application, you need to create a `.gatorconfig.json` file in the root directory of the project. This file should contain the following structure:

```json
{
  "db_url": "postgres://<username>:<password>@localhost:<port>/<dbname>?sslmode=disable",
  "current_user_name": "<username>"
}
