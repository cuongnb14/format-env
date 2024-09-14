build:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o bin/fenv_darwin_arm64 fenv.go \
	&& CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/fenv_darwin_amd64 fenv.go
