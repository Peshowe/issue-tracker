FROM golang:1.15.5-alpine AS base
WORKDIR /src
COPY go.* ./
RUN go mod download

FROM base AS build
COPY . .
RUN go build -o /out/tracker-service


CMD ["/out/tracker-service"]