# Authentication microservice

# Endpoints

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

#### /api/v1/reg - registration
```
curl --location 'http://localhost:8080/api/v1/auth' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "test@test.com",
    "password": "test"
}'
```

## Structure
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
