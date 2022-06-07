# Go, Gorilla/mux, Gorm, Postgres Restful API

This is simple restful api crud with Golang

## Install Package

#### Gorilla/mux

```
  go get -u github.com/gorilla/mux
```

#### Gorm and Driver Postgres (i'm using postgres)

```
  go get -u gorm.io/gorm
  go get -u gorm.io/driver/postgres
```

## Run Server

#### type command below.

```
  go run main.go
```

## End Point

#### Users

| Method | URL                                          |
| ------ | -------------------------------------------- |
| GET    | http://localhost:3000/users/ (Get All)       |
| GET    | http://localhost:3000/users/1 (Get By Id)    |
| POST   | http://localhost:3000/users/ (Create)        |
| PUT    | http://localhost:3000/users/ (Update by id)  |
| DELETE | http://localhost:3000/users/1 (Delete By Id) |

#### Vehicle

| Method | URL                                                                    |
| ------ | ---------------------------------------------------------------------- |
| GET    | http://localhost:3000/vehicle/ (Get All)                               |
| GET    | http://localhost:3000/vehicle/1 (Get By Id)                            |
| POST   | http://localhost:3000/vehicle/ (Create)                                |
| PUT    | http://localhost:3000/vehicle/ (Update by id)                          |
| DELETE | http://localhost:3000/vehicle/1 (Delete by id)                         |
| GET    | http://localhost:3000/vehicle/all?sort=asc (Sort by id ex:desc or asc) |
| GET    | http://localhost:3000/vehicle/all?search=vespa (Search by name)        |
| GET    | http://localhost:3000/vehicle/popular (Get popular by Likes)           |

#### Historys

| Method | URL                                                                    |
| ------ | ---------------------------------------------------------------------- |
| GET    | http://localhost:3000/history/ (Get All)                               |
| GET    | http://localhost:3000/history/1 (Get By Id)                            |
| POST   | http://localhost:3000/history/ (Create)                                |
| PUT    | http://localhost:3000/history/2 (Update by id)                         |
| DELETE | http://localhost:3000/history/1 (Delete By Id)                         |
| GET    | http://localhost:3000/history/all?sort=asc (Sort by id ex:desc or asc) |
