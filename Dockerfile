###############################################################################
# https://github.com/roboll/getsss
###############################################################################
FROM alpine:3.2

ADD target/getsss-linux-amd64 /getsss
ENTRYPOINT ["/getsss"]
