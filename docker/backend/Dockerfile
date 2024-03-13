# ベースイメージを指定します。Golangを使用します。
FROM golang:latest
COPY ../../backend /go/src/app/go/

WORKDIR /go/src/app/go/

RUN go mod tidy

EXPOSE 8080

CMD ["go", "run", "."]