---
name: backend-architect-go
description: Go + Gin + クリーンアーキテクチャ + CQRS バックエンド設計ガイドライン。新規機能追加、リファクタリング、コード生成、アーキテクチャ相談時に使用。
---

# Go Backend Architecture Guidelines

## 技術スタック

Go / Gin / Clean Architecture / CQRS / wire (DI) / sqlc / sqldef / golangci-lint

## ディレクトリ構造

```
.
├── db/
│   ├── migrations/                  # マイグレーションSQL
│   │   ├── schema.sql               # sqldef用スキーマ定義
│   │   └── seed.sql                 # 初期データ（任意）
│   └── sqlc/
│       ├── sqlc.yaml                # sqlc設定
│       ├── queries/                 # SQLクエリ定義
│       │   └── users.sql
│       └── generated/               # sqlc生成コード（自動生成）
│           ├── db.go
│           ├── models.go
│           └── users.sql.go
│
└── src/
    ├── di/
    │   ├── wire.go                  # Provider定義
    │   └── wire_gen.go              # 自動生成（編集禁止）
    │
    ├── domain/                      # ドメイン層（ビジネスロジックの中核）
    │   ├── shared/                  # 汎用Value Object（世間一般で共通のもののみ）
    │   │   ├── created_at_vo.go
    │   │   └── email_vo.go
    │   └── userdm/                  # ドメインモデル（xxxdm = 集約名 + dm）
    │       ├── user_entity.go
    │       ├── user_repository.go
    │       ├── user_id_vo.go
    │       ├── user_name_vo.go
    │       └── is_exist_user_domain_service.go
    │
    ├── usecase/                     # ユースケース層（動詞から始まる）
    │   ├── createuserusecase/
    │   │   ├── create_user_input.go
    │   │   ├── create_user_output.go
    │   │   └── create_user_usecase.go
    │   └── user_query_service.go    # Query系インターフェース
    │
    ├── interface/                   # インターフェース層
    │   ├── controller/
    │   │   └── user_controller.go
    │   └── presentation/
    │       └── user_presenter.go
    │
    └── infra/                       # インフラ層（自由に構成可能）
        ├── router/                  # Ginルーター（FWは技術的詳細）
        │   └── router.go
        ├── rdb/
        │   ├── repoimpl/            # リポジトリ実装（sqlc使用）
        │   │   └── user_repository.go
        │   └── queryimpl/           # QueryService実装（sqlc使用）
        │       └── user_query_service.go
        └── gateway/                 # 外部SaaS連携
            └── sendgrid_gateway.go
```

## Gin ルーター設定

```go
// infra/router/router.go
package router

import (
    "github.com/gin-gonic/gin"
)

func NewRouter(
    userController *controller.UserController,
) *gin.Engine {
    r := gin.Default()

    // ミドルウェア
    r.Use(gin.Recovery())
    r.Use(gin.Logger())

    // API routes
    api := r.Group("/api/v1")
    {
        users := api.Group("/users")
        {
            users.GET("", userController.List)
            users.GET("/:id", userController.Get)
            users.POST("", userController.Create)
            users.PUT("/:id", userController.Update)
            users.DELETE("/:id", userController.Delete)
        }
    }

    return r
}
```

### Controller例（Gin）

```go
// interface/controller/user_controller.go
package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type UserController struct {
    createUserUsecase *createuserusecase.CreateUserUsecase
    userQueryService  usecase.UserQueryService
}

func NewUserController(
    createUserUsecase *createuserusecase.CreateUserUsecase,
    userQueryService usecase.UserQueryService,
) *UserController {
    return &UserController{
        createUserUsecase: createUserUsecase,
        userQueryService:  userQueryService,
    }
}

func (c *UserController) Create(ctx *gin.Context) {
    var req CreateUserRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    input := createuserusecase.CreateUserInput{
        Name:  req.Name,
        Email: req.Email,
    }

    output, err := c.createUserUsecase.Exec(ctx.Request.Context(), input)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{"user_id": output.UserID})
}

func (c *UserController) Get(ctx *gin.Context) {
    id := ctx.Param("id")
    user, err := c.userQueryService.FindByID(ctx.Request.Context(), id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }
    ctx.JSON(http.StatusOK, user)
}
```

