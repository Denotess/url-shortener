# URL Shortener

A simple URL shortening service built with Go, Gin, and SQLite.

## What It Does

This is a basic API that takes long URLs and generates short, easy-to-share codes. When someone visits the short code, they get redirected to the original URL.

## How to Use

### Prerequisites

- Go 1.25 or higher
- SQLite3

### Installation

Clone the repository and install dependencies:

```bash
git clone https://github.com/Denotess/url-shortener.git
cd url-shortener
go mod download
```

### Running the Server

```bash
go run main.go
```

The server starts on `http://localhost:8080`.

### API Endpoints

**Health Check**

```
GET /ping
```

Returns a simple pong response.

**Shorten a URL**

```
POST /shorten
Content-Type: application/json

{
  "original": "https://www.example.com/very/long/url/that/is/annoying"
}
```

Response:

```json
{
  "short": "1a"
}
```

The short code is generated using base36 encoding of the database ID. Each URL gets a unique sequential ID, which is then converted to a compact base36 string.

**Redirect to Original URL**

```
GET /:short
```

Redirects (302) to the original URL. For example:

```
GET /1a
```

Will redirect to the URL that was shortened.

## How It Works

1. When you submit a URL to `/shorten`, it gets stored in the SQLite database.
2. The database assigns it an auto-incrementing ID.
3. That ID is converted to base36, creating a short code (e.g., ID 1 = "1", ID 36 = "10", ID 1296 = "100").
4. The short code is stored back in the database.
5. When someone visits `/:short`, the API looks up the original URL and redirects them.

## Database

The app uses SQLite with a simple schema:

```
Table: urls
- id: Auto-incrementing integer (primary key)
- short: Base36-encoded short code (unique)
- original: The full original URL
```

The database file is created automatically at `./data.db` on first run.

## Dependencies

- `github.com/gin-gonic/gin` - Web framework
- `github.com/mattn/go-sqlite3` - SQLite driver
- `github.com/blakewilliams/go-base36` - Base36 encoding
