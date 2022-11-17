FROM golang:1.16-alpine as builder

ENV BIN_FILE /opt/banner-rotation-service/banner-rotation-service
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} ${CODE_DIR}/cmd/banner-rotation-service/

FROM alpine:3

LABEL SERVICE="banner-rotation-service"
LABEL MAINTAINERS="feride3d@icloud.com"

ENV BIN_FILE "/opt/banner-rotation-service/banner-rotation-service"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

ENV CONFIG_FILE /etc/banner-rotation-service/config.yml
COPY ./configs/config.yml ${CONFIG_FILE}

CMD ${BIN_FILE} -config ${CONFIG_FILE}