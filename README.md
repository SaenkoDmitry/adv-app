
## Advertisement app

* server by default running on :9100 port and postgres db running on :5432 port

* database mounting data in project directory /db/data

## Swagger docs

#### to see documentation you can open url :
http://localhost:9100/swagger/index.html

## Getting started

#### project use go modules so you can load all dependencies in your $GOPATH by command:

```
go mod download
```

#### to run server and db locally:
```
docker-compose up -d --build
```

#### to add migration to db:

```
migrate create -ext sql -dir db/migrations name_of_migration
```

#### to update swagger docs you need to add/update comment on method and execute :

```
swag init
```


## Technological Stack:

* Gin framework with inner http server (https://gin-gonic.com/)

* ORM library (https://github.com/jinzhu/gorm)

* migration library and CLI (https://github.com/golang-migrate/migrate)

* extension of gorm for pagination (https://github.com/biezhi/gorm-paginator)

* swagger docs (https://github.com/swaggo/swag)

* docker-compose (https://docs.docker.com/compose/)

