PWD := $(shell pwd)
GOPATH := $(shell go env GOPATH)
PKG_NAME := "sled"

clean:
	@echo "Cleaning..."
	@rm -f ./sled
	@rm -rf ./sled-*
	@rm -rf ./*.tar.gz
	@rm -rf ./sled_*
	@rm -rf ./*.txt
	@rm -rf ./*.pem
	@echo "Done cleaning..."

fmt:
	@echo "Running $@"
	@go fmt *.go

test-tls:
	@echo "Making Development TLS Certificates..."
	@openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem -subj "/C=US/ST=Texas/L=Austin/O=Local Development/OU=IT Department/CN=127.0.0.0"

travis-sizes:
	@echo "Building unstripped binary..."
	@go build -o sled-raw || (echo "Failed to build binary: $$?"; exit 1)
	@echo "Size of unstripped binary: $$(ls -l sled-raw | awk '{print $$5}') bytes or $$(ls -lh sled-raw | awk '{print $$5}')" > ./size-report.txt
	@echo "Building stripped binary..."
	@go build -ldflags="-s -w" -o sled-stripped || (echo "Failed to build stripped binary: $$?"; exit 1)
	@echo "Size of stripped binary: $$(ls -l sled-stripped | awk '{print $$5}') bytes or $$(ls -lh sled-stripped | awk '{print $$5}')" >> ./size-report.txt
	@echo "Compressing stripped binary..."
	@cp ./sled-stripped ./sled-compressed
	@upx -9 -q ./sled-compressed > /dev/null || (echo "Failed to compress stripped binary: $$?"; exit 1)
	@echo "Size of compressed stripped binary: $$(ls -l sled-compressed | awk '{print $$5}') bytes or $$(ls -lh sled-compressed | awk '{print $$5}')" >> ./size-report.txt
	@echo "Reported binary sizes for Go version $$(echo -n $$(go version) | grep -o '1\.[0-9|\.]*'): "
	@cat ./size-report.txt
	@rm -f ./*.txt
