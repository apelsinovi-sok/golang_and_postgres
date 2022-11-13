FROM golang:latest

RUN mkdir /server

WORKDIR /server

COPY  ./ ./

RUN go env -w GO111MODULE=on

RUN go mod download

RUN #go build -o main .

RUN go build main.go

CMD ["./main"]