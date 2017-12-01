FROM golang:1.8

RUN mkdir -p /go/src/goyo.in/gpstracker
WORKDIR /go/src/goyo.in/gpstracker
COPY . /go/src/goyo.in/gpstracker

RUN go get github.com/codegangsta/gin

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

ENV PORT 6969 6979

EXPOSE 6969 6979 27017

CMD ["go-wrapper", "run"] # ["app"]

