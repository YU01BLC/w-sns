FROM golang:1

RUN go install github.com/cosmtrek/air@latest

COPY src/backend /go/backend
WORKDIR /go/backend

RUN go mod download

EXPOSE 8080

WORKDIR /go/backend/user_api

CMD ["air", "-c", ".air.toml"]
