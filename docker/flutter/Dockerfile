FROM ubuntu:18.04

RUN apt-get update \
    && apt-get install -y \
    curl \
    wget \
    gnupg \
    less \
    lsof \
    net-tools \
    git \
    unzip \
    apt-utils \
    apt-transport-https \
    xz-utils

# DART
# https://dart.dev/get-dart#install
RUN sh -c "curl https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add -" \
    && sh -c "curl https://storage.googleapis.com/download.dartlang.org/linux/debian/dart_stable.list > /etc/apt/sources.list.d/dart_stable.list" \
    && apt-get update \
    && apt-get install -y dart
ENV PATH="${PATH}:/usr/lib/dart/bin/:/root/.pub-cache/bin"

# https://dart.dev/tutorials/web/get-started#3-get-cli-tools-or-an-ide-or-both
RUN pub global activate webdev \
    && pub global activate stagehand

# FLUTTER
# https://flutter.dev/docs/get-started/install/linux
ARG FLUTTER_VERSION
RUN cd /usr/local/bin \
    && wget https://storage.googleapis.com/flutter_infra/releases/stable/linux/${FLUTTER_VERSION}-stable.tar.xz \
    && tar xf ${FLUTTER_VERSION}-stable.tar.xz \
    && rm -f ${FLUTTER_VERSION}-stable.tar.xz
ENV PATH="${PATH}:/usr/local/bin/flutter/bin"
RUN flutter precache

# web support
# https://flutter.dev/docs/get-started/web
RUN flutter channel master \
    && flutter upgrade \
    && flutter config --enable-web
