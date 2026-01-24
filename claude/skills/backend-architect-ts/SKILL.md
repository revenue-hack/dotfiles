---
name: backend-architect-ts
description: TypeScript + クリーンアーキテクチャ + CQRS バックエンド設計ガイドライン。新規機能追加、リファクタリング、コード生成、アーキテクチャ相談時に使用。
---

# TypeScript Backend Architecture Guidelines

## 技術スタック

TypeScript (strict) / Clean Architecture / CQRS / DI (tsyringe) / Prisma / ESLint

## ディレクトリ構造

```
.
├── prisma/
│   ├── schema.prisma            # Prismaスキーマ定義
│   └── migrations/              # マイグレーション（自動生成）
│
└── src/
    ├── domain/                  # ドメイン層（ビジネスロジックの中核）
    │   ├── shared/              # 汎用Value Object（世間一般で共通のもののみ）
    │   │   ├── createdAt.vo.ts
    │   │   └── email.vo.ts
    │   └── user/                # ドメインモデル（集約名）
    │       ├── user.entity.ts
    │       ├── user.repository.ts       # リポジトリインターフェース
    │       ├── userId.vo.ts
    │       ├── userName.vo.ts
    │       └── isExistUser.domainService.ts
    │
    ├── usecase/                 # ユースケース層（動詞から始まる）
    │   ├── createUser/
    │   │   ├── createUser.input.ts
    │   │   ├── createUser.output.ts
    │   │   └── createUser.usecase.ts
    │   ├── updateUser/
    │   │   └── ...
    │   └── userQueryService.ts  # Query系インターフェース
    │
    ├── interface/               # インターフェース層
    │   ├── controller/
    │   │   └── userController.ts
    │   └── presentation/
    │       └── userPresenter.ts
    │
    ├── infra/                   # インフラ層（自由に構成可能）
    │   ├── router/              # ルーター（FWは技術的詳細）
    │   │   └── router.ts
    │   ├── rdb/
    │   │   ├── repoImpl/        # リポジトリ実装（Prisma使用）
    │   │   │   └── userRepository.ts
    │   │   └── queryImpl/       # QueryService実装
    │   │       └── userQueryService.ts
    │   ├── gateway/             # 外部SaaS連携
    │   │   └── sendGridGateway.ts
    │   └── prisma/
    │       └── client.ts        # Prismaクライアント
    │
    ├── di/
    │   └── container.ts         # DIコンテナ設定
    │
    └── main.ts                  # エントリーポイント
```

## ファイル命名規則

| 種類 | ファイル名 | 例 |
|-----|----------|-----|
| エンティティ | `xxx.entity.ts` | `user.entity.ts` |
| リポジトリIF | `xxx.repository.ts` | `user.repository.ts` |
| Value Object | `xxxYyy.vo.ts` | `userId.vo.ts`, `userName.vo.ts` |
| ドメインサービス | `xxxYyy.domainService.ts` | `isExistUser.domainService.ts` |
| ユースケース | `xxxYyy.usecase.ts` | `createUser.usecase.ts` |
| Input/Output | `xxxYyy.input.ts` / `xxxYyy.output.ts` | `createUser.input.ts` |
| コントローラ | `xxxController.ts` | `userController.ts` |
| 実装 | `xxxYyy.ts` (implディレクトリ内) | `userRepository.ts` |

**すべてキャメルケース（camelCase）を使用**

## CQRS パターン

### Query（参照系）

取得のみのユースケースは **QueryService** パターンを使用。

```typescript
// usecase/userQueryService.ts - インターフェース定義
export interface UserReadModel {
  id: string;
  name: string;
  email: string;
  createdAt: Date;
}

export interface UserQueryService {
  findById(id: string): Promise<UserReadModel | null>;
  list(params: ListParams): Promise<UserReadModel[]>;
}

// infra/rdb/queryImpl/userQueryService.ts - 実装
import { injectable, inject } from 'tsyringe';
import { PrismaClient } from '@prisma/client';

@injectable()
export class UserQueryServiceImpl implements UserQueryService {
  constructor(
    @inject('PrismaClient') private prisma: PrismaClient
  ) {}

  async findById(id: string): Promise<UserReadModel | null> {
    const user = await this.prisma.user.findUnique({ where: { id } });
    if (!user) return null;
    return {
      id: user.id,
      name: user.name,
      email: user.email,
      createdAt: user.createdAt,
    };
  }
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

- **必ず動詞から始まる**: `createUser`, `updateUser`, `deleteUser`
- ディレクトリで管理し、Input/Output/Usecaseを分離

### 構成ファイル

```typescript
// usecase/createUser/createUser.input.ts
export interface CreateUserInput {
  name: string;
  email: string;
}

