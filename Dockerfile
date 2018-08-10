# Golang build container
FROM golang:1.10

WORKDIR $GOPATH/src/github.com/logdisplayplatform/logdisplayplatform

COPY Gopkg.toml Gopkg.lock ./
COPY vendor vendor

ARG DEP_ENSURE=""
RUN if [ ! -z "${DEP_ENSURE}" ]; then \
      go get -u github.com/golang/dep/cmd/dep && \
      dep ensure --vendor-only; \
    fi

COPY pkg pkg
COPY build.go build.go
COPY package.json package.json

RUN go run build.go build

# Node build container
FROM node:8

WORKDIR /usr/src/app/

COPY package.json yarn.lock ./
RUN yarn install --pure-lockfile --no-progress

COPY Gruntfile.js tsconfig.json tslint.json ./
COPY public public
COPY scripts scripts
COPY emails emails

ENV NODE_ENV production
RUN ./node_modules/.bin/grunt build

# Final container
FROM debian:stretch-slim

ARG GF_UID="472"
ARG GF_GID="472"

ENV PATH=/usr/share/logdisplayplatform/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin \
    GF_PATHS_CONFIG="/etc/logdisplayplatform/logdisplayplatform.ini" \
    GF_PATHS_DATA="/var/lib/logdisplayplatform" \
    GF_PATHS_HOME="/usr/share/logdisplayplatform" \
    GF_PATHS_LOGS="/var/log/logdisplayplatform" \
    GF_PATHS_PLUGINS="/var/lib/logdisplayplatform/plugins" \
    GF_PATHS_PROVISIONING="/etc/logdisplayplatform/provisioning"

WORKDIR $GF_PATHS_HOME

RUN apt-get update && apt-get install -qq -y libfontconfig ca-certificates && \
    apt-get autoremove -y && \
    rm -rf /var/lib/apt/lists/*

COPY conf ./conf

RUN mkdir -p "$GF_PATHS_HOME/.aws" && \
    groupadd -r -g $GF_GID logdisplayplatform && \
    useradd -r -u $GF_UID -g logdisplayplatform logdisplayplatform && \
    mkdir -p "$GF_PATHS_PROVISIONING/datasources" \
             "$GF_PATHS_PROVISIONING/dashboards" \
             "$GF_PATHS_LOGS" \
             "$GF_PATHS_PLUGINS" \
             "$GF_PATHS_DATA" && \
    cp "$GF_PATHS_HOME/conf/sample.ini" "$GF_PATHS_CONFIG" && \
    cp "$GF_PATHS_HOME/conf/ldap.toml" /etc/logdisplayplatform/ldap.toml && \
    chown -R logdisplayplatform:logdisplayplatform "$GF_PATHS_DATA" "$GF_PATHS_HOME/.aws" "$GF_PATHS_LOGS" "$GF_PATHS_PLUGINS" && \
    chmod 777 "$GF_PATHS_DATA" "$GF_PATHS_HOME/.aws" "$GF_PATHS_LOGS" "$GF_PATHS_PLUGINS"

COPY --from=0 /go/src/github.com/logdisplayplatform/logdisplayplatform/bin/linux-amd64/logdisplayplatform-server /go/src/github.com/logdisplayplatform/logdisplayplatform/bin/linux-amd64/logdisplayplatform-cli ./bin/
COPY --from=1 /usr/src/app/public ./public
COPY --from=1 /usr/src/app/tools ./tools
COPY tools/phantomjs/render.js ./tools/phantomjs/render.js

EXPOSE 3000

COPY ./packaging/docker/run.sh /run.sh

USER logdisplayplatform
ENTRYPOINT [ "/run.sh" ]
