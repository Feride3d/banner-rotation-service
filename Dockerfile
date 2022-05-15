# Собираем в гошке
FROM golang:1.16 as build

ENV BIN_FILE /opt/banner-rotation-service
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

# Кэшируем слои с модулями
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

# Собираем статический бинарник Go (без зависимостей на Си API),
# иначе он не будет работать в alpine образе.
ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/banner-rotation-service/*

# На выходе тонкий образ
FROM alpine:3.9

LABEL SERVICE="banner-rotation-service"
LABEL MAINTAINERS="feride3d@icloud.com"

ENV BIN_FILE "/opt/banner-rotation-service"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

ENV CONFIG_FILE /etc/banner-rotation-service/config.yml
COPY ./configs/config.yml ${CONFIG_FILE}

CMD ${BIN_FILE} -config ${CONFIG_FILE}