## sqlc 設定

### sqlc.yaml

```yaml
# db/sqlc/sqlc.yaml
version: "2"
sql:
  - engine: "postgresql"  # または "mysql"
    queries: "queries/"
    schema: "../migrations/schema.sql"
    gen:
      go:
        package: "generated"
        out: "generated"
        sql_package: "pgx/v5"  # PostgreSQLの場合
        emit_json_tags: true
        emit_empty_slices: true
        emit_result_struct_pointers: true
```

### クエリ定義例

```sql
-- db/sqlc/queries/users.sql

-- name: GetUserByID :one
SELECT id, name, email, created_at, updated_at
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, name, email, created_at, updated_at
FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CreateUser :one
INSERT INTO users (id, name, email, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET name = $2, email = $3, updated_at = $4
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: ExistsUserByEmail :one
SELECT EXISTS(SELECT 1 FROM users WHERE email = $1);
```

### sqlc 実行コマンド

```bash
# インストール
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# コード生成
sqlc generate -f db/sqlc/sqlc.yaml

# 検証
sqlc compile -f db/sqlc/sqlc.yaml
```

### リポジトリ実装（sqlc使用）

```go
// infra/rdb/repoimpl/user_repository.go
package repoimpl

import (
    "context"
    "yourproject/db/sqlc/generated"
    "yourproject/src/domain/userdm"
)

type userRepositoryImpl struct {
    queries *generated.Queries
}

func NewUserRepository(queries *generated.Queries) userdm.UserRepository {
    return &userRepositoryImpl{queries: queries}
}

func (r *userRepositoryImpl) FindByID(ctx context.Context, id userdm.UserID) (*userdm.User, error) {
    row, err := r.queries.GetUserByID(ctx, id.String())
    if err != nil {
        return nil, err
    }
    return r.toEntity(row), nil
}

func (r *userRepositoryImpl) Save(ctx context.Context, user *userdm.User) error {
    _, err := r.queries.CreateUser(ctx, generated.CreateUserParams{
        ID:        user.ID().String(),
        Name:      user.Name().String(),
        Email:     user.Email().String(),
        CreatedAt: user.CreatedAt().Value(),
        UpdatedAt: user.UpdatedAt().Value(),
    })
    return err
}

// toEntity: sqlc生成モデル → ドメインエンティティ変換
func (r *userRepositoryImpl) toEntity(row *generated.User) *userdm.User {
    return userdm.ReconstructUser(
        userdm.NewUserID(row.ID),
        userdm.NewUserName(row.Name),
        shared.NewEmail(row.Email),
        shared.NewCreatedAt(row.CreatedAt),
    )
}
```

## sqldef マイグレーション

### スキーマ定義

```sql
-- db/migrations/schema.sql
CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);

CREATE TABLE posts (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL REFERENCES users(id),
    title VARCHAR(255) NOT NULL,
    body TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_posts_user_id ON posts(user_id);
```

### sqldef 実行コマンド

```bash
# インストール
# PostgreSQL
go install github.com/sqldef/sqldef/cmd/psqldef@latest
# MySQL
go install github.com/sqldef/sqldef/cmd/mysqldef@latest

# dry-run（差分確認）
psqldef -U postgres -h localhost -p 5432 yourdb --dry-run < db/migrations/schema.sql

# 適用
psqldef -U postgres -h localhost -p 5432 yourdb < db/migrations/schema.sql

# MySQL の場合
mysqldef -u root -h localhost yourdb --dry-run < db/migrations/schema.sql
mysqldef -u root -h localhost yourdb < db/migrations/schema.sql
```

### Makefile

```makefile
# Makefile
.PHONY: migrate migrate-dry sqlc lint

DB_HOST ?= localhost
DB_PORT ?= 5432
DB_USER ?= postgres
DB_NAME ?= yourdb

# マイグレーション
migrate:
	psqldef -U $(DB_USER) -h $(DB_HOST) -p $(DB_PORT) $(DB_NAME) < db/migrations/schema.sql

migrate-dry:
	psqldef -U $(DB_USER) -h $(DB_HOST) -p $(DB_PORT) $(DB_NAME) --dry-run < db/migrations/schema.sql

# sqlc
sqlc:
	sqlc generate -f db/sqlc/sqlc.yaml

sqlc-verify:
	sqlc compile -f db/sqlc/sqlc.yaml

# lint
lint:
	golangci-lint run ./...

lint-fix:
	golangci-lint run --fix ./...
```

