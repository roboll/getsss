###############################################################################
# https://github.com/roboll/getsss
###############################################################################
FROM alpine:3.2

RUN apk --update add ca-certificates && rm -rf /var/cache/apk/*

ADD target/getsss-linux-amd64 /getsss
ENTRYPOINT ["/getsss"]
