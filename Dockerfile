FROM golang:1.17-alpine

#RUN apk update \
#    && apk --no-cache upgrade \
#    && apk --no-cache add \
#    git \
#;

RUN mkdir -p /app/bin/mac
#COPY main.go /app/src

COPY ./ /app/src
WORKDIR /app/src

#RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/laravel-i18n -i main.go