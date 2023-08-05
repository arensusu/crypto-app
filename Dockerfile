FROM golang

WORKDIR /crypto-tgbot

COPY . .

RUN go build

CMD ["./funding-rate"]