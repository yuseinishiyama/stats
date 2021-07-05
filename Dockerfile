FROM alpine:3.14

COPY artifact/stats bin/stats
COPY config config
CMD stats