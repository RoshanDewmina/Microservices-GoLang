FROM golang:1.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main .
FROM gcr.io/distroless/base-debian11
COPY --from=builder /app/main .
CMD ["/main"]
