FROM golang:1.18-alpine3.15
WORKDIR /go/src

ENV PATH="$PATH:/bin/bash" \
    BENTO4_BIN="/opt/bento4/bin" \
    PATH="$PATH:/opt/bento4/bin"

RUN apk add --update bash build-base curl

RUN curl --output bento4.zip https://www.bok.net/Bento4/binaries/Bento4-SDK-1-6-0-639.x86_64-unknown-linux.zip
RUN unzip -d bento bento4.zip
RUN mkdir "/opt/bento4"
RUN cp -R bento/Bento4-SDK-1-6-0-639.x86_64-unknown-linux/bin /opt/bento4

#vamos mudar para o endpoint correto. Usando top apenas para segurar o processo rodando
ENTRYPOINT [ "top"]