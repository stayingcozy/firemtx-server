FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build .

CMD ["/app/firemtx-server"]

