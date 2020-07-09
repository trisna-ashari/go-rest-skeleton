# Docs
Golang RESTful API boilerplate with modern architectures.

## [Go-Rest-Skeleton](https://github.com/trisna-ashari/go-rest-skeleton)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/bxcodec/faker/blob/master/LICENSE)

## Table of Contents
- [Getting Started](#getting-started)
- [Structures](#structures)
- [Features](#features)
    - [Better API Response](#better-api-response)
    - [Authentication](#authentication)
        - [JWT](#jwt)
        - [Basic Auth](#basic-auth)
        - [LDAP](#ldap)
    - [DB Migration & Seeder](#db-migration-and-seeder)
        - [Auto Migrate](#auto-migrate)
        - [Builtin Seeders](#builtin-seeder)
    - [Internationalization](#internationalization)
    - [Logger](#logger)
- [Documentation](#documentation)
- [Credits](#credits)
- [License](#license)
- [Copyright](#copyright)
- [Donation](#donation)

## Getting Started
#### Requirements

- Database: `MySQL` or `Postgres`
- Redis
- Go v1.14.x

#### Install & Run

```shell script
# Download this project

go get https://github.com/trisna-ashari/go-rest-skeleton
```

Before running API server, you should set configs with yours.
Create & configure your `.env` based on: [.env.example](https://github.com/trisna-ashari/go-rest-skeleton/blob/master/.env.example)

Fast run with:

```shell script
go run main.go

# API Endpoint : http://127.0.0.1:8888
```

or Enable hot reload with [Air](https://github.com/cosmtrek/air):

```shell script
go get -u github.com/cosmtrek/air
```

then simply run:

```shell script
air
```

**NOTE**: hot reload very useful on development processes

## Structures

```shell script
├── application
├── domain
│   ├── entity
│   ├── repository
│   ├── seeds
├── infrastructure
│   ├── authorization
│   ├── persistence
│   ├── security
│   └── seed
├── interfaces
│   ├── handler
│   │   ├── v1.0
│   │   │   ├── ...
│   │   │   ├── role
│   │   │   └── user
│   │   └── v2.0
│   │   │   └── ...
│   └── middleware
├── languages
└── main.go
```

## Features
### Better API Response
All endpoints were designed with `prefix` and `versioning` support. For example: /`api`/`v1`/`external`/routes.

Supported HTTP Method: `POST`, `OPTIONS`, `GET`, `PUT`, `PATCH`, `DELETE`.

Api response generally includes three `keys`:
1) `code` as `HTTP Code`
2) `data` as `actual response (various type of data)`
3) `message` as `context message (success or error)`
4) `meta` as `additional response`. Check [example](#response-with-meta-pagination)

On api response's `headers`, its also included additional headers:
- `X-API-Version`
- `X-Request-Id`
- `Accept-Language`

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
    "message": "Successfully get user list",
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
curl --location --request POST 'http://localhost:8888/api/v1/external/login' --header 'Accept-Language: ' --header 'Content-Type: application/json' --data-raw '{
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
        "first_name": "Trisna",
        "language": "en",
        "last_name": "Ashari",
        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTQ4MjI0MDAsInJlZnJlc2hfdXVpZCI6ImUxY2U1NmI0LTM3MjQtNDA1OS1hOTY5LWUwMTY4YTBjYTllMisrY2FlYjBlZmUtMWZkMS00ZDVhLWI5Y2QtMjIxM2I5NzY4ZTkzIiwidXVpZCI6ImNhZWIwZWZlLTFmZDEtNGQ1YS1iOWNkLTIyMTNiOTc2OGU5MyJ9.oGrudpMa57pNz9qGOx1DpLjg1He-xWmqT2gGS84hEww",
        "uuid": "caeb0efe-1fd1-4d5a-b9cd-2213b9768e93"
    },
    "message": "Successfully login"
}
```

#### Basic Auth
Coming Soon
#### LDAP
Coming Soon

### DB Migration and Seeder
Yes, this skeleton has builtin db migration and seeders.
#### Auto Migrate
Why auto migrate? This feature is very helpful to keep your `table(s) schema` always update depends on changes in each `entities`. 

AutoMigrate automatically run when you start the application `manually` or with `hot reload`.

#### Builtin Seeder
This builtin seeder can help you to fill your schema with dummy data, so you don't need wast your time to type `an lorem ipsum`.


### Internationalization
Internationalization made easy with this skeleton. [go-i18n](https://github.com/nicksnyder/go-i18n) was used to handle multilingual support. All translations text stored in *.`yaml` file on `languages` directory.
```shell script
. . .
├── languages
│   ├── global.en.yaml
│   └── global.id.yaml
. . .
```
For example of *.`yaml` file:
```yaml
api:
  msg:
    error:
      an_error_occured: "An Error Occurred"
      internal_server_error: "Internal Server Error"
      not_found: "Not Found"
      bad_request: "Bad Request"
      unauthorized: "Unauthorized"
      forbidden: "Forbidden"
      unprocessable_entity: "Unprocessable Entity"
    success:
      successfully_login: "Successfully Login"
      successfully_logout: "Successfully Logout"
      successfully_switch_language: "Successfully Switch Language"
      successfully_refresh_token: "Successfully Get New Token"
      successfully_get_profile: "Successfully Get Profile"
      successfully_get_user_list: "Successfully Get User List"
      successfully_get_user_detail: "Successfully Get User Detail"
      successfully_create_user: "Successfully Create a New User"
```
`YAML` file was choosen because `nested declaration` can be done easily instead of `TOML`, `JSON`, etc. For more example please check this [language example](https://github.com/trisna-ashari/go-rest-skeleton/tree/master/languages).

### Logger
Yes, `logger` is very useful in development process. This skeleton has built in logger to watch any request are coming to your rest API. It was configurable :).

This is examples of what logger prints:
```shell script
8:05AM INF Request ip=127.0.0.1 latency=1.110047 method=GET path=/api/v1/external/welcome payloads={} status=200 user-agent=PostmanRuntime/7.26.1
8:05PM INF Request ip=127.0.0.1 latency=87.734547 method=POST path=/api/v1/external/login payloads={"email":"me@example.com","password":"123456"} status=200 user-agent=curl/7.68.0
8:06AM INF Request ip=127.0.0.1 latency=2.405021 method=GET path=/api/v1/external/profile payloads={} status=200 user-agent=PostmanRuntime/7.26.1
```

## Documentation
Coming Soon

## Credits
- [Go](https://github.com/golang/go) - The Go Programming Language
- [gin](https://github.com/gin-gonic/gin) - Gin is HTTP web framework written in Go (Golang)
- [gorm](https://github.com/go-gorm/gorm) - The fantastic ORM library for Golang
 
## License

MIT License. See [LICENSE](https://github.com/trisna-ashari/go-rest-skeleton/blob/master/LICENSE) for details.

## Copyright

Copyright (c) 2017-2019 Trisna Novi Ashari.

## Donation

[Paypal](https://paypal.me/trisnaashari)

Buy me a cup of coffee :coffee: :)