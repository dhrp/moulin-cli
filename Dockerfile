FROM golang

ENV GOPATH=/go/
COPY . $GOPATH/src/github.com/dhrp/moulincli
WORKDIR $GOPATH/src/github.com/dhrp/moulincli
RUN go get -u github.com/golang/dep/cmd/dep
RUN make build
RUN go install
