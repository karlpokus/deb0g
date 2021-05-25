FROM alpine:3.13.5
RUN apk add --no-cache ca-certificates
COPY bin/server /src/server
ENTRYPOINT ["/src/server"]
