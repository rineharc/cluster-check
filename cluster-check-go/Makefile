# Build targets
TARGETS = \
    cluster-check-windows-amd64.exe \
    cluster-check-linux-amd64 \
    cluster-check-darwin-amd64 \
    cluster-check-darwin-arm64

# Go build flags
GOFLAGS = -ldflags="-s -w"

# Build rules
all: $(TARGETS)

cluster-check-windows-amd64.exe:
	GOOS=windows GOARCH=amd64 go build $(GOFLAGS) -o $@

cluster-check-linux-amd64:
	GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -o $@

cluster-check-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build $(GOFLAGS) -o $@

cluster-check-darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build $(GOFLAGS) -o $@

clean:
	rm -f $(TARGETS)
