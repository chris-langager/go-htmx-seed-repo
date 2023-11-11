# build
FROM golang:1.21.4-alpine AS build

WORKDIR /
ADD . .

RUN go build -o server ./cmd/server/main.go 


# run
FROM alpine:3.18.4

WORKDIR /
COPY --from=build /server /server

ENTRYPOINT ./server