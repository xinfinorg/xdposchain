
FROM golang:1.10-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers
ADD . /xdcchain
RUN cd /xdcchain && make xdc

FROM alpine:latest

LABEL maintainer="ino@xinfin.org"
WORKDIR /xdcchain
COPY --from=builder /xdcchain/build/bin/xdc /usr/local/bin/xdc

RUN chmod +x /usr/local/bin/xdc

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/usr/local/bin/xdc", "--help"]
