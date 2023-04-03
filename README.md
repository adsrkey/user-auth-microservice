# Authentication microservice

## Structure

    ├── deploy
    ├── pkg               
    │   ├── store        
    │   │   └── postgres
    │   ├── tool
    │   │   └── transaction
    │   └── type
    │
    ├── service
    │   └── auth
    │       ├── cmd
    │       │   └── app
    │       ├── configs
    │       ├── internal
    │       │   ├── delivery
    │       │   │   └── http
    │       │   │       ├── cookie
    │       │   │       └── validator
    │       │   │           └── auth
    │       │   ├── domain
    │       │   │   └── user
    │       │   ├── repository
    │       │   │   └── storage
    │       │   │       └── postgres
    │       │   │           ├── dao
    │       │   │           └── migrations
    │       │   └── usecase
    │       │       ├── adapters
    │       │       │    └── storage
    │       │       └── user
    │       ├── proto
    │       └── utils
    │            └── jwt
    └── ...