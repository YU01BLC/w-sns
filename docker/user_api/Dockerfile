FROM golang:1

COPY src/backend /go/backend
WORKDIR /go/backend

RUN go mod tidy

EXPOSE 8080

WORKDIR /go/backend/user_api

ENTRYPOINT [ "go", "run", "main.go" ]