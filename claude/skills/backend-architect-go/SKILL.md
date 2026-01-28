---
name: backend-architect-go
description: Go + Gin + クリーンアーキテクチャ + CQRS バックエンド設計ガイドライン。新規機能追加、リファクタリング、コード生成、アーキテクチャ相談時に使用。
---

# Go Backend Architecture Guidelines

## 技術スタック

Go / Gin / Clean Architecture / CQRS / wire (DI) / sqlc / sqldef / golangci-lint / gomock / testify / goldentest

## ディレクトリ構造

```
.
├── cmd/
│   └── api/
│       └── main.go                  # エントリーポイント
├── db/
│   └── migrations/schema.sql        # sqldef用スキーマ
└── internal/
    ├── di/
    │   ├── wire.go                  # Provider定義
    │   └── wire_gen.go              # 自動生成（編集禁止）
    ├── mock/                        # gomock生成（internal構造をミラー）
    │   ├── domain/userdm/mock_user_repository.go
    │   └── usecase/mock_user_query_service.go
    ├── domain/
    │   ├── shared/                  # 汎用VO（世間一般で共通のもののみ）
    │   │   ├── email_vo.go
    │   │   └── created_at_vo.go
    │   └── userdm/                  # 集約（xxxdm = 集約名 + dm）
    │       ├── user_entity.go       # コンストラクタはprivate
    │       ├── user_factory.go      # public生成メソッド
    │       ├── user_repository.go   # //go:generate mockgen
    │       ├── user_id_vo.go        # ※userdm固有、sharedに置かない
    │       ├── user_name_vo.go
    │       └── is_exist_user_domain_service.go
    ├── usecase/
    │   ├── createuserusecase/       # 必ず動詞から始まる
    │   │   ├── create_user_input.go
    │   │   ├── create_user_output.go
    │   │   └── create_user_usecase.go
    │   └── user_query_service.go    # Query系IF（//go:generate）
    ├── interface/
    │   ├── controller/                  # 1Usecase = 1Controllerで凝集度を高める
    │   │   ├── createusercontroller/
    │   │   │   └── create_user_controller.go
    │   │   └── getusercontroller/
    │   │       └── get_user_controller.go
    │   └── presentation/user_presenter.go
    └── infra/
        ├── router/router.go         # Gin（FWは技術的詳細なのでinfra）
        ├── rdb/
        │   ├── sqlc.yaml
        │   ├── generated/           # sqlc自動生成（編集禁止）
        │   ├── queries/             # メソッドごとにファイル分割（コンフリクト回避）
        │   │   ├── get_user_by_id.sql
        │   │   ├── create_user.sql
        │   │   └── exists_user_by_email.sql
        │   ├── repoimpl/user_repository.go      # Repository実装
        │   └── queryimpl/user_query_service.go  # QueryService実装
        └── gateway/                 # 外部SaaS（SendGrid, SQS等）
```

## 命名規則

| 種類 | ファイル名 | 例 |
|-----|----------|-----|
| エンティティ | `xxx_entity.go` | `user_entity.go` |
| ファクトリー | `xxx_factory.go` | `user_factory.go` |
| リポジトリIF | `xxx_repository.go` | `user_repository.go` |
| Value Object | `xxx_vo.go` | `user_id_vo.go` |
| ドメインサービス | `xxx_domain_service.go` | `is_exist_user_domain_service.go` |

## CQRS

| ユースケース | パターン | 実装場所 |
|-------------|---------|---------|
| 一覧・検索（複雑） | QueryService | usecase/にIF → infra/rdb/queryimpl/に実装 |
| 詳細取得（単純） | Repository | Repositoryで完結する場合はRepositoryでOK |
| 作成・更新・削除 | Usecase + Domain + Repository | ドメインロジックを経由 |

## Domain層

### Value Object
```go
// domain/userdm/user_id_vo.go
type UserID struct { value string }
func NewUserID() UserID { return UserID{value: uuid.New().String()} }
func (id UserID) String() string { return id.value }
func (id UserID) Equals(other UserID) bool { return id.value == other.value }
```

### エンティティ（コンストラクタはprivate）
```go
// domain/userdm/user_entity.go
type User struct { id UserID; name UserName; email shared.Email; createdAt shared.CreatedAt }
func newUser(id UserID, name UserName, email shared.Email, createdAt shared.CreatedAt) *User {
    return &User{id: id, name: name, email: email, createdAt: createdAt}  // private
}
func (u *User) ID() UserID { return u.id }  // ゲッターは名詞のみ（GetXXX禁止）、セッターは必要時のみ
```

