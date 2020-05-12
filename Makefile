CLI_VERSION=1.0.0

build-go:
	bash go-executable-build.bash github.com/mattcanty/isp-monitor

build-image: build-go
	docker build -t isp-monitor:prerelease . --build-arg CLI_VERSION=$(CLI_VERSION) --build-arg LINUX_ARCH=x86_64

run:
	docker run -v ~/.config/isp-monitor:/etc/isp-monitor/config isp-monitor:prerelease

all: build-image
