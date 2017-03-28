FROM golang:1.8.0
MAINTAINER Adam Pahlevi

# installing cachable go dependency
RUN go get -u github.com/kardianos/govendor
RUN go get github.com/onsi/ginkgo/ginkgo
RUN go get github.com/onsi/gomega

# add the current banku dir
ADD . /go/src/banku
WORKDIR /go/src/banku

# installing the dependency
RUN govendor sync

# run with: docker build -t banku .
# connect: docker run -i -t banku /bin/bash