### ファクトリー（ユースケース別にpublic生成メソッド）
```go
// domain/userdm/user_factory.go
func GenUserForCreate(name, email string) (*User, error) {
    userName, err := NewUserName(name)  // バリデーションあり
    if err != nil { return nil, err }
    emailVO, err := shared.NewEmail(email)
    if err != nil { return nil, err }
    return newUser(NewUserID(), userName, emailVO, shared.NewCreatedAt(time.Now())), nil
}
func GenUserForReconstruct(id, name, email string, createdAt time.Time) *User {
    return newUser(UserID{value: id}, UserName{value: name}, shared.Email{value: email}, shared.CreatedAt{value: createdAt})
}
```

### リポジトリIF（取得系はドメインかプリミティブを返す）
```go
// domain/userdm/user_repository.go
//go:generate mockgen -source=$GOFILE -destination=../../mock/domain/userdm/mock_user_repository.go -package=mockuserdm
type UserRepository interface {
    FindByID(ctx context.Context, id UserID) (*User, error)
    FindByEmail(ctx context.Context, email shared.Email) (*User, error)
    ExistsByID(ctx context.Context, id UserID) (bool, error)  // プリミティブ返却もOK
    Save(ctx context.Context, user *User) error
    Delete(ctx context.Context, id UserID) error
}
```

### ドメインサービス（DBを使うドメインロジック）
```go
// domain/userdm/is_exist_user_domain_service.go
type IsExistUserDomainService struct { userRepo UserRepository }
func NewIsExistUserDomainService(repo UserRepository) *IsExistUserDomainService { return &IsExistUserDomainService{userRepo: repo} }
func (ds *IsExistUserDomainService) Exec(ctx context.Context, email shared.Email) (bool, error) {
    user, err := ds.userRepo.FindByEmail(ctx, email)
    if err != nil { return false, err }
    return user != nil, nil
}
// ※ Execメソッド1つのみ。命名例: IsExistXxx, CanXxx
```

### shared/の配置ルール
- **OK**: `email_vo.go`, `created_at_vo.go`（世間一般で汎用的）
- **NG**: `user_id_vo.go`（userの所有物 → userdmに置く。他ドメインからは`userdm.UserID`で参照）

## Usecase層

```go
// usecase/createuserusecase/create_user_usecase.go
type CreateUserUsecase struct {
    userRepo      userdm.UserRepository
    isExistUserDS *userdm.IsExistUserDomainService
}
func (u *CreateUserUsecase) Exec(ctx context.Context, input CreateUserInput) (*CreateUserOutput, error) {
    exists, _ := u.isExistUserDS.Exec(ctx, shared.NewEmail(input.Email))
    if exists { return nil, errors.New("user already exists") }
    user, err := userdm.GenUserForCreate(input.Name, input.Email)  // ファクトリー使用
    if err != nil { return nil, err }
    if err := u.userRepo.Save(ctx, user); err != nil { return nil, err }
    return &CreateUserOutput{UserID: user.ID().String()}, nil
}
```

## Infra層

### sqlc設定
```yaml
# internal/infra/rdb/sqlc.yaml
version: "2"
sql:
  - engine: "postgresql"
    queries: "queries/"
    schema: "../../../db/migrations/schema.sql"
    gen:
      go:
        package: "generated"
        out: "generated"
        sql_package: "pgx/v5"
        emit_json_tags: true
```

### sqlcクエリ定義（1ファイル1メソッド、コンフリクト回避）
```sql
-- internal/infra/rdb/queries/get_user_by_id.sql
-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- internal/infra/rdb/queries/create_user.sql
-- name: CreateUser :one
INSERT INTO users (id, name, email, created_at) VALUES ($1, $2, $3, $4) RETURNING *;

-- internal/infra/rdb/queries/exists_user_by_email.sql
-- name: ExistsUserByEmail :one
SELECT EXISTS(SELECT 1 FROM users WHERE email = $1);
```

### sqldef（マイグレーション）
```bash
psqldef -U postgres -h localhost yourdb --dry-run < db/migrations/schema.sql  # 差分確認
psqldef -U postgres -h localhost yourdb < db/migrations/schema.sql            # 適用
# MySQL: mysqldef -u root -h localhost yourdb < db/migrations/schema.sql
```

## Interface/Controller層

### 凝集度を高めるためのController設計

**1Usecase = 1Controller** のパターンを採用。これにより：
- コントローラーの依存が明確になる
- 単体テストが容易になる
- 変更の影響範囲が限定される

