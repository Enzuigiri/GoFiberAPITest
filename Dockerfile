FROM golang:1.20-alpine

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN cd ./src && go build -o /server-test

EXPOSE 8000

ENTRYPOINT [ "/server-test" ]