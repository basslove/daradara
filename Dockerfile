FROM golang:1.16.3-alpine3.13 as builder

ARG GITHUB_TOKEN
ARG SERVICE_NAME
ARG VERSION

ENV GOPRIVATE=github.com

RUN apk --no-cache add git

WORKDIR /app

COPY ./go.* ./

RUN git config --global url."https://$x-access-token:${GITHUB_TOKN}@github.com/".insteadOf "https://github.com/"

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go install -v -ldflags="-w -s -X main.version=${VERSION}" -X main.serviceName=${SERVICE_NAME} ./cmd/${SERVICE_NAME}

FROM alpine:3.13
RUN apk --no-cache add ca-certificates
COPY --from-builder /go/bin/${SERVICE_NAME} /app/${SERVICE_NAME}

RUN addgroup -g 1001 loc && adduser -D -G loc -u 1001 loc
USER 1001

CMD ["/app/${SERVICE_NAME}"]

