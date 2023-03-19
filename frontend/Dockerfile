FROM golang:1.19.7-bullseye
SHELL ["/bin/bash", "-c"]
COPY ./frontend /app/frontend
COPY ./common /app/common
WORKDIR /app/frontend/cmd
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct && go install .
ENV HTTP_PORT=8080
ENTRYPOINT exec /go/bin/cmd --httpPort ${HTTP_PORT}
EXPOSE ${HTTP_PORT}