## CQRS パターン

### Query（参照系）

取得のみのユースケースは **QueryService** パターンを使用。

```go
// usecase/user_query_service.go - インターフェース定義
type UserQueryService interface {
    FindByID(ctx context.Context, id string) (*UserReadModel, error)
    List(ctx context.Context, params ListParams) ([]*UserReadModel, error)
}

// infra/rdb/queryimpl/user_query_service.go - 実装（sqlc使用）
type userQueryServiceImpl struct {
    queries *generated.Queries
}

func NewUserQueryService(queries *generated.Queries) usecase.UserQueryService {
    return &userQueryServiceImpl{queries: queries}
}

func (s *userQueryServiceImpl) FindByID(ctx context.Context, id string) (*usecase.UserReadModel, error) {
    row, err := s.queries.GetUserByID(ctx, id)
    if err != nil {
        return nil, err
    }
    return &usecase.UserReadModel{
        ID:        row.ID,
        Name:      row.Name,
        Email:     row.Email,
        CreatedAt: row.CreatedAt,
    }, nil
}
```

**例外**: 取得のみでも Repository で完結する場合は Repository を使用してOK

### Command（更新系）

永続化を伴うユースケースは **ドメインモデル + Repository** を使用。

### 判断基準

| ユースケース | パターン | 理由 |
|-------------|---------|------|
| 一覧取得・検索 | QueryService | 読み取り専用、ドメインロジック不要 |
| 詳細取得（単純） | Repository | ドメインエンティティをそのまま返す |
| 詳細取得（複雑な集約） | QueryService | 複数テーブル結合、ReadModel返却 |
| 作成・更新・削除 | Usecase + Repository | ドメインロジック・整合性が必要 |

## Usecase層の規約

### 命名規則

- **必ず動詞から始まる**: `createuserusecase`, `updateuserusecase`, `deleteuserusecase`
- ディレクトリで管理し、Input/Output/Usecaseを分離

### 構成ファイル

```go
// usecase/createuserusecase/create_user_input.go
package createuserusecase

type CreateUserInput struct {
    Name  string
    Email string
}

// usecase/createuserusecase/create_user_output.go
package createuserusecase

type CreateUserOutput struct {
    UserID string
}

// usecase/createuserusecase/create_user_usecase.go
package createuserusecase

type CreateUserUsecase struct {
    userRepo        userdm.UserRepository
    isExistUserDS   *userdm.IsExistUserDomainService
}

func NewCreateUserUsecase(
    userRepo userdm.UserRepository,
    isExistUserDS *userdm.IsExistUserDomainService,
) *CreateUserUsecase {
    return &CreateUserUsecase{
        userRepo:      userRepo,
        isExistUserDS: isExistUserDS,
    }
}

func (u *CreateUserUsecase) Exec(ctx context.Context, input CreateUserInput) (*CreateUserOutput, error) {
    // ドメインサービスで存在チェック
    exists, err := u.isExistUserDS.Exec(ctx, userdm.NewUserID(input.Email))
    if err != nil {
        return nil, err
    }
    if exists {
        return nil, errors.New("user already exists")
    }

    // ドメインでエンティティ生成（ドメインルールはここに集約）
    user, err := userdm.NewUser(input.Name, input.Email)
    if err != nil {
        return nil, err
    }

    // 永続化
    if err := u.userRepo.Save(ctx, user); err != nil {
        return nil, err
    }

    return &CreateUserOutput{UserID: user.ID().String()}, nil
}
```

## Domain層の規約

### ディレクトリ命名

- `xxxdm/` （xxx = 集約名、dm = domain model）

### ファイル命名規則