```go
// interface/controller/createusercontroller/create_user_controller.go
package createusercontroller

type CreateUserController struct {
    createUserUsecase *createuserusecase.CreateUserUsecase
}

func NewCreateUserController(uc *createuserusecase.CreateUserUsecase) *CreateUserController {
    return &CreateUserController{createUserUsecase: uc}
}

func (c *CreateUserController) Handle(ctx *gin.Context) {
    var req CreateUserRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        _ = ctx.Error(apperr.BadRequestWrap(err, apperr.CodeInvalidRequest))
        return
    }
    output, err := c.createUserUsecase.Exec(ctx.Request.Context(), createuserusecase.CreateUserInput{
        Name:  req.Name,
        Email: req.Email,
    })
    if err != nil {
        _ = ctx.Error(err)
        return
    }
    ctx.JSON(http.StatusCreated, gin.H{"id": output.UserID})
}
```

### Ginルーター
```go
// infra/router/router.go
func NewRouter(
    createUserCtrl *createusercontroller.CreateUserController,
    getUserCtrl *getusercontroller.GetUserController,
) *gin.Engine {
    r := gin.Default()
    api := r.Group("/api/v1")
    users := api.Group("/users")
    users.GET("/:id", getUserCtrl.Handle)
    users.POST("", createUserCtrl.Handle)
    return r
}
```

### リポジトリ実装
```go
// infra/rdb/repoimpl/user_repository.go
func (r *userRepositoryImpl) FindByID(ctx context.Context, id userdm.UserID) (*userdm.User, error) {
    row, err := r.queries.GetUserByID(ctx, id.String())
    if err != nil { return nil, err }
    return userdm.GenUserForReconstruct(row.ID, row.Name, row.Email, row.CreatedAt), nil  // ファクトリー使用
}
```

## DI（wire）※手動DI禁止

```go
// internal/di/wire.go
//go:build wireinject
package di

var infraSet = wire.NewSet(
    repoimpl.NewUserRepository,
    wire.Bind(new(userdm.UserRepository), new(*repoimpl.UserRepositoryImpl)),
    queryimpl.NewUserQueryService,
)
var domainServiceSet = wire.NewSet(userdm.NewIsExistUserDomainService)
var usecaseSet = wire.NewSet(createuserusecase.NewCreateUserUsecase)

// Controller providers（1Usecase = 1Controller）
var controllerSet = wire.NewSet(
    createusercontroller.NewCreateUserController,
    getusercontroller.NewGetUserController,
)

func InitializeRouter(queries *generated.Queries) *gin.Engine {
    wire.Build(infraSet, domainServiceSet, usecaseSet, controllerSet, router.NewRouter)
    return nil
}
```

## テスト（必須）※テストなしはマージ禁止

**errorは絶対に無視しない。require.NoError等で必ずチェック。**

| 層 | 種類 | 方法 |
|----|------|------|
| domain | 単体テスト | gomock + testify |
| usecase | 単体テスト | gomock + testify |
| interface/controller | 単体テスト | gomock + testify |
| interface/presentation | 不要 | - |
| infra/rdb | 単体テスト | SQLite（RDB制約で無理なら実RDB） |
| infra/router | goldentest | SQLite |
| infra/gateway | 都度判断 | emailtrap等 |

### モック生成（go:generateを各IFファイルに記載）
```bash
go generate ./...
```

### 単体テスト例
```go
func TestCreateUserUsecase_Exec(t *testing.T) {
    ctrl := gomock.NewController(t)
    mockRepo := mockuserdm.NewMockUserRepository(ctrl)
    mockDS := mockuserdm.NewMockIsExistUserDomainService(ctrl)
    mockDS.EXPECT().Exec(gomock.Any(), gomock.Any()).Return(false, nil)
    mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)

    uc := createuserusecase.NewCreateUserUsecase(mockRepo, mockDS)
    out, err := uc.Exec(context.Background(), createuserusecase.CreateUserInput{Name: "John", Email: "john@example.com"})
    require.NoError(t, err)  // errorは必ずチェック
    assert.NotEmpty(t, out.UserID)
}
```

### RDBテスト（SQLite）
```go
func setupTestDB(t *testing.T) *generated.Queries {
    db, err := sql.Open("sqlite3", ":memory:")
    require.NoError(t, err)  // errorは必ずチェック
    _, err = db.Exec(`CREATE TABLE users (id TEXT PRIMARY KEY, name TEXT, email TEXT UNIQUE, created_at DATETIME)`)
    require.NoError(t, err)
    t.Cleanup(func() { db.Close() })
    return generated.New(db)
}
```

### goldentest更新
```bash
go test ./internal/infra/router/... -update
```

## golangci-lint

```yaml
# .golangci.yml
linters:
  enable:
    - errcheck, gosimple, govet, staticcheck, unused
    - goimports, gofmt, misspell, unconvert, unparam
    - gosec, bodyclose, sqlclosecheck, contextcheck, nilerr, errorlint
linters-settings:
  goimports:
    local-prefixes: github.com/your-org/your-project
  govet:
    enable-all: true
issues:
  exclude-rules:
    - path: _test\.go
      linters: [errcheck, gosec]
    - path: db/sqlc/generated/
      linters: [all]
```

