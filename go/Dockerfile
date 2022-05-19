FROM golang:1.18

ENV GOOS=linux
ENV GOARCH=amd64

RUN mkdir /src
WORKDIR /src

COPY go.mod /src
COPY go.sum /src
RUN go mod download

COPY * /src
RUN go build -o /app .

CMD ["/app"]
