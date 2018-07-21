
build_linux_amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -o release/linux/amd64/telo

build_linux_i386:
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -v -a -o release/linux/i386/telo

docker:
	docker build -t lhong/telo .

test:
	go test -v .
