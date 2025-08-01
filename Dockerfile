# Build
FROM golang:alpine as builder
WORKDIR /app
COPY . .

RUN go env -w GO111MODULE=on \
 && go env -w CGO_ENABLED=0 \
 && go mod tidy \
 && go build -o server .

# Runtime
FROM alpine:latest
LABEL maintainer="zzdpk2"

WORKDIR /app
COPY --from=builder /app/server ./server
COPY --from=builder /app/config.yaml ./config.yaml
COPY --from=builder /app/.kube/config ./.kube/config

EXPOSE 8082
ENTRYPOINT ["./server"]
