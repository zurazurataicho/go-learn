FROM golang:bullseye

WORKDIR /usr/src/app

COPY go.mod main.go ./
RUN go mod tidy

RUN go build -v -o /usr/local/bin/app ./main.go

CMD ["app"]
