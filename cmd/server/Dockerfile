# syntax=docker/dockerfile:1
FROM golang AS builder
ARG GITHUB_USER_BUILD
ARG GITHUB_TOKEN_BUILD
ARG timezone=Asia/Bangkok
ENV GITHUB_USER $GITHUB_USER_BUILD
ENV GITHUB_TOKEN $GITHUB_TOKEN_BUILD
ENV TZ $timezone
WORKDIR /app

# install dumb-init
RUN apt-get update && \
    apt-get -y install dumb-init && \
    apt-get -y clean

# for CI/CD with private repo
# RUN go env -w GOPRIVATE=github.com/thn-lee

# use ssh instead of https
# RUN git config --global url."git@github.com:".insteadOf "https://github.com"

# use https with github token
# RUN git config --global url."https://${GITHUB_USER}:${GITHUB_TOKEN}@github.com".insteadOf "https://github.com"

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.* ./
RUN go mod download && go mod verify

# let's build project
COPY . .
RUN CGO_ENABLED=0 go build -v \
    -buildvcs=false \
    -installsuffix 'static' \
    -ldflags="-X 'main.version=$(git rev-parse --short HEAD)' -X 'main.build=$(date --iso-8601=seconds)'" \
    -o dist/server ./cmd/server

# pack PRD image
FROM gcr.io/distroless/base:nonroot
LABEL maintainer="Thanawat Leelawatcharamas <thanawat.lee@kkumail.com>"

ARG timezone=Asia/Bangkok

ENV LANG C.UTF-8
ENV LC_ALL C.UTF-8
ENV TZ $timezone

# Create app dir
WORKDIR /app
COPY --from=builder /usr/bin/dumb-init /usr/bin/dumb-init
COPY --from=builder /app/dist/server /app/server
COPY ./configs /app/configs
COPY ./internal/assets /app/internal/assets
COPY ./web /app/web

EXPOSE 8080 8443

# default run entrypoint
ENTRYPOINT ["/usr/bin/dumb-init", "--", "/app/server"]
CMD ["--env=prd"]
