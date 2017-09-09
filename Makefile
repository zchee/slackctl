GO_SRCS = $(shell find . -type f \( -name '*.go' -and -not -iwholename '*testdata*' \))

build: bin/slackctl

bin/slackctl: bin $(GO_SRCS)
	go build -v -x -o $@ ./cmd/slackctl

bin:
	@mkdir ./bin

debug:
	@echo ${GO_SRCS}

.PHONY: build debug
