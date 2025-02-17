SHELL := /bin/bash

# 基础配置
APP_NAME ?= khan
DOCKER_IMAGE ?= khan
VERSION ?= $(shell git rev-parse --short HEAD)
GO_MODULE = smallBot/cmd
GO_PATH := $(shell go env GOPATH)
GO_ROOT := $(shell go env GOROOT)
DOWNLOAD_PATH := $(shell realpath ../download)
HOST_DOWNLOAD ?= $(shell [ -d $(DOWNLOAD_PATH) ] || mkdir -p $(DOWNLOAD_PATH); echo $(DOWNLOAD_PATH))
CONTAINER_DOWNLOAD := /app/public/download
HOST_PORT ?= 8074
CONTAINER_PORT ?= 8073
LOG_FILE ?= /var/log/supervisor/supervisord.log

.PHONY: all all_mac_arm build docker-build docker-build-mac-arm docker-run docker-run-mac-arm clean help env demo

all: docker-build
	@echo "镜像构建成功，开始运行容器..."
	@make docker-run
	@echo "容器运行成功，访问地址: http://localhost:${HOST_PORT}"

all_mac_arm: docker-build-mac-arm
	@echo "镜像构建成功，开始运行容器..."
	@make docker-run-mac-arm
	@echo "容器运行成功，访问地址: http://localhost:${HOST_PORT}"

build:
	@echo "编译二进制文件..."
	CGO_ENABLED=0 go build -ldflags="-s -w" -o ${APP_NAME} ${GO_MODULE}

docker-build:
	@echo "构建Docker镜像..."
	docker build \
		--build-arg APP_NAME=${APP_NAME}_${VERSION} \
		-t ${DOCKER_IMAGE}:${VERSION} \
		.

docker-build-mac-arm:
	@echo "构建苹果系统的arm64镜像..."
	docker buildx build \
		--platform linux/arm64 \
		--build-arg APP_NAME=${APP_NAME}_mac_${VERSION} \
		-t ${DOCKER_IMAGE}_mac:${VERSION} \
		.

docker-run:
	@echo "开始运行容器..."
	@echo "使用的下载目录: ${DOWNLOAD_PATH}"
	@echo "容器中使用的下载目录: ${CONTAINER_DOWNLOAD}"
	docker run -d \
		--name ${APP_NAME}_${VERSION} \
		-p ${HOST_PORT}:${CONTAINER_PORT} \
		-v ${DOWNLOAD_PATH}:${CONTAINER_DOWNLOAD} \
		${DOCKER_IMAGE}:${VERSION}

docker-run-mac-arm:
	@echo "开始运行容器..."
	@echo "使用的下载目录: ${DOWNLOAD_PATH}"
	@echo "容器中使用的下载目录: ${CONTAINER_DOWNLOAD}"
	docker run -d \
		--name ${APP_NAME}_mac_${VERSION} \
		-p ${HOST_PORT}:${CONTAINER_PORT} \
		-v ${DOWNLOAD_PATH}:${CONTAINER_DOWNLOAD} \
		${DOCKER_IMAGE}_mac:${VERSION}

clean:
	@echo "清理构建产物..."
	rm -f ${APP_NAME}
	@echo "清理容器..."
	docker rm -f ${APP_NAME}_${VERSION} || true
	docker rm -f ${APP_NAME}_mac_${VERSION} || true
	@echo "清理Docker镜像..."
	docker rmi ${DOCKER_IMAGE}:${VERSION} || true
	docker rmi ${DOCKER_IMAGE}_mac:${VERSION} || true
	@echo "清理Docker镜像缓存..."
	docker image prune -f
	@echo "清理覆盖率报告..."
	rm -f coverage.out coverage.html

install-deps:
	@echo "安装依赖..."
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

fmt:
	@echo "格式化代码..."
	${GO_PATH}/bin/goimports -w .
	${GO_ROOT}/bin/gofmt -s -w .

lint:
	@echo "运行代码静态检测..."
	${GO_PATH}/bin/golangci-lint run ./...

test:
	@echo "运行测试用例..."
	go test -v -coverprofile=coverage.out ./...

coverage:
	@echo "生成覆盖率报告..."
	go tool cover -html=coverage.out -o coverage.html

logs:
	@echo "查看容器日志..."
	docker exec ${APP_NAME} tail -f ${LOG_FILE}

env:
	@echo "环境变量:"
	@echo "APP_NAME: $(APP_NAME)"
	@echo "DOCKER_IMAGE: $(DOCKER_IMAGE)"
	@echo "VERSION: $(VERSION)"
	@echo "GOPATH: $(GO_PATH)"
	@echo "GOROOT: $(GO_ROOT)"
	@echo "GO_MODULE: $(GO_MODULE)"
	@echo "SHELL: $(SHELL)"
	@echo "DOWNLOAD_PATH: $(DOWNLOAD_PATH)"
	@echo "HOST_DOWNLOAD: $(HOST_DOWNLOAD)"
	@echo "CONTAINER_DOWNLOAD: $(CONTAINER_DOWNLOAD)"
	@echo "HOST_PORT: $(HOST_PORT)"
	@echo "CONTAINER_PORT: $(CONTAINER_PORT)"
	@echo "LOG_FILE: $(LOG_FILE)"

demo:
	@echo "demo测试..."
	go run cmd/main.go demo

help:
	@echo "可执行的命令:"
	@echo "  all          - 编译本地二进制文件并构建Docker镜像"
	@echo "  all_mac_arm       - 编译本地二进制文件并构建苹果系统的arm64镜像"
	@echo "  build        - 编译本地二进制文件"
	@echo "  docker-build - 构建Docker镜像"
	@echo "  docker-build-mac-arm - 构建苹果系统的arm64镜像"
	@echo "  docker-run   - 运行Docker容器"
	@echo "  docker-run-mac-arm - 运行苹果系统的arm64镜像"
	@echo "  test         - 运行测试"
	@echo "  logs         - 查看容器日志"
	@echo "  clean        - 清理构建产物"
	@echo "  fmt          - 格式化代码"
	@echo "  lint         - 代码静态检查"
	@echo "  coverage     - 生成覆盖率报告"
	@echo "  install-deps - 安装开发依赖"
	@echo "  env          - 查看环境变量"
	@echo "  demo         - demo测试"