BUILD_DIR=./bin
BUILD_NAME=llc

default: build

build:
	@go build -o $(BUILD_DIR)/$(BUILD_NAME) .
