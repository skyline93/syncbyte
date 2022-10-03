FROM golang:1.18

COPY output/syncbyte-agent /usr/bin/syncbyte-agent
COPY output/syncbyte-engine /usr/bin/syncbyte-engine
COPY output/syncbyte /usr/bin/syncbyte
COPY conf/agent.yaml /root/.syncbyte-agent.yaml
COPY conf/engine.yaml /root/.syncbyte-engine.yaml
COPY conf/syncbyte.yaml /root/.syncbyte.yaml

RUN mkdir /var/log/syncbyte \
    && mkdir /var/run/backup \
    && mkdir /var/run/restore \
    && ln -s /usr/bin/syncbyte-agent /usr/bin/agent \
    && ln -s /usr/bin/syncbyte-engine /usr/bin/engine

WORKDIR /

VOLUME [ "/var/log/syncbyte", "/var/run/backup", "/var/run/restore" ]
