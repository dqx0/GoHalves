# ベースイメージを指定します。Golangを使用します。
FROM golang:latest
COPY backend /go/src/app/go/

RUN mkdir -p /usr/local/go/src/web/app

WORKDIR /go/src/app/go/

RUN go mod tidy
#&& apk add --no-cache git \
#&& go get github.com/gin-gonic/gin \
#&& go get github.com/jinzhu/gorm \
#&& go get github.com/lib/pq \

#&& go mod init app \
#&& apk add git

EXPOSE 8080

CMD ["go", "run", "main.go"]