FROM golang:1.20.6

RUN apt update

RUN apt install -y libmapnik-dev git

WORKDIR /usr/src/app
COPY . .
RUN go get github.com/gin-gonic/gin
RUN git clone https://github.com/kki-org/go-mapnik $(go env GOROOT)/src/go-mapnik
