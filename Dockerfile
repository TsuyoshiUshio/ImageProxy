FROM golang:1.7.5

ADD server.go .
RUN go get -u github.com/labstack/echo
RUN go build server.go

ENTRYPOINT ./server

EXPOSE 8090