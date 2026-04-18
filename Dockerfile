FROM alpine:latest

RUN apk add --no-cache dante-server bash

COPY sockd.conf /etc/sockd.conf
COPY start.sh /start.sh
RUN chmod +x /start.sh

EXPOSE 1080

CMD ["/start.sh"]
