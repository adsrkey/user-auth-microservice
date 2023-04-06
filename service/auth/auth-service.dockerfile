FROM golang:1.20.1 as builder

WORKDIR /go-microservice/

COPY . .

RUN CGO_ENABLED=0 go build -o microservice service/auth/cmd/app/main.go

FROM alpine:latest

WORKDIR /go-microservice

COPY --from=builder /go-microservice/ /go-microservice/

EXPOSE 8080

CMD ./microservice