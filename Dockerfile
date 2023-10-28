FROM golang:1.20-bullseye AS build-env

FROM build-env AS builder


ARG GOPROXY=https://goproxy.cn,direct

WORKDIR /go/src
COPY ./ ./

# build
RUN make build WORKSPACE=demo-docker

# runtime
FROM debian:11

COPY --from=builder /go/src/cmd/demo-docker/demo-docker /go/bin/demo-docker

COPY --from=builder /go/src/cmd/demo-docker/openapi.json /go/bin/cmd/demo-docker/openapi.json
		
EXPOSE 9090

EXPOSE 80

ARG PROJECT_NAME
ARG PROJECT_VERSION
ENV PROJECT_NAME=${PROJECT_NAME} PROJECT_VERSION=${PROJECT_VERSION}

WORKDIR /go/bin
ENTRYPOINT ["/go/bin/demo-docker"]
