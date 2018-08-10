FROM debian:stretch-slim

ARG GRAFANA_TGZ="logdisplayplatform-latest.linux-x64.tar.gz"

RUN apt-get update && apt-get install -qq -y tar && \
    apt-get autoremove -y && \
    rm -rf /var/lib/apt/lists/*

COPY ${GRAFANA_TGZ} /tmp/logdisplayplatform.tar.gz

RUN mkdir /tmp/logdisplayplatform && tar xfvz /tmp/logdisplayplatform.tar.gz --strip-components=1 -C /tmp/logdisplayplatform

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

COPY --from=0 /tmp/logdisplayplatform "$GF_PATHS_HOME"

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

EXPOSE 3000

COPY ./run.sh /run.sh

USER logdisplayplatform
ENTRYPOINT [ "/run.sh" ]
