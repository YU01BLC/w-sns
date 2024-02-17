FROM golang:1

RUN go install github.com/cosmtrek/air@latest

WORKDIR /go/backend/user_api

CMD ["air", "-c", ".air.toml"]