| 種類 | ファイル名 | 例 |
|-----|----------|-----|
| エンティティ | `xxx_entity.go` | `user_entity.go` |
| リポジトリIF | `xxx_repository.go` | `user_repository.go` |
| Value Object | `xxx_vo.go` | `user_id_vo.go`, `user_name_vo.go` |
| ドメインサービス | `xxx_domain_service.go` | `is_exist_user_domain_service.go` |

### Value Object

```go
// domain/userdm/user_id_vo.go
package userdm

type UserID struct {
    value string
}

func NewUserID(value string) UserID {
    return UserID{value: value}
}

func (id UserID) String() string {
    return id.value
}

func (id UserID) Equals(other UserID) bool {
    return id.value == other.value
}
```

### 共通Value Object（shared/）

**世間一般で共通のもののみ** を配置。

```go
// domain/shared/email_vo.go - OK（汎用的）
// domain/shared/created_at_vo.go - OK（汎用的）
// domain/shared/user_id_vo.go - NG！（userの所有物なのでuserdmに置く）
```

他ドメインで `UserID` を使う場合でも、それは `userdm.UserID` として参照する。

### エンティティ

```go
// domain/userdm/user_entity.go
package userdm

type User struct {
    id        UserID
    name      UserName
    email     shared.Email
    createdAt shared.CreatedAt
}

func NewUser(name string, email string) (*User, error) {
    // ドメインルールはここに集約
    userName, err := NewUserName(name)
    if err != nil {
        return nil, err
    }
    emailVO, err := shared.NewEmail(email)
    if err != nil {
        return nil, err
    }

    return &User{
        id:        NewUserID(uuid.New().String()),
        name:      userName,
        email:     emailVO,
        createdAt: shared.NewCreatedAt(time.Now()),
    }, nil
}

// 再構築用（DBからの復元）
func ReconstructUser(id UserID, name UserName, email shared.Email, createdAt shared.CreatedAt) *User {
    return &User{
        id:        id,
        name:      name,
        email:     email,
        createdAt: createdAt,
    }
}

// ゲッターのみ公開（不変性を保つ）
func (u *User) ID() UserID              { return u.id }
func (u *User) Name() UserName          { return u.name }
func (u *User) Email() shared.Email     { return u.email }
```

### リポジトリインターフェース

取得系メソッドは **ドメインかプリミティブ** を返す。

```go
// domain/userdm/user_repository.go
package userdm

type UserRepository interface {
    // 取得系: ドメインを返す
    FindByID(ctx context.Context, id UserID) (*User, error)
    FindByEmail(ctx context.Context, email shared.Email) (*User, error)

    // 取得系: プリミティブを返す（存在確認など）
    ExistsByID(ctx context.Context, id UserID) (bool, error)

    // 永続化系
    Save(ctx context.Context, user *User) error
    Delete(ctx context.Context, id UserID) error
}
```

### ドメインサービス

**DBを使うドメインロジック** はドメインサービスに切り出す。

```go
// domain/userdm/is_exist_user_domain_service.go
package userdm

type IsExistUserDomainService struct {
    userRepo UserRepository
}

func NewIsExistUserDomainService(userRepo UserRepository) *IsExistUserDomainService {
    return &IsExistUserDomainService{userRepo: userRepo}
}

// メソッドは1つだけ（Exec）
func (ds *IsExistUserDomainService) Exec(ctx context.Context, userID UserID) (bool, error) {
    user, err := ds.userRepo.FindByID(ctx, userID)
    if err != nil {
        return false, err
    }
    return user != nil, nil
}
```

**ドメインサービスの特徴**:
- Repositoryインターフェースを経由してデータ取得
- `Exec` メソッド1つのみ
- 命名: `is_xxx_domain_service.go`, `can_xxx_domain_service.go` など

## DI（依存性注入）- wire

Google Wire を使用してコンパイル時にDIを解決。

### ディレクトリ構造

```
src/
└── di/
    ├── wire.go           # Provider定義
    └── wire_gen.go       # 自動生成（編集禁止）
```

### wire.go

