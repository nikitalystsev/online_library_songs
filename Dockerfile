FROM golang:latest

WORKDIR /usr/local/src

COPY ["./go.mod", "./go.sum", "./"]

RUN go mod download

COPY ./ ./

RUN go build -o ./app cmd/app/main.go

CMD ["./app"]