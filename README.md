## Build & Run

```
apt install docker.io -y && apt install docker-compose -y
git clone https://github.com/1makarov/go-crud-example
cd go-crud-example/
// add .env file
docker-compose up -d --build
```

##.env file:

```dotenv
POSTGRES_DB=library
POSTGRES_PASSWORD=postgres
POSTGRES_USER=root
POSTGRES_HOST=postgres

APP_PORT=80
JWT_SIGNING_KEY=cVRFw)29qgC|d4p
```

## Resources and Actions

    URL                           HTTP Method  Operation
    /api/v1/auth/create           GET          returns an auth key

    /api/v1/books/get/:id         GET          returns the book with id of :id
    /api/v1/books/get-all         GET          returns an array of books
    /api/v1/books/create          POST         create new book
    /api/v1/books/update/:id      POST         update the book by :id
    /api/v1/books/get/:id         GET          returns the book with id of :id
    /api/v1/books/delete/:id      DELETE       delete a book by :id



##book struct
```go
type Book struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	ISBN        string `json:"isbn"`
	Description string `json:"description"`
}
```