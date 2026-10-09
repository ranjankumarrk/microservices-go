# microservices-go

A small Go REST API for a book-swap service: list books from MySQL and a login endpoint, routed with `gorilla/mux`.

## Endpoints

| Method | Path | Description |
| --- | --- | --- |
| GET | `/api/books` | All books (`id`, `title`, `author`, `owner`, `publishYear`, `price`) |
| POST | `/api/v1/login` | Login (stub) |

## Run

Needs MySQL on `127.0.0.1:3306` with a `bookswap` database and a `books` table. The DSN in `main.go` is a local placeholder.

```bash
go run .
# server on :8080
curl localhost:8080/api/books
```

## Author

**Ranjan Kumar** (`ranjankumarrk`) — [Golang developer](https://ranjankumarrk.github.io/golang-developer.html) and [Node.js / NestJS developer](https://ranjankumarrk.github.io/nodejs-developer.html), Senior Software Engineer, Hyderabad.

[Portfolio](https://ranjankumarrk.github.io/) · [Ranjan Kumar on LinkedIn](https://www.linkedin.com/in/ranjankumarrk/) · [GitHub](https://github.com/ranjankumarrk)
