# Gator CLI

Gator is a command-line application that allows users to manage RSS feeds, fetch posts, and browse content directly in the terminal. It’s built with Go and PostgreSQL for simplicity and performance.

---

## Features

- **Follow RSS Feeds**: Add and follow feeds by URL.
- **Unfollow Feeds**: Stop following feeds you no longer need.
- **Browse Posts**: View the latest posts from the feeds you follow.
- **Continuous Aggregation**: Periodically fetch new posts from followed feeds.
- **User Accounts**: Each user can maintain their own feed subscriptions.

---

## Prerequisites

Ensure the following are installed:

1. **PostgreSQL**: Required for database management.
2. **Go (v1.20 or newer)**: Required to build and run the application.
3. **Goose**: For database migrations.
   ```bash
   go install github.com/pressly/goose/v3/cmd/goose@latest
4. **sqlc**: 
    ```bash
   go install github.com/kyleconroy/sqlc/cmd/sqlc@latest


---

## Installation

1. **Clone the Repository**
   ```bash
   git clone https://github.com/your-username/gator-cli.git
   cd gator-cli

2. **Install Gator Build and install the CLI tool**
    ```bash
    go install .

3. **Verify Installation Confirm the tool is installed by running**
    ```bash
    gator --help

## Configuration
    Create a config.json file in the root directory to connect Gator to your PostgreSQL database:

    ```json
    {  
    "db_url": "postgres://username:password@localhost:5432/gator",  
    "current_user_name": "default_user"  
    }  
    
    Replace username, password, and localhost with your database credentials.

---

## Database Setup

1. **Create a PostgreSQL database**
   ```bash
   createdb gator

2. **Run Migrations** Apply all required database migrations using Goose:

    ```bash
    goose up

3. **Generate SQL Queries** Use sqlc to generate Go code from the SQL queries:
    ```bash
    sqlc generate


---

## Usage

Here are some of the commands available in Gator:

**Add a Feed**
Add a new feed to the database:

    ```bash
    gator addfeed "Feed Name" "Feed URL"

**Follow a Feed**
Follow an existing feed by its URL:

    ```bash
    gator follow "Feed URL"

**Unfollow a Feed**
Unfollow a feed by its URL:

    ```bash
    gator unfollow "Feed URL"

**Browse Posts**
View the most recent posts from the feeds you follow:

    ```bash
    gator browse [limit]

The optional limit parameter specifies the number of posts to display (default is 2).

**Aggregate Feeds**
Fetch posts from all followed feeds periodically:

    ```bash
    gator agg [interval]
The interval parameter is a duration string like 1s, 1m, or 1h.

## Tools and Technologies
**Goose**: Used for managing database migrations.
**sqlc**: Used for generating type-safe SQL query code.
**Go**: The primary language for building Gator CLI.
**PostgreSQL**: The database for storing users, feeds, and posts.

## License
This project is licensed under the MIT License. See the LICENSE file for details.