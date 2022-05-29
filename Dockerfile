FROM golang:1.18 AS builder
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download -x

COPY . ./
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-X 'main.release=`git rev-parse --short=8 HEAD`'" -o /bin/server cmd/server/*.go

FROM ubuntu
WORKDIR /app

COPY --from=builder /bin/server ./
COPY --from=builder /src/truffle.config ./
COPY --from=builder /src/wait-for-it.sh ./

# CMD ["./server"]
