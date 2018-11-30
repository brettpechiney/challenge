FROM golang as builder

# setup
RUN apt-get update -y
RUN apt-get upgrade -y

# install wget
RUN apt-get install wget

# install CockroachDB
RUN wget -qO- https://binaries.cockroachdb.com/cockroach-v2.1.1.linux-amd64.tgz | tar  xvz
RUN cp -i cockroach-v2.1.1.linux-amd64/cockroach /usr/local/bin

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

FROM scratch
COPY --from=builder /app/server /app/
EXPOSE 8083
ENTRYPOINT ["/app/server"]