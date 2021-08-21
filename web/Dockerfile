FROM golang:alpine AS builder
WORKDIR /go/src/stats
COPY . .
RUN go build -o stats ./cmd/stats

FROM alpine
COPY --from=builder /go/src/stats/stats bin/stats
COPY config config
ENTRYPOINT ["stats"]
