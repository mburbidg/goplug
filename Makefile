.DEFAULT_GOAL := build

build:
	go build -o bin/goplug github.com/mburbidg/goplug
	go build -buildmode=plugin -o bin/plugins/plugin_a.so github.com/mburbidg/goplug/plugin_a
	go build -buildmode=plugin -o bin/plugins/plugin_b.so github.com/mburbidg/goplug/plugin_b
.PHONY:build
