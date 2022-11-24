FROM golang:1.19 AS build

WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/
RUN go mod download

COPY . /app/
RUN go build -o /main

EXPOSE 3333

CMD [ "/main" ]