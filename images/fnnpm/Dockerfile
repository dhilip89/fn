# build stage
FROM golang:1.10-alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ENV D=/go/src/github.com/fnproject/fn
ADD . $D
RUN cd $D/poolmanager/server && go build -o fnnpm && cp fnnpm /tmp/

# final stage: the local fnproject/dind:latest will be either built afresh or
# whatever is the latest from master, depending on whether we're releasing
# a newer cut.
FROM alpine:latest
WORKDIR /app
COPY --from=build-env /tmp/fnnpm /app/fnnpm
CMD ["./fnnpm"]