## Makefile

```makefile
migrate-dry: psqldef -U $(DB_USER) -h $(DB_HOST) $(DB_NAME) --dry-run < db/migrations/schema.sql
migrate:     psqldef -U $(DB_USER) -h $(DB_HOST) $(DB_NAME) < db/migrations/schema.sql
sqlc:        sqlc generate -f internal/infra/rdb/sqlc.yaml
wire:        cd internal/di && wire
mock:        go generate ./...
lint:        golangci-lint run ./...
test:        go test ./... -v
generate:    sqlc wire mock
run:         go run ./cmd/api/main.go
```

## コード実装の原則

- **必要最小限のコードのみ実装**（使わないCRUDメソッドを全部作らない、今必要なものだけ）
- 使用しないメソッド・フィールド・構造体は作成しない
- 将来使うかもしれないコードは書かない（YAGNI）

## エラーハンドリング

### パッケージ構成

```
internal/apperr/
├── error_code.go   # エラーコード定義 (Code型)
└── app_error.go    # AppError構造体とコンストラクタ
```

### HTTPステータスの使い分け

| ステータス | 用途 | コンストラクタ |
|-----------|------|---------------|
| 400 Bad Request | リクエスト構文エラー（JSON不正など） | `BadRequest`, `BadRequestWrap` |
| 401 Unauthorized | 認証エラー | `Unauthorized`, `UnauthorizedWrap` |
| 403 Forbidden | 認可エラー | `Forbidden`, `ForbiddenWrap` |
| 404 Not Found | リソース未検出 | `NotFound`, `NotFoundWrap` |
| 409 Conflict | 重複エラー | `Conflict`, `ConflictWrap` |
| 422 Unprocessable Entity | ドメインバリデーションエラー | `UnprocessableEntity`, `UnprocessableEntityWrap` |
| 500 Internal Server Error | DB/インフラエラー | `Internal` |

### コンストラクタの使い分け

```go
// 元のエラーがない場合
apperr.Conflict(apperr.CodeKeywordAlreadyExists)

// 元のエラーをラップする場合（スタックトレース保持）
apperr.UnprocessableEntityWrap(err, apperr.CodeInvalidRequest)

// 500系は常にラップ
apperr.Internal(err)
```

### Controller での使用

```go
// JSONバインドエラー → 400
if err := ctx.ShouldBindJSON(&req); err != nil {
    _ = ctx.Error(apperr.BadRequestWrap(err, apperr.CodeInvalidRequest))
    return
}

// 認証チェック → 401
if userID == "" {
    _ = ctx.Error(apperr.Unauthorized(apperr.CodeUnauthorized))
    return
}

// リソース検索 → 404
user, err := c.userQueryService.FindByID(ctx, userID)
if err != nil {
    _ = ctx.Error(apperr.NotFoundWrap(err, apperr.CodeUserNotFound))
    return
}

// Usecase呼び出し → Usecaseが返すAppErrorをそのまま渡す
output, err := c.createRuleUsecase.Exec(ctx, input)
if err != nil {
    _ = ctx.Error(err)
    return
}
```

### Usecase での使用

```go
// ドメインバリデーション → 422
keyword, err := ruledm.NewKeyword(input.Keyword)
if err != nil {
    return nil, apperr.UnprocessableEntityWrap(err, apperr.CodeInvalidRequest)
}

// 重複チェック → 409
if exists {
    return nil, apperr.Conflict(apperr.CodeKeywordAlreadyExists)
}

// DB操作 → 500
if err := u.ruleRepo.Save(ctx, rule); err != nil {
    return nil, apperr.Internal(err)
}
```

### エラーコード追加手順

1. `internal/apperr/error_code.go` に追加:
```go
const (
    CodeNewError Code = "NEW_ERROR"
)
```

2. フロントエンド `frontend/src/lib/errorCodes.ts` に対応メッセージ追加:
```typescript
export const ERROR_MESSAGES: Record<string, string> = {
    NEW_ERROR: '新しいエラーメッセージ',
};
```

## 禁止事項

- domain層からinfra層への依存
- usecase層からinterface層への依存
- 不要なセッター追加（必要なセッターはOK、ゲッターは名詞のみ）
- ディレクトリ名に `command/` `query/` を使用
- `user_id_vo.go` を shared/ に置く（ドメイン固有VOは各xxxdmに）
- any型の使用
- 手動DI（必ずwireを使用）
- 生成コード（sqlc/wire_gen/mock）の手動編集
- テストなしでのマージ
- golangci-lintエラーを無視してコミット
- エラーをラップせずに新規エラーを作成（スタックトレースが失われる）
