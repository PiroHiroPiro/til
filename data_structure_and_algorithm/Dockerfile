FROM python:3.6

RUN apt-get update
RUN apt-get -y install locales && \
    localedef -f UTF-8 -i ja_JP ja_JP.UTF-8

ENV LANG=ja_JP.UTF-8
ENV LANGUAGE=ja_JP:ja
ENV LC_ALL=ja_JP.UTF-8
ENV TZ=JST-9
ENV PYTHONUNBUFFERED=1
ENV PIPENV_VENV_IN_PROJECT=1

RUN mkdir /usr/src/source

WORKDIR /usr/src/source

COPY ./Pipfile  ./
COPY ./Pipfile.lock ./

RUN pip install --upgrade pip
RUN pip install --upgrade setuptools
RUN pip install pipenv
RUN pipenv install --system --dev
