clean:
	rm -Rfv bin
	mkdir bin

build: clean
	go build -o bin/covid-exporter main.go

build-all: clean
	GOOS="linux"   GOARCH="amd64"       go build -o bin/covid-exporter__linux-amd64 main.go
	GOOS="linux"   GOARCH="arm" GOARM=6 go build -o bin/covid-exporter__linux-armv6 main.go
	GOOS="linux"   GOARCH="arm" GOARM=7 go build -o bin/covid-exporter__linux-armv7 main.go
	GOOS="linux"   GOARCH="arm"         go build -o bin/covid-exporter__linux-arm   main.go
	GOOS="darwin"  GOARCH="amd64"       go build -o bin/covid-exporter__macos-amd64 main.go
	GOOS="windows" GOARCH="amd64"       go build -o bin/covid-exporter__win-amd64 main.go