// usecase/createUser/createUser.output.ts
export interface CreateUserOutput {
  userId: string;
}

// usecase/createUser/createUser.usecase.ts
import { injectable, inject } from 'tsyringe';
import { UserRepository } from '@/domain/user/user.repository';
import { IsExistUserDomainService } from '@/domain/user/isExistUser.domainService';
import { User } from '@/domain/user/user.entity';
import { CreateUserInput } from './createUser.input';
import { CreateUserOutput } from './createUser.output';

@injectable()
export class CreateUserUsecase {
  constructor(
    @inject('UserRepository') private userRepo: UserRepository,
    @inject('IsExistUserDomainService') private isExistUserDS: IsExistUserDomainService
  ) {}

  async exec(input: CreateUserInput): Promise<CreateUserOutput> {
    // ドメインサービスで存在チェック
    const exists = await this.isExistUserDS.exec(input.email);
    if (exists) {
      throw new Error('User already exists');
    }

    // ドメインでエンティティ生成（ドメインルールはここに集約）
    const user = User.create({
      name: input.name,
      email: input.email,
    });

    // 永続化
    await this.userRepo.save(user);

    return { userId: user.id.value };
  }
}
```

## Domain層の規約

### Value Object

```typescript
// domain/user/userId.vo.ts
export class UserId {
  private constructor(public readonly value: string) {}

  static create(value?: string): UserId {
    return new UserId(value ?? crypto.randomUUID());
  }

  static reconstruct(value: string): UserId {
    return new UserId(value);
  }

  equals(other: UserId): boolean {
    return this.value === other.value;
  }
}

// domain/user/userName.vo.ts
export class UserName {
  private constructor(public readonly value: string) {}

  static create(value: string): UserName {
    if (value.length < 1 || value.length > 100) {
      throw new Error('UserName must be between 1 and 100 characters');
    }
    return new UserName(value);
  }

  static reconstruct(value: string): UserName {
    return new UserName(value);
  }
}
```

### 共通Value Object（shared/）

**世間一般で共通のもののみ** を配置。

```typescript
// domain/shared/email.vo.ts - OK（汎用的）
// domain/shared/createdAt.vo.ts - OK（汎用的）
// domain/shared/userId.vo.ts - NG！（userの所有物なので user/ に置く）
```

他ドメインで `UserId` を使う場合でも、それは `import { UserId } from '@/domain/user/userId.vo'` として参照する。

### エンティティ

```typescript
// domain/user/user.entity.ts
import { UserId } from './userId.vo';
import { UserName } from './userName.vo';
import { Email } from '@/domain/shared/email.vo';
import { CreatedAt } from '@/domain/shared/createdAt.vo';

interface UserProps {
  id: UserId;
  name: UserName;
  email: Email;
  createdAt: CreatedAt;
}

interface CreateUserProps {
  name: string;
  email: string;
}

export class User {
  private constructor(private props: UserProps) {}

  // 新規作成
  static create(input: CreateUserProps): User {
    return new User({
      id: UserId.create(),
      name: UserName.create(input.name),
      email: Email.create(input.email),
      createdAt: CreatedAt.create(),
    });
  }

  // 再構築（DBからの復元）
  static reconstruct(props: UserProps): User {
    return new User(props);
  }

  // ゲッターのみ公開（不変性を保つ）
  get id(): UserId { return this.props.id; }
  get name(): UserName { return this.props.name; }
  get email(): Email { return this.props.email; }
  get createdAt(): CreatedAt { return this.props.createdAt; }

  // ドメインロジック
  changeName(newName: string): void {
    this.props.name = UserName.create(newName);
  }
}
```

### リポジトリインターフェース

取得系メソッドは **ドメインか null** を返す。

```typescript
// domain/user/user.repository.ts
import { User } from './user.entity';
import { UserId } from './userId.vo';
import { Email } from '@/domain/shared/email.vo';

export interface UserRepository {
  // 取得系: ドメインを返す
  findById(id: UserId): Promise<User | null>;
  findByEmail(email: Email): Promise<User | null>;

  // 取得系: プリミティブを返す（存在確認など）
  existsById(id: UserId): Promise<boolean>;

  // 永続化系
  save(user: User): Promise<void>;
  delete(id: UserId): Promise<void>;
}
```

### ドメインサービス

**DBを使うドメインロジック** はドメインサービスに切り出す。

```typescript
// domain/user/isExistUser.domainService.ts
import { injectable, inject } from 'tsyringe';
import { UserRepository } from './user.repository';
import { Email } from '@/domain/shared/email.vo';

