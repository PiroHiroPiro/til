# UNIX command

システムプログラミング

## Requirement

- [Docker](https://www.docker.com/)

## Try it

```console
$ docker run -w /usr/src -it system_programming /bin/bash
/usr/src# gcc -o fileinfo fileinfo.c
/usr/src# ./fileinfo fileinfo.c
   mode: 100644
  links: 1
   user: 0
  group: 0
   size: 996
modtime: 1559240549
   name: fileinfo.c
```

## Install

Build images:

```console
$ docker build -t system_programming .
```
