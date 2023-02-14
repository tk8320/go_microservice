FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app

ENV GO111MODULE off

RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/mux

ENV GOPATH /app


RUN go build -o main .

EXPOSE 8080
RUN go test -v . 
CMD ["/app/main"]

