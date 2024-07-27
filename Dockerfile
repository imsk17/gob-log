# Setup Base Image
FROM alpine:3.13.1 AS base
EXPOSE 4000

FROM golang:1.17-alpine AS builder
RUN apk update
RUN apk add build-base npm
RUN mkdir /build
ADD . /build
# Build the minified CSS Stylesheet
WORKDIR /build
RUN npm i -g yarn
RUN yarn install
RUN yarn release
ENV PORT 4000
RUN go build -o blog -ldflags "-s" main.go

FROM base as FINAL
WORKDIR /app
COPY --from=builder /build/blog .
CMD [ "/app/blog" ]