NAME	:=	sfcc_exporter
GO		:=	go
FMT		=	gofmt
pkgs	=	$(shell env GO111MODULE=on $(GO) list -m)
API_URL =	https://store.ubi.com

FILE	=	main.go\
			metrics_module.go\
			collector.go\
			const_var.go\
			type_struct.go\
			collector_promotion.go\
			get_token.go\
			get_sfcc_metrics.go\
			collector_coupon.go

DOCKER_IMAGE_NAME       ?= sfcc_exporter

all: format build

test:
	@echo ">> running tests"
	@go test -short $(pkgs)

format:
	@echo ">> formatting code"
	@$(FMT) -w $(FILE)

module:
	@echo ">> creating module"
	@$(GO) mod init $(NAME)

build: 
	@echo ">> building binaries"
	@$(GO) build -o $(NAME)

docker: all
	@echo ">> building docker image"
	@docker build -t $(DOCKER_IMAGE_NAME) .

fclean:
	rm -rf $(NAME) go.sum go.mod

re: fclean module all test

.PHONY: all format build test