@injectable()
export class IsExistUserDomainService {
  constructor(
    @inject('UserRepository') private userRepo: UserRepository
  ) {}

  // メソッドは1つだけ（exec）
  async exec(email: string): Promise<boolean> {
    const emailVO = Email.create(email);
    const user = await this.userRepo.findByEmail(emailVO);
    return user !== null;
  }
}
```

**ドメインサービスの特徴**:
- Repositoryインターフェースを経由してデータ取得
- `exec` メソッド1つのみ
- 命名: `isExistXxx.domainService.ts`, `canXxx.domainService.ts` など

## Infra層の規約

### Prisma クライアント

```typescript
// infra/prisma/client.ts
import { PrismaClient } from '@prisma/client';

export const prisma = new PrismaClient();
```

### リポジトリ実装（Prisma使用）

```typescript
// infra/rdb/repoImpl/userRepository.ts
import { injectable, inject } from 'tsyringe';
import { PrismaClient } from '@prisma/client';
import { UserRepository } from '@/domain/user/user.repository';
import { User } from '@/domain/user/user.entity';
import { UserId } from '@/domain/user/userId.vo';
import { UserName } from '@/domain/user/userName.vo';
import { Email } from '@/domain/shared/email.vo';
import { CreatedAt } from '@/domain/shared/createdAt.vo';

@injectable()
export class UserRepositoryImpl implements UserRepository {
  constructor(
    @inject('PrismaClient') private prisma: PrismaClient
  ) {}

  async findById(id: UserId): Promise<User | null> {
    const row = await this.prisma.user.findUnique({
      where: { id: id.value }
    });
    if (!row) return null;
    return this.toEntity(row);
  }

  async save(user: User): Promise<void> {
    await this.prisma.user.upsert({
      where: { id: user.id.value },
      create: {
        id: user.id.value,
        name: user.name.value,
        email: user.email.value,
        createdAt: user.createdAt.value,
      },
      update: {
        name: user.name.value,
        email: user.email.value,
      },
    });
  }

  // toEntity: Prismaモデル → ドメインエンティティ変換
  private toEntity(row: { id: string; name: string; email: string; createdAt: Date }): User {
    return User.reconstruct({
      id: UserId.reconstruct(row.id),
      name: UserName.reconstruct(row.name),
      email: Email.reconstruct(row.email),
      createdAt: CreatedAt.reconstruct(row.createdAt),
    });
  }
}
```

### gateway（外部SaaS連携）

```typescript
// infra/gateway/sendGridGateway.ts
import sgMail from '@sendgrid/mail';

export class SendGridGateway {
  constructor(apiKey: string) {
    sgMail.setApiKey(apiKey);
  }

  async sendEmail(to: string, subject: string, body: string): Promise<void> {
    await sgMail.send({
      to,
      from: 'noreply@example.com',
      subject,
      text: body,
    });
  }
}
```

## DI（依存性注入）- tsyringe

```typescript
// di/container.ts
import 'reflect-metadata';
import { container } from 'tsyringe';
import { PrismaClient } from '@prisma/client';
import { prisma } from '@/infra/prisma/client';

// Prisma
container.register('PrismaClient', { useValue: prisma });

// Repository
container.register('UserRepository', { useClass: UserRepositoryImpl });

// Domain Service
container.register('IsExistUserDomainService', { useClass: IsExistUserDomainService });

// Query Service
container.register('UserQueryService', { useClass: UserQueryServiceImpl });

// Usecase
container.register('CreateUserUsecase', { useClass: CreateUserUsecase });

export { container };
```

### main.ts

```typescript
// main.ts
import 'reflect-metadata';
import '@/di/container';
import { container } from 'tsyringe';
import express from 'express';
import { UserController } from '@/interface/controller/userController';

const app = express();
app.use(express.json());

const userController = container.resolve(UserController);

app.get('/api/v1/users/:id', (req, res) => userController.get(req, res));
app.post('/api/v1/users', (req, res) => userController.create(req, res));

app.listen(3000, () => {
  console.log('Server running on port 3000');
});
```

## Controller例

```typescript
// interface/controller/userController.ts
import { injectable, inject } from 'tsyringe';
import { Request, Response } from 'express';
import { CreateUserUsecase } from '@/usecase/createUser/createUser.usecase';
import { UserQueryService } from '@/usecase/userQueryService';

@injectable()
export class UserController {
  constructor(
    @inject('CreateUserUsecase') private createUserUsecase: CreateUserUsecase,
    @inject('UserQueryService') private userQueryService: UserQueryService
  ) {}

