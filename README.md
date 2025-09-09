# url-shortener

A simple and efficient URL shortening service that converts long, messy links into short, shareable ones.

## Features

- Shorten any URL to a unique short code
- Redirect short URLs to their original destinations
- RESTful API built with [Gin](https://github.com/gin-gonic/gin)
- Secure random code generation
- In-memory storage for URLs
- Easy to run locally

## Getting Started

### Prerequisites

- Go 1.18+
- (Optional) Docker

### Installation

Clone the repository:

```sh
git clone https://github.com/spike510/url-shortener.git
cd url-shortener
```

### Running Locally

```sh
go run ./cmd/server/main.go
```

By default, the server runs on `http://localhost:8080`.

You can set environment variables:

- `PORT` – server port (default: 8080)
- `BASE_URL` – base URL for shortened links (default: http://localhost:PORT)

### API Endpoints

#### Shorten a URL

**POST** `/api/shorten`

Request body:
```json
{
  "url": "https://example.com/very/long/url"
}
```

Response:
```json
{
  "code": "abc123",
  "short_url": "http://localhost:8080/abc123"
}
```

#### Redirect

**GET** `/:code`

Redirects to the original URL.

### Example Usage

```sh
curl -X POST http://localhost:8080/api/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com"}'
```

### Testing

Run unit tests with:

```sh
go test ./...
```

### Project Structure

```
cmd/server/         # Main server entrypoint
internal/http/      # HTTP handlers and routing
internal/generator/ # Short code generator
internal/storage/   # In-memory URL storage
```

### License

MIT

---
