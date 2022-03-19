.PHONY: build
build:
	cd server && go build -o ../bin/recordings -tags=jsoniter