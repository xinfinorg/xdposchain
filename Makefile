# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: xdc android ios xdc-cross swarm evm all test clean
.PHONY: xdc-linux xdc-linux-386 xdc-linux-amd64 xdc-linux-mips64 xdc-linux-mips64le
.PHONY: xdc-linux-arm xdc-linux-arm-5 xdc-linux-arm-6 xdc-linux-arm-7 xdc-linux-arm64
.PHONY: xdc-darwin xdc-darwin-386 xdc-darwin-amd64
.PHONY: xdc-windows xdc-windows-386 xdc-windows-amd64

GOBIN = $(shell pwd)/build/bin
GO ?= latest

xdc:
	build/env.sh go run build/ci.go install ./cmd/xdc
	@echo "Done building."
	@echo "Run \"$(GOBIN)/xdc\" to launch xdc."

swarm:
	build/env.sh go run build/ci.go install ./cmd/swarm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/swarm\" to launch swarm."

all:
	build/env.sh go run build/ci.go install

android:
	build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/xdc.aar\" to use the library."

ios:
	build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/xdc.framework\" to use the library."

test: all
	build/env.sh go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/kevinburke/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go get -u github.com/golang/protobuf/protoc-gen-go
	env GOBIN= go install ./cmd/abigen
	@type "npm" 2> /dev/null || echo 'Please install node.js and npm'
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

# Cross Compilation Targets (xgo)

xdc-cross: xdc-linux xdc-darwin xdc-windows xdc-android xdc-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/xdc-*

xdc-linux: xdc-linux-386 xdc-linux-amd64 xdc-linux-arm xdc-linux-mips64 xdc-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/xdc-linux-*

xdc-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/xdc
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/xdc-linux-* | grep 386

xdc-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/xdc
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/xdc-linux-* | grep amd64

xdc-linux-arm: xdc-linux-arm-5 xdc-linux-arm-6 xdc-linux-arm-7 xdc-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/xdc-linux-* | grep arm

xdc-linux-arm-5:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/xdc
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/xdc-linux-* | grep arm-5

xdc-linux-arm-6:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/xdc
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/xdc-linux-* | grep arm-6

xdc-linux-arm-7:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/xdc
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/xdc-linux-* | grep arm-7

xdc-linux-arm64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/xdc
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/xdc-linux-* | grep arm64

xdc-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/xdc
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/xdc-linux-* | grep mips

xdc-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/xdc
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/xdc-linux-* | grep mipsle

xdc-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/xdc
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/xdc-linux-* | grep mips64

xdc-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/xdc
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/xdc-linux-* | grep mips64le 

xdc-darwin: xdc-darwin-386 xdc-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/xdc-darwin-*

xdc-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/xdc
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/xdc-darwin-* | grep 386

xdc-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/xdc
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/xdc-darwin-* | grep amd64

xdc-windows: xdc-windows-386 xdc-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/xdc-windows-*

xdc-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/xdc
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/xdc-windows-* | grep 386

xdc-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/xdc
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/xdc-windows-* | grep amd64
