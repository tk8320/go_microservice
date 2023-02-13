FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
## Add this go mod download command to pull in any dependencies
# RUN go mod download

ENV GO111MODULE off

RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/mux

ENV GOPATH /app

## Our project will now successfully build with the necessary go libraries included.

RUN go build -o main .
## Our start command which kicks off
## our newly created binary executable
EXPOSE 8080
CMD ["/app/main"]

# RUN go test -v . > test.out