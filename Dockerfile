# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
COPY sa.json / 

RUN go build -o /stonehenge

ENV HTTP_PORT=8080

EXPOSE 8080

CMD [ "/stonehenge" ]