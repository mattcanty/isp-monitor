FROM alpine

ARG CLI_VERSION
ARG LINUX_ARCH

RUN apk add wget
RUN wget "https://bintray.com/ookla/download/download_file?file_path=ookla-speedtest-${CLI_VERSION}-${LINUX_ARCH}-linux.tgz" -O speedtest.tgz
RUN tar -x -f ./speedtest.tgz

COPY isp-monitor-linux-amd64 isp-monitor

ENV CONFIG_ROOT=/etc/isp-monitor/config

CMD ./isp-monitor
