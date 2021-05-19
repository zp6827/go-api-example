FROM golang:1.8.5-jessie
RUN go get github.com/IncSW/geoip2
COPY src/testdata/GeoLite2-Country.mmdb src/testdata/GeoLite2-Country.mmdb
WORKDIR /go/src
ADD src/ src/
CMD ["go", "run", "src/main.go"]
