# Flutter tutorial

- [hello_world](https://flutter.dev/docs/get-started/codelab#step-1-create-the-starter-flutter-app)
- [startup_namer](https://flutter.dev/docs/get-started/codelab#step-1-create-the-starter-flutter-app)
- [building_layout](https://flutter.dev/docs/development/ui/layout/tutorial)

## Requirement

- [Docker](https://www.docker.com/)
  - docker-compose

## Usage

Run docker container:

```console
$ docker-compose up -d
```

Create new app:

```console
$ docker-compose exec flutter bash
> flutter create [APP_NAME]
> cd [APP_NAME]
```

Run app:

```console
> flutter pub get
> flutter run -d web-server --web-port=8080 --web-hostname=0.0.0.0
```

Go to `http://localhost:8080` and you'll see the web page.

## Install

Clone repository:

```console
$ git clone https://github.com/PiroHiroPiro/til.git
$ cd til/flutter_tutorial
```

Build image:

```console
$ docker-compose build
```

## Author

[Hiroyuki Nishizawa](https://github.com/PiroHiroPiro)
