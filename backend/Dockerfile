FROM golang:1.18

ADD . /backend

WORKDIR /backend

RUN go mod init backend

RUN go mod tidy
RUN go mod download
RUN go build -o backend
RUN chmod +x /backend

EXPOSE 8090

CMD ["./backend"]