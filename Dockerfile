FROM alpine:3.14

COPY artifact/stats stats
COPY config config
CMD ./stats