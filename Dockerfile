RUN go build -o main .

FROM golang:latest

WORKDIR /app/campaignweb

COPY --from=0 /app/main .


EXPOSE 8080


ENV DB_HOST=localhost
ENV DB_PORT=8080
ENV DB_USER=myuser
ENV DB_PASSWORD=mypassword
ENV DB_NAME=mydatabase


CMD ["./main"]
