# build
FROM golang:1.21.4-alpine AS build

WORKDIR /
ADD . .

RUN go build -o server ./cmd/server/main.go 


# run
FROM alpine:latest AS run

WORKDIR /
COPY --from=build /server /server

ENTRYPOINT ./server