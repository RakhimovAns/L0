FROM golang:1.24.5

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./bin/app ./cmd/app/main.go

RUN go build -o ./bin/migrator ./cmd/migrator/main.go

EXPOSE 9000

CMD ["./bin/app"]