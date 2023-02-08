# sardine

LMS(Leaning Management System)のAPI等を提供するためのリポジトリ

## APIドキュメント

https://ae.gitlab.kaonavi.jp/sardine/main/oas/index.html

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

---

## MinIOのAccessKey/SecretKeyを発行

開発の際にAccessKey/SecretKeyが必要になった場合は以下の手順でキーを発行してください。

### MinIOのコンソールにログイン

```sh
$ open http://localhost:29001
```

※ID/Passwordは `minioadmin` です

### キーを発行

1. サイドメニューの `Access Keys` を押下
2. `Create access key` を押下
3. `Create` を押下
4. 表示されるキーを保持する
   - ローカル専用でいつでも再発行できるのでDownloadは不要
