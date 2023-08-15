FROM golang

WORKDIR /crypto-tgbot

COPY . .

RUN go build -o crypto-tgbot cmd/main.go

CMD ["./crypto-tgbot"]