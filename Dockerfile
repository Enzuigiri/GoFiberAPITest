FROM golang:alpine

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /server-test

EXPOSE 8000

ENTRYPOINT [ "/server-test" ]
