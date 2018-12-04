FROM golang:latest 

ADD . /go/src/github.com/golang/parse_article
RUN go get -u github.com/golang/protobuf/proto
RUN go get -u google.golang.org/grpc
RUN go install github.com/golang/protobuf/protoc-gen-go
COPY wrapper.sh /go/src/github.com/golang/parse_article/wrapper.sh
RUN chmod 755 /go/src/github.com/golang/parse_article/wrapper.sh

CMD ["/go/src/github.com/golang/parse_article/wrapper.sh"]
EXPOSE 50051 8080