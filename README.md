## Build & Run
```
apt install docker.io -y && apt install docker-compose -y
git clone https://github.com/1makarov/go-crud-example
cd go-crud-example/
// add .env file
docker-compose up -d --build
```
## Swagger
```
http://localhost:8080/swagger/index.html
```
## .env

```dotenv
POSTGRES_DB=library
POSTGRES_PASSWORD=postgres
POSTGRES_USER=root
POSTGRES_HOST=postgres

JWT_SIGNING_KEY=cVRFw)29qgC|d4p
```
