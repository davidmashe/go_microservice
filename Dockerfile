FROM golang:1.9

ENV GO_MICROSERVICE_PORT=3025

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
COPY . .

RUN go get github.com/go-martini/martini 
RUN go get golang.org/x/crypto/bcrypt

CMD ["go", "run", "server.go"]