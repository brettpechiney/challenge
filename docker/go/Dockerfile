FROM golang as builder

# setup
RUN apt-get update -y
RUN apt-get upgrade -y

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o service

FROM scratch
COPY --from=builder /app/service /app/
COPY --from=builder /app/application-properties.toml /app/
EXPOSE 8083
ENTRYPOINT ["/app/service"]