
FROM golang:latest

WORKDIR /app

COPY . .

COPY go.mod go.sum ./
RUN go mod download

RUN go build -o main .

FROM golang:latest

WORKDIR /app/campaignweb

COPY --from=0 /app/main .


EXPOSE 8080


ENV DB_HOST=localhost
ENV DB_PORT=5432
ENV DB_USER=myuser
ENV DB_PASSWORD=mypassword
ENV DB_NAME=mydatabase


CMD ["./main"]
