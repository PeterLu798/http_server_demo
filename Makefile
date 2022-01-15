export tag=v1.0
root: 
	export ROOT=github.com/PeterLu798/http_server_demo

build: 
    echo "building http_server_demo binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64

release: build
	echo "building http_server_demo container"
	docker build -t cncamp/httpserver:${tag} .

push: release
	echo "pushing cncamp/httpserver"
	docker push cncamp/httpserver:${tag}
