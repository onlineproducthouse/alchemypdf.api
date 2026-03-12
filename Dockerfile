# syntax=docker/dockerfile:1

ARG IMAGE_REGISTRY_BASE_URL

FROM ${IMAGE_REGISTRY_BASE_URL}/golang:1.25.5-alpine AS builder

LABEL maintainer="onlineproducthouse <info@onlineproducthouse.com>"

RUN mkdir -p /app

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN go install github.com/swaggo/swag/cmd/swag

FROM builder

ARG TARGETOS
ARG TARGETARCH

ENV ALCHEMYPDF_BIN_FOLDER=./bin/${TARGETOS}/${TARGETARCH}

COPY ./lib/alchemypdf.api.constant ./lib/alchemypdf.api.constant
COPY ./lib/alchemypdf.api.contract ./lib/alchemypdf.api.contract
COPY ./lib/alchemypdf.api.infrastructure ./lib/alchemypdf.api.infrastructure
COPY ./lib/alchemypdf.api.ioc ./lib/alchemypdf.api.ioc
COPY ./lib/alchemypdf.api.model ./lib/alchemypdf.api.model
COPY ./lib/alchemypdf.api.service ./lib/alchemypdf.api.service
COPY ./lib/alchemypdf.api.testhelpers ./lib/alchemypdf.api.testhelpers

COPY ./services/alchemypdf.api.app ./services/alchemypdf.api.app

RUN $(go env GOPATH)/bin/swag init --parseDependency -g services/alchemypdf.api.app/app/app.go

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o ${ALCHEMYPDF_BIN_FOLDER}/app ./services/alchemypdf.api.app/main.go

ENTRYPOINT [ "sh", "-c", "$(echo ${ALCHEMYPDF_BIN_FOLDER}/app)" ]
