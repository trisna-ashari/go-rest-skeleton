# Docs
Golang RESTful API boilerplate with modern architectures.

## [Go-Rest-Skeleton](https://github.com/trisna-ashari/go-rest-skeleton)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/bxcodec/faker/blob/master/LICENSE)
[![Go Report](https://goreportcard.com/badge/github.com/trisna-ashari/go-rest-skeleton)](https://goreportcard.com/report/github.com/trisna-ashari/go-rest-skeleton)

## Table of Contents
- [Getting Started](#getting-started)
- [Structures](#structures)
- [Features](#features)
    - [Better API Response](#better-api-response)
    - [Authentication](#authentication)
        - [JWT](#jwt)
        - [Basic Auth](#basic-auth)
        - [LDAP](#ldap)
    - [Role Based Access Permission](#cr-based-access-permission)
    - [DB Migration & Seeder](#db-migration-and-seeder)
        - [Auto Migrate](#auto-migrate)
        - [Builtin Seeders](#builtin-seeder)
    - [Internationalization](#internationalization)
    - [Logger](#logger)
    - [Test](#test)
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

git clone https://github.com/trisna-ashari/go-rest-skeleton
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
│   ├── tableEntity
│   ├── repository
│   ├── seeds
├── infrastructure
│   ├── authorization
│   ├── persistence
│   ├── security
│   ├── util
│   └── seed
├── interfaces
│   ├── handler
│   │   ├── v1.0       // handler version 1.0
│   │   │   ├── ...
│   │   │   ├── cr
│   │   │   └── userEntity
│   │   └── v2.0       // hanlder version 2.0
│   │   │   └── ...
│   └── middleware
├── languages
└── main.go
```

## Features
### Better API Response
All RESTful endpoints were designed with `prefix` and `versioning` support. Prefix format is: /`api`/`v1`/`external`/routes.

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

### Role Based Access Permission
There is builtin middleware called `policy`. It a middleware uses to handle access permission for each URI based on (method on the handler). This `policy` works by defined `cr` and `permission`.

- Each `userEntity` can have more than one cr.
- Each `cr` can have multiple permissions.
- Based on database to achieve dynamic and custom roles.

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
8:05AM INF Request ip=127.0.0.1 latency=1.110047 method=GET path=/api/v1/external/welcome payloads={} status=200 userEntity-agent=PostmanRuntime/7.26.1
8:05PM INF Request ip=127.0.0.1 latency=87.734547 method=POST path=/api/v1/external/login payloads={"email":"me@example.com","password":"123456"} status=200 userEntity-agent=curl/7.68.0
8:06AM INF Request ip=127.0.0.1 latency=2.405021 method=GET path=/api/v1/external/profile payloads={} status=200 userEntity-agent=PostmanRuntime/7.26.1
```

### Test
Coming Soon

## Documentation
### Entities
Entities represent each table schema on this skeleton. All entities stored in [domain/tableEntity](domain/tableEntity). Entities can contains definition of schema or structure, method, and validations. 


Here is current available entities:

- Module `modules`
- Permission `permissions`
- Role `roles`
- RolePermission `role_permissions`
- User `users`
- UserRole `user_roles`


### Middleware
##### API Version Middleware
File:
[interfaces/middleware/api_version.go](interfaces/middleware/api_version.go)

Description:
It's responsible to inject header `X-Api-Version` on response header.
Header `X-Api-Version` value based on api prefix: `api/v1`, `api/v2`, `api/v...`.

If an api route doesn't have prefix `v(1...n)`, this middleware wouldn't inject `X-Api-Version` on response header.

Example an api call with `v1` prefix:

```shell script
curl http://localhost:8888/api/v1/external/ping -i
```

will return:

```shell script
HTTP/1.1 200 OK
Accept-Language: en
Access-Control-Allow-Credentials: true
Access-Control-Allow-Headers: Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With
Access-Control-Allow-Methods: POST, OPTIONS, GET, PUT, PATCH, DELETE
Access-Control-Allow-Origin: *
Content-Type: application/json; charset: utf-8
X-Api-Version: v1
X-Request-Id: cff21ac8-770f-41e6-b575-7b631ffcef33
Date: Sat, 11 Jul 2020 23:01:41 GMT
Content-Length: 41

{"code":200,"data":null,"message":"pong"}
```

```shell script
curl http://localhost:8888/api/v2/external/ping -i
```

Example an api call with `v2` prefix:

will return:

```shell script
HTTP/1.1 200 OK
Accept-Language: en
Access-Control-Allow-Credentials: true
Access-Control-Allow-Headers: Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With
Access-Control-Allow-Methods: POST, OPTIONS, GET, PUT, PATCH, DELETE
Access-Control-Allow-Origin: *
Content-Type: application/json; charset: utf-8
X-Api-Version: v2
X-Request-Id: c7b66211-391e-4a2b-a9f6-626bdbbf5aac
Date: Sat, 11 Jul 2020 23:01:41 GMT
Content-Length: 41

{"code":200,"data":null,"message":"pong"}
```


##### Formatter Middleware
File:
[interfaces/middleware/formatter.go](interfaces/middleware/formatter.go)

Description:
Formatter middleware responsible to formulate all `success` response. Generally, there are three response keys (max four):
- `status` type: `integer`, represent `http code`
- `data` type: `interface{}` represent `response data`
- `message` type `string` represent `response message`
- `meta` type `interface` represent `response meta` such as `page`, `per_page`, `total` which only returned on get list of data.

Each `success` response through function named `Formatter`.

```go
Formatter(c *gin.Context, data interface{}, message string, meta interface{})
```

Here is an example usage:

```go
middleware.Formatter(c, nil, "PONG v1.0", nil)
```
will return: 
```json
{
    "code": 200,
    "data": null,
    "message": "PONG V1.0"
}
```

Parameter `message` type `string` will be translated if match with key in [translation](languages) files. If not match with any `key` in translation files, it will directly returned.

Current `active language` based: `Accept-Language` on request headers (if present).
If it not exists, will be applied default language that configured on `environment variale` named `APP_LANG`.

All translation process handled by [infrastructure/util/translation.go](infrastructure/util/translation.go). Check [this section](#translation) for more detail.

##### Logger Middleware
##### Policy Middleware
##### Request-Id Middleware
File: [interfaces/middleware/request_id.go](interfaces/middleware/request_id.go)

Description: 
It's responsible to `set` `Set-Request-Id` on request header. It uses for tracking the request that can be useful for debugging and logging. If the request send `Set-Request-Id` on request header, the value of `Set-Request-Id` will be forwarded to response header. If `Set-Request-Id` not exists on the request header, this middleware will inject `X-Request-Id` on response header automatically.

Example request with `Set-Request-Id` header:


##### Response Middleware

### Util
##### Translation
Go-rest-skeleton provide `translation util` to make `internalization` more easy. Translation file contain predefined key stored in`YAML` file.

Here is an example how to define key and use it in response:
- Define key in `YAML` file:
```yaml
api:
  msg:
    error:
      email_is_required: "Email is required"
      field_is_required: "Filed {{.Field}} is required"
    success:
      welcome_back: "Welcome back {{.Name}}, You have {{.TotalNotification}} notification"
```

- Call translation util function named `NewTranslation`:
  - Without passing data:
    ```go
    util.NewTranslation(c, "error", "api.msg.error.email_is_required", map[string]interface{}{})
    ```
    
    will return: 

    ```json
    {
        "code": 422,
        "data": null,
        "message": "Field email is required"
    }
    ```

   - With passing single data
    ```go
    errMsgData := make(map[string]interface{})
    errMsgData['Field'] = "email"
    util.NewTranslation(c, "error", "api.msg.error.field_is_required", errMsgData)
    ```
    will return: 
    ```json
    {
        "code": 422,
        "data": null,
        "message": "Field email is required"
    }
    ```
   - With passing multiple data
    ```go
    errMsgData := make(map[string]interface{})
    errMsgData['Name'] = "John"
    errMsgData['TotalNotification'] = 7
    util.NewTranslation(c, "error", "api.msg.success.welcome_back", errMsgData)
    ```
    will return:
    ```json
    {
        "code": 200,
        "data": null,
        "message": "Welcome back John, you have 7 notification"
    }
    ```

### Endpoints
##### Preparation API

| Method | URI Path                      | Description                                               |
|:------:|-------------------------------|-----------------------------------------------------------|
|  `GET` | /secret                       | Generate strong `APP_PRIVATE_KEY` and `APP_PUBLIC_KEY` that can be used on `.env` |

##### Authentication API

| Method | URI Path                      | Description                                               |
|:------:|-------------------------------|-----------------------------------------------------------|
| `POST` | /api/v1/external/auth/login   | Perform login with `email` and `password`                 |
| `POST` | /api/v1/external/auth/logout  | Perform logout for authenticated userEntity with `access_token` |
| `POST` | /api/v1/external/auth/refresh | Retrieve new `access_token` with `refresh_token`          |
| `GET`  | /api/v1/external/auth/profile | Retrieve `current authenticated userEntity` profile             |
| `PUT`  | /api/v1/external/auth/profile | Update `current authenticated userEntity` profile               |
| `GET`  | /api/v1/external/auth/switch-language | Change `language` preference by `language code`   |

##### User API
| Method | URI Path                      | Description                                               |
|:------:|-------------------------------|-----------------------------------------------------------|
| `POST` | /api/v1/external/users        | Retrieve userEntity list                                        |
| `GET`  | /api/v1/external/users        | Retrieve userEntity detail                                      |
| `POST` | /api/v1/external/users/:uuid  | Create a new userEntity                                         |
| `PUT`  | /api/v1/external/users/:uuid  | Update specified userEntity with `uuid`                         |
|`DELETE`| /api/v1/external/users/:uuid  | Delete specified userEntity with `uuid`                         |

##### Role API
| Method | URI Path                      | Description                                               |
|:------:|-------------------------------|-----------------------------------------------------------|
| `POST` | /api/v1/external/roles        | Retrieve cr list                                        |
| `GET`  | /api/v1/external/roles        | Retrieve cr detail                                      |
| `POST` | /api/v1/external/roles/:uuid  | Create a new cr                                         |
| `PUT`  | /api/v1/external/roles/:uuid  | Update specified cr with `uuid`                         |
|`DELETE`| /api/v1/external/roles/:uuid  | Delete specified cr with `uuid`                         |

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