# Build Geth in a stock Go builder container
FROM golang:1.12-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

ADD . /XDCchain
RUN cd /XDCchain && make XDC

FROM alpine:latest

LABEL maintainer="anil@xinfin.org"

WORKDIR /XDCchain

COPY --from=builder /XDCchain/build/bin/XDC /usr/local/bin/XDC

RUN chmod +x /usr/local/bin/XDC

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/usr/local/bin/XDC"]

CMD ["--help"]
