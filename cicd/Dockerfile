FROM golang:1.14 as builder

RUN apt-get update && apt-get install -y git build-essential

COPY . /builder
RUN cd /builder && make && mv /builder/build/bin/XDC /builder/build/bin/XDC-mainnet

RUN mv /builder/common/constants/constants.go.devnet /builder/common/constants.go
RUN cd /builder && make && mv /builder/build/bin/XDC /builder/build/bin/XDC-devnet

RUN mv /builder/common/constants/constants.go.testnet /builder/common/constants.go
RUN cd /builder && make && mv /builder/build/bin/XDC /builder/build/bin/XDC-testnet

# The actual image
FROM alpine:3

WORKDIR /work

COPY --from=builder /builder/build/bin/XDC-devnet /usr/bin
COPY --from=builder /builder/build/bin/XDC-testnet /usr/bin
COPY --from=builder /builder/build/bin/XDC-mainnet /usr/bin

RUN apk add --no-cache bash
RUN chmod +x /usr/bin/XDC*
    
# Copy over files
ADD cicd/devnet /work/devnet
ADD cicd/testnet /work/testnet
ADD cicd/mainnet /work/mainnet
ADD cicd/entry.sh /work/

# Create an empty pwd file
RUN touch /work/.pwd

# rpc
EXPOSE 8545
# ws
EXPOSE 8555
# port
EXPOSE 30303

ENTRYPOINT ["bash","/work/entry.sh"]