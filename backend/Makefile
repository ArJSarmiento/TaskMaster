build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/server

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose