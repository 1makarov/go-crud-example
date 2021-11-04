## Build & Run
```
apt update
apt install docker.io

git clone github.com/1makarov/go-crud-example

// add .env file

cd go-crud-example/

docker-compose up -d --build
```

.env file:
```dotenv
POSTGRES_DB=library
POSTGRES_PASSWORD=postgres
POSTGRES_USER=root
POSTGRES_HOST=postgres

APP_PORT=80
JWT_SIGNING_KEY=cVRFw)29qgC|d4p
```