run:
	go run .

build: linux windows

linux:
	GOOS=linux GOARCH=amd64 go build -o ./build/webshare .

windows:
	GOOS=windows GOARCH=amd64 go build -o ./build/webshare.exe .
