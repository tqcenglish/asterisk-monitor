buildDateTime = $(shell date '+%Y-%m-%d %H:%M:%S')
gitCommitCode = $(shell git rev-parse --short HEAD)
gitBranch = $(shell git rev-parse --abbrev-ref HEAD)
goVersion = $(shell go version)
version = $(shell cat scripts/version.conf |grep softVersion=|cut -d '=' -f 2)

run: build
	./deployments/asterisk-monitor  web -c ./configs/config.yaml
run-ui:
	cd frontend && npm dev
build-ui:
	rm -rf web/www
	cd frontend && npm run build
build:
	go build -tags "admin pprof" -o ./deployments/asterisk-monitor main.go
build-arm:
	GOOS=linux GOARCH=arm64 go build -tags "admin pprof" -o ./deployments/asterisk-monitor main.go
