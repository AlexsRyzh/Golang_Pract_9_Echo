FROM golang:1.21

WORKDIR /usr/src/pract9

COPY go.mod go.sum ./
COPY main.go ./
COPY internal ./internal
COPY .env ./

CMD ["go","run","main.go"]

EXPOSE 8080
