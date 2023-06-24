FROM golang:1.20-alpine

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o /server
RUN chmod +x /server

EXPOSE ${PORT}

CMD ["/server" ]