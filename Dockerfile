FROM golang:1.15.6-alpine3.12

LABEL maintaner="Hadyan <hadyanadamn@gmail.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . /app

ENV PORT 8000

RUN go build

RUN find . -name "*.go" -type f -delete

EXPOSE ${PORT}

CMD ./go-gorm-gin