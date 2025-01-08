FROM golang:1.23-alpine AS builder

WORKDIR /build

COPY go.mod go.sum main.go ./

COPY cmd ./cmd

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o nordvpn-docker

FROM ubuntu:24.04 AS runtime

RUN apt-get update && \
apt-get install -y --no-install-recommends wget apt-transport-https ca-certificates && \
    apt-get install -y --no-install-recommends wget apt-transport-https ca-certificates && \ 
    wget -qO /etc/apt/trusted.gpg.d/nordvpn_public.asc https://repo.nordvpn.com/gpg/nordvpn_public.asc && \
    echo "deb https://repo.nordvpn.com/deb/nordvpn/debian stable main" > /etc/apt/sources.list.d/nordvpn.list && \
    apt-get update && \
    apt-get install -y --no-install-recommends nordvpn && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /build/nordvpn-docker /usr/local/bin/nordvpn-docker

ENTRYPOINT ["/usr/local/bin/nordvpn-docker"]
