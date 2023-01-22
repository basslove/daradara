ARG GO_IMAGE_NAME

# develop
FROM ${GO_IMAGE_NAME} as dev

ARG SERVICE
ARG APP_NAME

ENV WORK_DIR=/${APP_NAME}/
ENV LANG C.UTF-8
ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk update && apk add git

RUN mkdir -p ${WORK_DIR}
WORKDIR ${WORK_DIR}

COPY go.mod go.sum ./
# COPY ./ ./
RUN go mod download

CMD ["go", "run", "cmd/daradara/main.go"]

# Build stage
FROM ${GO_IMAGE_NAME} as builder

ARG SERVICE
ARG APP_NAME

ENV WORK_DIR=/${APP_NAME}/
ENV LANG C.UTF-8
ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk add --update-cache git ca-certificates tzdata make gcc g++ && git config --global http.sslVerify false

RUN mkdir -p ${WORK_DIR}
WORKDIR ${WORK_DIR}

COPY go.mod go.sum ./
COPY cmd ./cmd
COPY db ./db
COPY internal ./internal
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ${WORK_DIR}/main ./cmd/daradara/main.go

## Run staging
#FROM alpine:3.17
#
#ENV LANG C.UTF-8
#ENV ROOT=/go/src/daradara
#RUN mkdir -p ${ROOT}
#WORKDIR ${ROOT}
#
#COPY --from=builder ${ROOT}/main .
#COPY .env .
#
#CMD ["/go/src/daradara/main"]
