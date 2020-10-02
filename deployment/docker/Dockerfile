
# First Step: from full golang toolchain container
FROM golang:alpine AS build-env

# declare working directory
ENV BOJ_WORKING_DIR $GOPATH/src/github.com/Myriad-Dreamin/boj-v6
# expose go mod env
ENV GO111MODULE on
# both https://goproxy.{io, cn} are ok
ENV GOPROXY https://goproxy.cn

RUN go version
RUN go env

# checkout working directory
WORKDIR $BOJ_WORKING_DIR

# Copy dependencies to container, without other codes
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Add source codes
ADD . .

# Build codes
RUN env go build -v -o ./server ./cmd/server

RUN mv server /boj-server-bin

# Second Step: from alpine application container
FROM amd64/alpine

# set alpine apk source to ustc mirror
# RUN sed -i 's/archive.ubuntu.com/mirrors.ustc.edu.cn/g' /etc/apt/sources.list

WORKDIR /
RUN apk add -U tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime
COPY --from=build-env /boj-server-bin /server
ADD docs/main_spec.json docs/
COPY deployment/docker/run.sh /run.sh
COPY config.toml /config.toml
# COPY rbac.conf /rbac.conf
RUN chmod +x /run.sh
EXPOSE 23336
ENTRYPOINT "/run.sh"
