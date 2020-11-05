VERSION=`git describe --tags --abbrev=0`
VERSION_FLAG=main.Version=$(VERSION)

all:
	go build -ldflags "-X ${VERSION_FLAG} -s -w" -o easi-trm-gateway ./...
