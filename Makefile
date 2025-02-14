# REVISION := $(shell git rev-parse --short HEAD)
REVISION := "1.0.0"
# -X main.revision=$(REVISION) を "" で囲わないと go build が main.revision=$(REVISION) を .go ファイルだと認識し、エラーになる
LD_FLAGS := "-X main.revision=$(REVISION)"

run:
	go run .

build-mac:
	GOOS=darwin GOARCH=amd64 go build -ldflags $(LD_FLAGS) -o ./binary_mac ./main.go