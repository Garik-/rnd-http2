FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
