FROM golang:1.18

ENV GOOS=linux
ENV GOARCH=amd64

RUN mkdir /app
WORKDIR /app

COPY go.mod /app
COPY go.sum /app
RUN go mod download

COPY * /app
RUN go build

CMD ["./app"]
