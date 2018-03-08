FROM golang

RUN go get github.com/gin-gonic/gin \
 && go get gopkg.in/mgo.v2

ADD . /go/src/github.com/memclutter/gontacts

RUN go install github.com/memclutter/gontacts

ENTRYPOINT /go/bin/gontacts

EXPOSE 8000
