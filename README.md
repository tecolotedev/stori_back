# Local Setup

## Create a postgres db with docker

```bash
$ docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=secret -d postgres:15.0-alpine

$ docker exec -it postgres15 createdb --username=admin --owner=admin stori_db
```

## Drop database if needed

```bash
$ docker exec -it postgres15 dropdb --username=admin stori_db
```

## Install golang-migrate

_docs:_ https://github.com/golang-migrate/migrate?tab=readme-ov-file

**mac:**

```bash
$ brew install golang-migrate
```

**windows**

```bash
$ scoop install migrate
```

**linux**

```bash
$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
$ apt-get update
$ apt-get install -y migrate
```

## Migrate tables with golang-migrate

```bash
$ migrate -path db/migrations -database "postgresql://admin:secret@127.0.0.1:5432/stori_db?sslmode=disable" --verbose up
```

## Install sqlc

```bash
$ brew install sqlc
```
