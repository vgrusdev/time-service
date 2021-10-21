FROM golang:1.17 AS builder

WORKDIR /go/src
COPY src .
#RUN go build -o time-service time-service.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o time-service time-service.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /go/src/time-service .

EXPOSE 8080

CMD ["./time-service"]

