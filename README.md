# Go, Gorilla/mux, Gorm, Postgres Restful API

This is simple restful api crud with Golang

## 🖥️ Tech Stack

**Backend:**

![golang](https://img.shields.io/badge/Go-100000?style=for-the-badge&logo=Go&logoColor=white&labelColor=51DEF0&color=51DEF0)&nbsp;
![postgresql](https://img.shields.io/badge/PostgreSQL-100000?style=for-the-badge&logo=PostgreSQL&logoColor=white&labelColor=3A7373&color=384A5F)&nbsp;
![jwt](https://img.shields.io/badge/JWT-100000?style=for-the-badge&logo=JSONWebTokens&logoColor=white&labelColor=000000&color=000000)&nbsp;

**Deployed On:**

![heroku](https://img.shields.io/badge/heroku-100000?style=for-the-badge&logo=Heroku&logoColor=white&labelColor=3C8932&color=3C8932)&nbsp;

## Install Package

#### Gorilla/mux

```
  go get -u github.com/gorilla/mux
```

#### Gorm and Driver Postgres

```
  go get -u gorm.io/gorm
  go get -u gorm.io/driver/postgres
```

## Run Server

#### type command below.

```
  go run main.go serve
```

## End Point

#### Users

| Service      | Method | URL      |
| ------------ | ------ | -------- |
| Get All      | GET    | /users/  |
| Get By Id    | GET    | /users/1 |
| Create       | POST   | /users/  |
| Update by id | PUT    | /users/  |
| Delete By Id | DELETE | /users/1 |

#### Vehicle

| Service                   | Method | URL                       |
| ------------------------- | ------ | ------------------------- |
| Get All                   | GET    | /vehicle/                 |
| Get By Id                 | GET    | /vehicle/1                |
| Create                    | POST   | /vehicle/                 |
| Update by id              | PUT    | /vehicle/                 |
| Delete by id              | DELETE | /vehicle/1                |
| Sort by id ex:desc or asc | GET    | /vehicle/all?sort=asc     |
| Search by name            | GET    | /vehicle/all?search=vespa |
| Get popular by Likes      | GET    | /vehicle/popular          |

#### Historys

| Service                   | Method | URL                   |
| ------------------------- | ------ | --------------------- |
| Get All                   | GET    | /history/             |
| Get By Id                 | GET    | /history/1            |
| Create                    | POST   | /history/             |
| Update by id              | PUT    | /history/2            |
| Delete By Id              | DELETE | /history/1            |
| Sort by id ex:desc or asc | GET    | /history/all?sort=asc |
