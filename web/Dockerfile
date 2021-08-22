FROM golang:alpine AS builder
WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd
COPY command command
COPY pkg pkg
RUN CGO_ENABLED=0 go build ./cmd/stats

FROM gcr.io/distroless/base
COPY --from=builder /workspace/stats /bin
COPY config config
ENTRYPOINT ["stats"]
