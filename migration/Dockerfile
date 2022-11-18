FROM golang:1.16-alpine as builder

ENV MIGRATION_DIR /migration
ENV CODE_DIR /go/src/
WORKDIR ${CODE_DIR}
RUN go get -u github.com/pressly/goose/v3/cmd/goose
COPY migration ${MIGRATION_DIR}

## Wait script for migration
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

CMD /wait && goose -dir ${MIGRATIONS_DIR} postgres "user=postgres password=postgres host=localhost port=5432 dbname=banner-rotation-service sslmode=disable" up 