FROM golang:1.20.6

RUN apt update

RUN apt install -y libmapnik-dev

WORKDIR /usr/src/app
COPY . .
RUN go get github.com/gin-gonic/gin github.com/omniscale/go-mapnik