  async get(req: Request, res: Response): Promise<void> {
    const { id } = req.params;
    const user = await this.userQueryService.findById(id);
    if (!user) {
      res.status(404).json({ error: 'User not found' });
      return;
    }
    res.json(user);
  }

  async create(req: Request, res: Response): Promise<void> {
    try {
      const output = await this.createUserUsecase.exec(req.body);
      res.status(201).json({ userId: output.userId });
    } catch (error) {
      res.status(400).json({ error: (error as Error).message });
    }
  }
}
```

## Prisma 設定

### schema.prisma

```prisma
// prisma/schema.prisma
generator client {
  provider = "prisma-client-js"
}

datasource db {
  provider = "postgresql"  // または "mysql"
  url      = env("DATABASE_URL")
}

model User {
  id        String   @id
  name      String
  email     String   @unique
  createdAt DateTime @default(now()) @map("created_at")
  updatedAt DateTime @updatedAt @map("updated_at")
  posts     Post[]

  @@map("users")
}

model Post {
  id        String   @id
  userId    String   @map("user_id")
  title     String
  body      String?
  createdAt DateTime @default(now()) @map("created_at")
  updatedAt DateTime @updatedAt @map("updated_at")
  user      User     @relation(fields: [userId], references: [id])

  @@index([userId])
  @@map("posts")
}
```

### Prisma コマンド

```bash
# マイグレーション作成・適用
npx prisma migrate dev --name init

# 本番適用
npx prisma migrate deploy

# クライアント生成
npx prisma generate

# スキーマをDBに反映（開発用）
npx prisma db push
```

## ESLint 設定

```json
// .eslintrc.json
{
  "root": true,
  "parser": "@typescript-eslint/parser",
  "parserOptions": {
    "project": "./tsconfig.json"
  },
  "plugins": ["@typescript-eslint", "import"],
  "extends": [
    "eslint:recommended",
    "plugin:@typescript-eslint/recommended",
    "plugin:@typescript-eslint/recommended-requiring-type-checking",
    "plugin:import/recommended",
    "plugin:import/typescript"
  ],
  "rules": {
    "@typescript-eslint/no-explicit-any": "error",
    "@typescript-eslint/explicit-function-return-type": "warn",
    "@typescript-eslint/no-unused-vars": ["error", { "argsIgnorePattern": "^_" }],
    "@typescript-eslint/strict-boolean-expressions": "error",
    "@typescript-eslint/no-floating-promises": "error",
    "import/order": [
      "error",
      {
        "groups": ["builtin", "external", "internal", "parent", "sibling", "index"],
        "newlines-between": "never",
        "alphabetize": { "order": "asc" }
      }
    ],
    "import/no-duplicates": "error"
  },
  "settings": {
    "import/resolver": {
      "typescript": {
        "project": "./tsconfig.json"
      }
    }
  }
}
```

### tsconfig.json

```json
{
  "compilerOptions": {
    "target": "ES2022",
    "module": "NodeNext",
    "moduleResolution": "NodeNext",
    "lib": ["ES2022"],
    "outDir": "./dist",
    "rootDir": "./src",
    "strict": true,
    "esModuleInterop": true,
    "skipLibCheck": true,
    "forceConsistentCasingInFileNames": true,
    "resolveJsonModule": true,
    "declaration": true,
    "declarationMap": true,
    "sourceMap": true,
    "experimentalDecorators": true,
    "emitDecoratorMetadata": true,
    "baseUrl": ".",
    "paths": {
      "@/*": ["./src/*"]
    }
  },
  "include": ["src/**/*"],
  "exclude": ["node_modules", "dist"]
}
```

## package.json scripts

```json
{
  "scripts": {
    "dev": "tsx watch src/main.ts",
    "build": "tsc",
    "start": "node dist/main.js",
    "lint": "eslint . --ext .ts",
    "lint:fix": "eslint . --ext .ts --fix",
    "prisma:generate": "prisma generate",
    "prisma:migrate": "prisma migrate dev",
    "prisma:push": "prisma db push"
  }
}
```

## パスalias

`@/*` → `./src/*`

## 禁止事項

- domain層から infra層への依存
- usecase層から interface層への依存
- エンティティへの直接的なプロパティ変更（セッター経由のみ）
- ディレクトリ名に `command/` `query/` を使用
- `userId.vo.ts` を shared/ に置く（ドメイン固有のVOは各ドメインディレクトリに）
- `any` 型の使用（`unknown` か適切な型を使用）
- ESLint エラーを無視してコミット
- Prisma生成コードの手動編集
