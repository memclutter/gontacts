FROM golang

RUN go get github.com/gin-gonic/gin \
 && go get gopkg.in/mgo.v2 \
 && go get github.com/pkg/errors \
 && go get github.com/dgrijalva/jwt-go \
 && go get github.com/pkg/profile \
 && go get golang.org/x/crypto/bcrypt \
 && go get gopkg.in/gomail.v2

ADD . /go/src/github.com/memclutter/gontacts

RUN go install github.com/memclutter/gontacts

ENTRYPOINT /go/bin/gontacts

EXPOSE 8000
