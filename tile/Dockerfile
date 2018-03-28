FROM golang:1.8
RUN mkdir -p /go/src/tile38/tile38ser

WORKDIR /go/src/tile38/tile38ser
COPY . /go/src/tile38/tile38ser

# RUN go-wrapper download   # "go get -d -v ./..."
# RUN go-wrapper install    # "go install -v ./..."

ENV PORT 9851
EXPOSE 9851


CMD ["./tile38-server"] # ["app"]