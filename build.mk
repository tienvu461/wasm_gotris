.PHONY: build clean build_gotris

NAME := main.wasm
GOTRIS := gotris.wasm
ENV := local
GOOS := js
GOARCH := wasm
HASH := $$(git rev-parse --short --verify HEAD)
DATE := $$(date -u '+%Y%m%dT%H%M%S')
GOVERSION = $$(go version)
PUBLIC_DIR := ./public

build: $(NAME).$(ENV).$(GOOS).$(GOARCH)

$(NAME).$(ENV).$(GOOS).$(GOARCH):
	GOOS=$(GOOS) GOARCH=$(GOARCH) \
	    go build -tags=$(ENV) \
	        -o $(NAME) \
	        .
	# cp $(NAME) $@
	mv $(NAME) $(PUBLIC_DIR)/$(NAME)

# build_gotris: $(GOTRIS).$(ENV).$(GOOS).$(GOARCH)
#
# $(GOTRIS).$(ENV).$(GOOS).$(GOARCH):
# 	GOOS=$(GOOS) GOARCH=$(GOARCH) \
# 	    go build -tags=$(ENV) \
# 	        -o $(GOTRIS) \
# 	        ./tetris/
# 	# cp $(GOTRIS) $@
# 	mv $(GOTRIS) $(PUBLIC_DIR)/$(GOTRIS)

clean:
	-rm -rf $(NAME).$(ENV).$(GOOS).$(GOARCH) $(GOTRIS).$(ENV).$(GOOS).$(GOARCH)
