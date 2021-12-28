
#Build stage
FROM golang:1.17.5-alpine3.15 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

#Run stage
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/main .


EXPOSE 5050
CMD [ "/app/main" ] 