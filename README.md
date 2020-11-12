# Docs
Golang RESTful API boilerplate with modern architectures.

## [Go-Rest-Skeleton](https://github.com/trisna-ashari/go-rest-skeleton)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/bxcodec/faker/blob/master/LICENSE)
[![Go Report](https://goreportcard.com/badge/github.com/trisna-ashari/go-rest-skeleton)](https://goreportcard.com/report/github.com/trisna-ashari/go-rest-skeleton)

## Table of Contents
- [Getting Started](#getting-started)
- [Api Documentation](#api-documentation)
- [Structures](#structures)
- [Features](#features)
    - [Better API Response](#better-api-response)
    - [Authentication](#authentication)
        - [JWT](#jwt)
        - [Basic Auth](#basic-auth)
        - [Oauth](#oauth)
    - [Role Based Access Permission](#role-based-access-permission)
    - [DB Migration & Seeder](#db-migration-and-seeder)
        - [Auto Migrate](#auto-migrate)
        - [Builtin Seeders](#builtin-seeder)
    - [Internationalization](#internationalization)
    - [Logger](#logger)
    - [Test](#test)
- [Credits](#credits)
- [License](#license)
- [Copyright](#copyright)
- [Donation](#donation)

## Getting Started
#### Requirements

- Database: `MySQL` or `Postgres`
- Minio Server
- Redis
- Go v1.14.x

#### Install & Run
Download this project:

```shell script
git clone https://github.com/trisna-ashari/go-rest-skeleton
```

Download project dependencies:
```shell script
go mod download
```

Before run this project, you should set configs with yours.
Create & configure your `.env` based on: [.env.example](https://github.com/trisna-ashari/go-rest-skeleton/blob/master/.env.example)

Create app secret (private and public key):
```shell script
go run main.go create:secret
```

**NOTE**: you can use this generated key pair for `APP_PRIVATE_KEY` and `APP_PUBLIC_KEY` for your `.env`

Run migration:
```shell script
go run main.go db:migrate
```

Run initial seeder:
```shell script
go run main.go db:init
```

Fast run with:
```shell script
go run main.go

# running on default port 8888
```

or Enable hot reload with [Air](https://github.com/cosmtrek/air):

```shell script
go get -r github.com/cosmtrek/air
```

then simply run:

```shell script
air
```

**NOTE**: hot reload very useful on development processes

## Api documentation
This skeleton has builtin API documentation using [swagger](https://github.com/swaggo/swag). Just run this project and open this link:

http://localhost:8888/swagger.index.html

To rebuild api docs, simply run:
```shell script
swag init
```

## Structures

```md
├── application
├── config
├── docs                    // swagger
├── domain
│   ├── domain
│   ├── registry
│   ├── repository
│   ├── seeds
├── graph                   // an example of graphQL rpcServer
├── grpc                    // an example of gRPC rpcServer
├── infrastructure
│   ├── authorization
│   ├── message
│   ├── notify
│   ├── persistence
│   └── storage
├── interfaces
│   ├── cmd
│   ├── handler
│   │   ├── v1.0            // handler version 1.0
│   │   │   └── ...
│   │   └── v2.0            // hanlder version 2.0
│   │   │   └── ...
│   └── middleware
│   └── routers
│   └── service
├── languages
├── pkg                     // internal package
├── tests
└── main.go
```

## Features
### Better API Response
All RESTful endpoint has `prefix` and `versioning` support. Prefix format is: /`api`/`v1`/`external`/routes.

Supported HTTP Method: 
- `POST`
- `GET` 
- `PUT`
- `PATCH`
- `DELETE`
- `OPTIONS`

Api response generally consists by three `keys` (max four):
1) `code` as `HTTP Code`
2) `data` as `actual response (various type of data)`
3) `message` as `context message (success or error)`
4) `meta` as `additional response`. Check [example](#response-with-meta-pagination)

On api response's `headers`, its also included additional headers:
- `Accept-Language`
- `X-API-Version`
- `X-Request-Id`

#### Basic response
```shell script
curl http://localhost:8888/ping
```
will return:
```json
{
    "code": 200,
    "data": null,
    "message": "pong"
}
```

#### Response with meta pagination
Builtin meta pagination that easy to configure.
This is an example of response with meta pagination including `page`, `per_page`, `total`:
```json
{
    "code": 200,
    "data": [
        {
            "uuid": "5d082e28-7e8f-42a6-913e-8a939b77d1eb",
            "first_name": "Hortense",
            "last_name": "Lebsack",
            "email": "BAnSVKy@TrDbGME.info",
            "phone": "734-109-1286"
        },
        {
            "uuid": "b467e87e-0858-476e-b47a-174045dcdf71",
            "first_name": "Sylvan",
            "last_name": "Krajcik",
            "email": "DQDuvba@yCJshBR.net",
            "phone": "107-398-4261"
        }
    ],
    "message": "Successfully get userEntity list",
    "meta": {
        "page": 2,
        "per_page": 2,
        "total": 9
    }
}
```

### Authentication
#### JWT
This skeleton has builtin `JWT` based authentication. For an example:
```shell script
curl --location --request POST 'http://localhost:8888/api/v1/external/login' \
--header 'Accept-Language: ' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "me@example.com",
    "password": "123456"
}'
```
will return:
```json
{
    "code": 200,
    "data": {
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6ImUxY2U1NmI0LTM3MjQtNDA1OS1hOTY5LWUwMTY4YTBjYTllMiIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTU5NDIxODUwMCwidXVpZCI6ImNhZWIwZWZlLTFmZDEtNGQ1YS1iOWNkLTIyMTNiOTc2OGU5MyJ9.2uSkiwISaL4R5wwpAgx7isRpnpSE6GtE8gumOY_yR1E",
        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTQ4MjI0MDAsInJlZnJlc2hfdXVpZCI6ImUxY2U1NmI0LTM3MjQtNDA1OS1hOTY5LWUwMTY4YTBjYTllMisrY2FlYjBlZmUtMWZkMS00ZDVhLWI5Y2QtMjIxM2I5NzY4ZTkzIiwidXVpZCI6ImNhZWIwZWZlLTFmZDEtNGQ1YS1iOWNkLTIyMTNiOTc2OGU5MyJ9.oGrudpMa57pNz9qGOx1DpLjg1He-xWmqT2gGS84hEww",
    },
    "message": "Successfully login"
}
```

#### Basic Auth
There also builtin authentication using `Basic Auth` by passing auth through header.

Basic auth constructed from based64 encoded of:
```
email:password
me@example.com:123456
```

Example:

```shell script
curl --location --request GET 'http://localhost:8888/api/v1/external/profile' \
--header 'Accept-Language: en' \
--header 'Authorization: Basic dHJpc25hLngyQGdtYWlsLmNvbToxMjM0NTY='
```

#### Oauth
There are builtin oauth2 server and client. For an example click this link:

http://localhost:8181/oauth/login

### Role Based Access Permission
There is builtin middleware called `policy`. It a middleware uses to handle access permission for each URI based on (method on the handler). This `policy` works by defined `custom role` and `permission`.

- Each `userEntity` can have more than one `custom role`.
- Each `custom role` can have multiple permissions.
- Based on the database to achieve dynamic and custom roles.

See this [example](#policy-middleware) how easy to implement.

### DB Migration and Seeder
Yes, this skeleton has builtin db migration and seeders.
#### Auto Migrate
Why auto migrate? This feature is very helpful to keep your `table(s) schema` always update depends on changes in each `entities`. 

`AutoMigrate` automatically run when you `manually` start the application or with triggered by `hot reload`.

#### Builtin Seeder
This builtin seeder can help you to fill your schema with dummy data, so you don't need wast your time to type `an lorem ipsum`.


### Internationalization
Internationalization made easy with this skeleton. [go-i18n](https://github.com/nicksnyder/go-i18n) was used to handle multilingual support. All translations text stored in *.`yaml` file on `languages` directory.
```
. . .
├── languages
│   ├── global.en.yaml
│   └── global.id.yaml
. . .
```

`YAML` file was choosen because `nested declaration` can be done easily instead of `TOML`, `JSON`, etc. For more example please check this [language example](https://github.com/trisna-ashari/go-rest-skeleton/tree/master/languages).

### Logger
Yes, `logger` is very useful in development process. This skeleton has built in logger to watch any request are coming to your rest API. It was configurable :).

This is examples of what logger prints:
```shell script
2020/11/04 16:01:56 /Users/sentinel/Gitrepos/go-rest-skeleton/infrastructure/persistence/user_repository.go:176
[7.798ms] [rows:1] SELECT * FROM `users` WHERE email = "me@example.com" AND `users`.`deleted_at` IS NULL LIMIT 1
4:01PM INF Request headers={"Accept":"*/*","Accept-Encoding":"gzip, deflate, br","Accept-Language":"id","Cache-Control":"no-cache","Connection":"keep-alive","Content-Length":"66","Content-Type":"application/json","Postman-Token":"fff3d097-e315-4103-8ec2-7fb945fb654f","User-Agent":"PostmanRuntime/7.26.5"} ip=127.0.0.1 latency=160.543176 method=POST path=/api/v1/external/login request-form={} request-id=ba38e9a6-4f7e-4e67-b499-7620ef3bef0a request-payload={"email":"me@example.com","password":"123456"} response={"code":200,"data":{"access_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjNiYzdiZTk4LTgzOWMtNDExZi04ZGZhLWMyNjU4OGQ5Njg5YiIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTYwNDQ4NDExNiwidXVpZCI6IjNhY2E2NWE5LWQ3MWUtNGJhNy05YzQxLThmMGE0ZjFiZGVhYiJ9.EcniXexsASuPapM0SpKDNrkUE-RR0TmgwfsNXn7DC5A","refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDUwODUzMTYsInJlZnJlc2hfdXVpZCI6IjNiYzdiZTk4LTgzOWMtNDExZi04ZGZhLWMyNjU4OGQ5Njg5YisrM2FjYTY1YTktZDcxZS00YmE3LTljNDEtOGYwYTRmMWJkZWFiIiwidXVpZCI6IjNhY2E2NWE5LWQ3MWUtNGJhNy05YzQxLThmMGE0ZjFiZGVhYiJ9.LO2W2DUpFuEM9gWbPR0s1Enud-dNZq1IkuXARA1tBuw"},"message":"Berhasil masuk"} status=200 user-agent=PostmanRuntime/7.26.5
[GIN] 2020/11/04 - 16:01:56 | 200 |   165.35311ms |       127.0.0.1 | POST     "/api/v1/external/login"

```

### Test
Absolutely yes, just run:
```shell script
go test -p 1 ./... -cover -coverprofile=coverage.out 
```

or using the [Makefile](https://github.com/trisna-ashari/go-rest-skeleton/blob/master/Makefile):
```
make unit test
make integration test
```

## Credits
- [Go](https://github.com/golang/go) - The Go Programming Language
- [gin](https://github.com/gin-gonic/gin) - Gin is HTTP web framework written in Go (Golang)
- [gorm](https://github.com/go-gorm/gorm) - The fantastic ORM library for Golang
- [swag](https://github.com/swaggo/swag) - Automatically generate RESTful API documentation with Swagger 2.0 for Go
- [oauth2](https://github.com/go-oauth2/oauth2) - OAuth 2.0 server library for the Go programming language
- [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) - An idiomatic Go (golang) validation package

## License

MIT License. See [LICENSE](https://github.com/trisna-ashari/go-rest-skeleton/blob/master/LICENSE) for details.

## Copyright

Copyright (c) 2020 Trisna Novi Ashari.

## Donation

[Paypal](https://paypal.me/trisnaashari)

Buy me a cup of coffee :coffee: :)