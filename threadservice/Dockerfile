FROM golang:1.19.7-bullseye
SHELL ["/bin/bash", "-c"]
COPY ./threadservice /app/threadservice
COPY ./common /app/common
WORKDIR /app/threadservice
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct && go install .

ENTRYPOINT exec /go/bin/threadservice --rpcPort ${RPC_PORT} --httpPort ${HTTP_PORT} --address ${ADDRESS}
EXPOSE ${RPC_PORT}
EXPOSE ${HTTP_PORT}