```go
//go:build wireinject
// +build wireinject

package di

import (
    "github.com/google/wire"
    "github.com/jackc/pgx/v5"
    "yourproject/db/sqlc/generated"
    "yourproject/src/domain/userdm"
    "yourproject/src/usecase/createuserusecase"
    "yourproject/src/infra/rdb/repoimpl"
    "yourproject/src/infra/rdb/queryimpl"
    "yourproject/src/interface/controller"
    "yourproject/src/interface/router"
)

// Provider Sets
var infraSet = wire.NewSet(
    repoimpl.NewUserRepository,
    wire.Bind(new(userdm.UserRepository), new(*repoimpl.UserRepositoryImpl)),
    queryimpl.NewUserQueryService,
)

var domainServiceSet = wire.NewSet(
    userdm.NewIsExistUserDomainService,
)

var usecaseSet = wire.NewSet(
    createuserusecase.NewCreateUserUsecase,
)

var controllerSet = wire.NewSet(
    controller.NewUserController,
)

var routerSet = wire.NewSet(
    router.NewRouter,
)

// InitializeRouter はアプリケーション全体のDIを解決
func InitializeRouter(queries *generated.Queries) *router.Router {
    wire.Build(
        infraSet,
        domainServiceSet,
        usecaseSet,
        controllerSet,
        routerSet,
    )
    return nil
}
```

### main.go

```go
// main.go
package main

import (
    "context"
    "os"

    "github.com/jackc/pgx/v5"
    "yourproject/db/sqlc/generated"
    "yourproject/src/di"
)

func main() {
    // DB接続
    db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
    if err != nil {
        panic(err)
    }
    defer db.Close(context.Background())

    // sqlc Queries
    queries := generated.New(db)

    // wire で DI 解決
    r := di.InitializeRouter(queries)

    // サーバー起動
    r.Run(":8080")
}
```

### wire コマンド

```bash
# インストール
go install github.com/google/wire/cmd/wire@latest

# コード生成（src/di ディレクトリで実行）
cd src/di && wire

# または
wire ./src/di/...
```

### Makefile に追加

```makefile
# wire
wire:
	cd src/di && wire

# 全生成
generate: sqlc wire
```

### Provider の書き方

```go
// 単純なProvider
func NewUserRepository(queries *generated.Queries) *UserRepositoryImpl {
    return &UserRepositoryImpl{queries: queries}
}

// インターフェースへのバインド
wire.Bind(new(userdm.UserRepository), new(*repoimpl.UserRepositoryImpl))

// 複数のProviderをセットにまとめる
var infraSet = wire.NewSet(
    NewUserRepository,
    wire.Bind(new(userdm.UserRepository), new(*repoimpl.UserRepositoryImpl)),
    NewUserQueryService,
)
```
```

## golangci-lint 設定

`.golangci.yml` をプロジェクトルートに配置。

```yaml
run:
  timeout: 5m
  go: "1.21"

linters:
  enable:
    # デフォルト有効
    - errcheck
    - gosimple
    - govet
    - ineffassign
    - staticcheck
    - unused

    # 追加推奨
    - goimports
    - gofmt
    - misspell
    - unconvert
    - unparam
    - nakedret
    - prealloc
    - exportloopref
    - nilerr
    - errorlint
    - gosec
    - bodyclose
    - sqlclosecheck
    - contextcheck

linters-settings:
  goimports:
    local-prefixes: github.com/your-org/your-project
  errcheck:
    check-type-assertions: true
    check-blank: true
  govet:
    enable-all: true
  nakedret:
    max-func-lines: 30
  gosec:
    excludes:
      - G104

issues:
  exclude-rules:
    - path: _test\.go
      linters:
        - errcheck
        - gosec
    # sqlc生成コードは除外
    - path: db/sqlc/generated/
      linters:
        - all
  max-issues-per-linter: 0
  max-same-issues: 0
```

## 禁止事項

- domain層から infra層への依存
- usecase層から interface層への依存
- エンティティへのセッター追加
- ディレクトリ名に `command/` `query/` を使用
- `user_id_vo.go` を shared/ に置く（ドメイン固有のVOは各xxxdmに）
- any型の使用
- golangci-lint のエラーを無視してコミット
- sqlc生成コードの手動編集
- wire_gen.go の手動編集
- 手動DI（wireを使用すること）
