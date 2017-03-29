FROM alpine:3.5

MAINTAINER "voidint <voidint@126.com>"

ENV SWAGGER_HUB_HOME /opt/swagger-hub

ENV UI_DOWNLOAD_URL https://github.com/voidint/swagger-hub/raw/master/releases/swagger-hub-ui.tar.gz
ENV SERVER_DOWNLOAD_URL https://github.com/voidint/swagger-hub/raw/master/releases/doc-server-linux-amd64.tar.gz

RUN mkdir -p "$SWAGGER_HUB_HOME" \
    && cd $SWAGGER_HUB_HOME \
    && apk update \
    && apk add ca-certificates wget \
    && wget "$UI_DOWNLOAD_URL" \
    && wget -O ./doc-server.tar.gz "$SERVER_DOWNLOAD_URL" \
    && tar -xz -f ./doc-server.tar.gz \
    && tar -xz -f ./swagger-hub-ui.tar.gz \
    && chown -R root:root ./* \
    && chmod u+x ./doc-server \
    && rm -f doc-server.tar.gz swagger-hub-ui.tar.gz


WORKDIR $SWAGGER_HUB_HOME

EXPOSE 8080

CMD ["./doc-server", "--dir", "/opt/swagger-hub/web", "--domain", "localhost", "--port", "8080"]
