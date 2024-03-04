FROM golang:latest

ENV CONFIG_FILE=config.yaml

WORKDIR /app

COPY server /app/server
COPY configs /app/configs
COPY go.mod /app
COPY go.sum /app
COPY merkle_tree /app/merkle_tree
COPY proto /app/proto
COPY utils /app/utils

RUN go build /app/server/src/main.go && chmod +x /app/main

CMD ["/app/main"]