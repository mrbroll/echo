FROM alpine:3.7

ADD ./server /usr/bin/

CMD ["/usr/bin/server"]
