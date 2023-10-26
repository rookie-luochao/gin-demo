FROM private-harbor.xxx.com/xxx/builder:v1.0.0 AS build-env

FROM build-env AS builder


ARG GOPROXY=https://goproxy.cn,direct

WORKDIR /go/src
COPY ./ ./

# build
RUN make build WORKSPACE=demo-docker

# runtime
FROM private-harbor.xxx.com/xxx/runtime:v1.0.0

COPY --from=builder /go/src/cmd/demo-docker/demo-docker /go/bin/demo-docker


EXPOSE 9090

EXPOSE 80

ARG PROJECT_NAME
ARG PROJECT_VERSION
ENV PROJECT_NAME=${PROJECT_NAME} PROJECT_VERSION=${PROJECT_VERSION}

WORKDIR /go/bin
ENTRYPOINT ["/go/bin/demo-docker"]
