FROM golang:1.22

WORKDIR /usr/src/user_service

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o ./bin/app ./cmd/user/main.go

CMD ["./bin/app"]