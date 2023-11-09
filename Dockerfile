FROM golang

WORKDIR /crypto-app

COPY . .

RUN go build -o crypto-app cmd/main.go

CMD ["./crypto-app"]