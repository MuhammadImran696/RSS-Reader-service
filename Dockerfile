FROM golang:1.12.0-alpine3.9
# ENV GO111MODULE=on
# RUN go mod download
RUN mkdir /app
ADD . /app
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /app

RUN apk update
RUN apk add git
## we run go build to compile the binary
## executable of our Go program
RUN go mod download


RUN go build -o main 
## Our start command which kicks off
## our newly created binary executable
EXPOSE 9000
CMD ["/app/main"]