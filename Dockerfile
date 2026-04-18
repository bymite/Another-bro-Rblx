FROM alpine:latest

RUN apk add --no-cache git make gcc musl-dev

RUN git clone https://github.com/rofl0r/microsocks && \
    cd microsocks && make && cp microsocks /usr/local/bin/ && \
    cd .. && rm -rf microsocks

EXPOSE 1080

CMD ["microsocks", "-p", "1080"]
