# build stage
FROM golang:alpine as builder

RUN apk add -U make

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY Makefile .

RUN make download

COPY . .

RUN make linux BINARY=cli

# final stage
FROM scratch

COPY --from=builder /app/cli /app/

EXPOSE 3000

ENTRYPOINT ["/app/cli"]
