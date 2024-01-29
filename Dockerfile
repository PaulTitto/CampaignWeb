FROM golang:latest

WORKDIR /app/campaignweb

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
