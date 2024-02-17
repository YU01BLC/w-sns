FROM golang:1

COPY src/backend/user_api /go/user_api
WORKDIR /go/user_api

EXPOSE 8080

ENTRYPOINT [ "go", "run", "main.go" ]