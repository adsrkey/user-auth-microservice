# Authentication microservice

## How to build

### Docker-compose with make:

```
    cd .\deploy\
    make up_build
```

---

## Endpoints

### cURL examples:

#### /api/v1/reg - registration

```
curl --location 'http://localhost:8080/api/v1/reg' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "test@test.com",
    "password": "password"
}'
```

#### Response status code example:

- 201 | Created,
- 409 | Conflict,
- 400 | Bad Request,
- 500 | Internal Server Error,
- 503 | Service Unavailable

###### request / response format: application/json

###### ok:

``
{
"message": "user registered"
}
``

###### error:

``
{
"code": 409,
"developer_message": "user with such data is already registered"
}
``

---

#### /api/v1/auth - authentication

```
curl --location --request POST 'http://localhost:8080/api/v1/auth' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjOWE5ZjY3MS1lMmI1LTRiMzQtYWQzOS0yYjBkZDY2MjZhY2IiLCJleHAiOjE2ODE3NTU5OTcsImlhdCI6MTY4MTc1MjM5NywiaXNzIjoiYzlhOWY2NzEtZTJiNS00YjM0LWFkMzktMmIwZGQ2NjI2YWNiIiwic3ViIjoibG9nZ2VkX2luIiwiU2Vzc2lvbklEIjoiNzlmZmIwMzAtMGViMC00ZTM1LWFmM2ItOWI2NjBhNThlYTg3IiwiVXNlcklEIjoiYzlhOWY2NzEtZTJiNS00YjM0LWFkMzktMmIwZGQ2NjI2YWNiIn0.p96FzRNCZH9m8Vezb3aLdmyfK69aWHxmcv7DWexu22E' \
--data ''
```

#### Response status code example:

- 200 | Status OK,
- 401 | Unauthorized,
- 400 | Bad Request,
- 500 | Internal Server Error,
- 503 | Service Unavailable

###### request / response format: application/json

###### ok:

``
{
"message": "token is valid"
}
``

###### error:

``
{
"code": 400,
"developer_message": "user not authorized"
}
``

#### /api/v1/login - login

```
curl --location 'http://localhost:8080/api/v1/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "test@test.com",
    "password": "password"
}'
```

#### Response status code example:

- 200 | Status OK,
- 401 | Unauthorized,
- 400 | Bad Request,
- 500 | Internal Server Error,
- 503 | Service Unavailable

###### request / response format: application/json

###### ok:

``
{
"message": "user authorized",
"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJmMTVhYTc1OC03NjJjLTQzYWYtYThlZC02ZTU3ZWQ4OWM5ZWMiLCJleHAiOjE2ODE3NTY4ODMsImlhdCI6MTY4MTc1MzI4MywiaXNzIjoiZjE1YWE3NTgtNzYyYy00M2FmLWE4ZWQtNmU1N2VkODljOWVjIiwic3ViIjoibG9nZ2VkX2luIiwiU2Vzc2lvbklEIjoiOTMyNGYzNmItZGJlYi00MTYwLThhNjUtMWYyZjNlMDEyMzg5IiwiVXNlcklEIjoiZjE1YWE3NTgtNzYyYy00M2FmLWE4ZWQtNmU1N2VkODljOWVjIn0.G6InRQU1--pebGjcLLxZavJvykMlvse7E_jz1oWHVu4"
}
``

###### error:

``
{
"code": 400,
"developer_message": "user not authorized"
}
``

## Folder structure

    .
    ├── deploy
    ├── pkg               
    │   ├── store        
    │   │   └── postgres
    │   ├── tool
    │   │   └── transaction
    │   └── type
    │
    └── service
        └── auth
            ├── cmd
            │   └── app
            ├── configs
            ├── internal
            │   ├── delivery
            │   │   └── http
            │   │       ├── cookie
            │   │       ├── middleware
            │   │       ├── response
            │   │       └── validator
            │   ├── domain
            │   │   └── user
            │   ├── repository
            │   │   └── storage
            │   │       └── postgres
            │   │           ├── dao
            │   │           ├── migrations
            │   │           └── worker
            │   └── usecase
            │       ├── adapters
            │       │    └── storage
            │       └── user
            ├── proto
            └── utils
                └── jwt
