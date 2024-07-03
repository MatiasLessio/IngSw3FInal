FROM golang:1.20

EXPOSE 8090

ADD . /backend
WORKDIR /backend

RUN go mod init backend
RUN go mod tidy
RUN go build -o backend
RUN chmod +x /backend

CMD ["./backend"]