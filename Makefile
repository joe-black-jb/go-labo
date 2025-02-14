REVISION := $(shell git rev-parse --short HEAD)
# REVISION := "1.0.0"
# -X main.revision=$(REVISION) を "" で囲わないと go build が main.revision=$(REVISION) を .go ファイルだと認識し、エラーになる
# main.revision と main.Revision は違うので注意
LD_FLAGS := "-X main.Revision=$(REVISION)"

# build-mac の元々のコマンド
# -ldflags $(LD_FLAGS) だとタイミング的に変数が展開されないので、-ldflags "$(LD_FLAGS)" に変更
# GOOS=darwin GOARCH=amd64 go build -ldflags $(LD_FLAGS) -o ./binary_mac ./main.go
# REVISION を直接指定するパターン
# GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.revision=$(REVISION)" -o ./binary_mac ./main.go

flags:
	@echo $(REVISION)
	@echo $(LD_FLAGS)

run:
	go run .

build-mac:
	rm -f ./binary_mac
	GOOS=darwin GOARCH=amd64 go build -ldflags $(LD_FLAGS) -o ./binary_mac ./main.go

