FROM golang:1.16-alpine

RUN apk update

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /build

# install goose
RUN go get 'github.com/pressly/goose/cmd/goose'

# mount the app
COPY ./migrations migrations

# define goose as the entrypoint
ENTRYPOINT echo "Running migrations..." && goose -v -dir migrations postgres "host=${DB_HOST} user=${DB_USER} password=${DB_PASS} dbname=${DB_NAME} sslmode=disable" up
