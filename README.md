# sardine

LMS(Leaning Management System)のAPI等を提供するためのリポジトリ

## APIドキュメント

https://ae.gitlab.kaonavi.jp/sardine/master/oas/index.html

※警告は無視してください。


## 環境構築手順

### 前提条件

- shironegiの環境構築が終わっていること

### 1. .envの作成

```sh
$ cp .env.example .env
```

`.env` の以下のキーにshironegiの構築時に設定したDB名と同じ値（通常は社員番号）を設定。

```
DB_DATABASE=[社員番号]
MYSQL_DATABASE=[社員番号]
```

### 2. docker構築

```sh
$ docker compose up -d --build
```

### 3. DBのマイグレーション

```sh
$ docker compose exec go make migrate
```

