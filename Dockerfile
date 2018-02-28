FROM alpine:3.7

ADD ./echo /usr/bin/

CMD ["/usr/bin/echo"]
