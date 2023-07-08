FROM golang:1.18.7-alpine3.15 as builder

WORKDIR /app

# copy modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on

# cache modules
RUN go mod download

# copy source code
COPY cmd/ cmd/
COPY pkg/ pkg/

COPY n1/ n1/
COPY n2/ n2/
COPY n3/ n3/

COPY n1.yaml n1.yaml
COPY n2.yaml n2.yaml
COPY n3.yaml n3.yaml

# build
RUN CGO_ENABLED=0 go build \
    -a -o kv-app cmd/main.go

FROM alpine:3.13

RUN apk --no-cache add ca-certificates

USER nobody

COPY --from=builder --chown=nobody:nobody /app/kv-app .

ENTRYPOINT ["./kv-app"]