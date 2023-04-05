# Authentication microservice

## Endpoints

### cURL examples:

#### /api/v1/reg - registration
```
curl --location 'http://localhost:8080/api/v1/reg' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "test@test.com",
    "password": "test"
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
curl --location 'http://localhost:8080/api/v1/auth' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "test@test.com",
    "password": "test"
}'
```

#### Response status code example:

- 202 | Status Accepted,
- 401 | Unauthorized,
- 400 | Bad Request,
- 500 | Internal Server Error,
- 503 | Service Unavailable

###### request / response format: application/json

###### ok:
``
{
"message": "user authorized"
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
