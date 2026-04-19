FROM golang:1.21-alpine
WORKDIR /app
COPY server.go .
RUN go build -o server server.go
EXPOSE 1080
CMD ["./server"]
