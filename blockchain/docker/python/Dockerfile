FROM python:3.7

ENV PYTHONUNBUFFERED=1

RUN apt-get update -y --fix-missing \
    && apt-get install -y --no-install-recommends \
    locales

ENV DEBIAN_FRONTEND noninteractive

RUN echo "ja_JP.UTF-8 UTF-8" > /etc/locale.gen \
    && locale-gen ja_JP.UTF-8 \
    && dpkg-reconfigure locales \
    && /usr/sbin/update-locale LANG=ja_JP.UTF-8

ENV LC_ALL=ja_JP.UTF-8
ENV LANG=ja_JP.UTF8

WORKDIR /usr/src/source

ENV PIPENV_VENV_IN_PROJECT=1

COPY ./source/Pipfile ./

RUN pip install --upgrade pip \
    && pip install pipenv \
    && pipenv install --skip-lock --system --dev
