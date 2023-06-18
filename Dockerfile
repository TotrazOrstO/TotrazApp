FROM golang:1.19-alpine
WORKDIR $GOPATH/src/totrazapp

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone

COPY go.mod go.mod .env ./
COPY ./ .
RUN go get -d ./...
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ./totrazapp ./cmd/main.go


CMD ["./totrazapp"]
EXPOSE 2641