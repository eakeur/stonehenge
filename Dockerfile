# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY sa.json ./

COPY . ./

RUN go build -o /stonehenge

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /stonehenge /stonehenge

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/stonehenge"]