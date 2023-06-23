FROM golang:1.20-alpine

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o /server
RUN go run migrate/migrate.go

EXPOSE ${PORT}

CMD [ "go" , "run", "main.go" ]