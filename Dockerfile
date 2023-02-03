# syntax=docker/dockerfile:1

# Build
FROM golang:1.19-alpine as build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build -o /german-angst-api


# Deploy
FROM alpine:latest

WORKDIR /

COPY --from=build /german-angst-api /german-angst-api

EXPOSE 8080

ENTRYPOINT [ "/german-angst-api" ]