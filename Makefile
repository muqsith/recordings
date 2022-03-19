SHELL := /bin/bash

.PHONY: frontend-build
frontend-build:
	source ${NVM_DIR}/nvm.sh \
	&& cd http-pub && nvm use && npm run build

.PHONY: build
build:
	cd server && go build -o ../bin/recordings -tags=jsoniter

.PHONY: build-app
build-app: frontend-build build