FROM golang:1.18

ADD . /backend
WORKDIR /backend

RUN go mod init backend
RUN go mod tidy
RUN go mod download
RUN go build -o backend

EXPOSE 8080

CMD ["./backend"]