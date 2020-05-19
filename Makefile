DIRS=build build/bins
CLI_VERSION=1.0.0

build-go:
	go build -o build/bins/isp-monitor-local github.com/mattcanty/isp-monitor
	GOOS=linux GOARCH=arm GOARM=5 go build -o build/bins/isp-monitor-linux-arm-5 github.com/mattcanty/isp-monitor
	

build-image: build-go
	docker build -t isp-monitor:prerelease . --build-arg CLI_VERSION=$(CLI_VERSION) --build-arg LINUX_ARCH=x86_64

run:
	docker run -v ~/.config/isp-monitor:/etc/isp-monitor/config isp-monitor:prerelease

publish-go:
	scp isp-monitor-linux-amd64

build: build-image

$(shell mkdir -p $(DIRS))
