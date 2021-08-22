FROM golang:alpine AS builder
WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd
COPY command command
COPY pkg pkg

ARG GOARCH
ARG GOARM
RUN GOARCH=${GOARCH} GOARM=${GOARM} CGO_ENABLED=0 go build ./cmd/stats

FROM gcr.io/distroless/static
COPY --from=builder /workspace/stats /bin
COPY config config
ENTRYPOINT ["stats"]
