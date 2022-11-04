FROM golang:latest

WORKDIR /usr/app

COPY . . 

RUN go mod tidy
RUN go build

EXPOSE 8080

CMD ["./conceitoExato"]