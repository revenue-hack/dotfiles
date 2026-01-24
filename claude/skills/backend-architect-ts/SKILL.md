---
name: backend-architect-ts
description: TypeScript + クリーンアーキテクチャ + CQRS バックエンド設計ガイドライン。新規機能追加、リファクタリング、コード生成、アーキテクチャ相談時に使用。
---

# TypeScript Backend Architecture Guidelines

## 技術スタック

TypeScript (strict) / Clean Architecture / CQRS / tsyringe (DI) / Prisma / ESLint / Jest

## ディレクトリ構造

```
.
├── prisma/
│   ├── schema.prisma                # スキーマ定義
│   └── migrations/                  # 自動生成
├── src/
│   ├── domain/
│   │   ├── shared/                  # 汎用VO（世間一般で共通のもののみ）
│   │   │   ├── email.vo.ts
│   │   │   └── createdAt.vo.ts
│   │   └── user/                    # 集約
│   │       ├── user.entity.ts       # コンストラクタはprivate
│   │       ├── user.factory.ts      # public生成メソッド
│   │       ├── user.repository.ts   # IF定義
│   │       ├── userId.vo.ts         # ※user固有、sharedに置かない
│   │       ├── userName.vo.ts
│   │       └── isExistUser.domainService.ts
│   ├── usecase/
│   │   ├── createUser/              # 必ず動詞から始まる
│   │   │   ├── createUser.input.ts
│   │   │   ├── createUser.output.ts
│   │   │   └── createUser.usecase.ts
│   │   └── userQueryService.ts      # Query系IF
│   ├── interface/
│   │   ├── controller/userController.ts
│   │   └── presentation/userPresenter.ts
│   ├── infra/
│   │   ├── router/router.ts         # FWは技術的詳細なのでinfra
│   │   ├── rdb/
│   │   │   ├── repoImpl/userRepository.ts      # Repository実装
│   │   │   └── queryImpl/userQueryService.ts   # QueryService実装
│   │   ├── gateway/                 # 外部SaaS（SendGrid, SQS等）
│   │   └── prisma/client.ts
│   ├── di/container.ts
│   └── main.ts
└── tests/                           # src構造をミラー
    ├── domain/user/
    │   ├── user.entity.test.ts
    │   └── user.factory.test.ts
    ├── usecase/createUser/createUser.usecase.test.ts
    ├── infra/
    │   ├── router/router.test.ts
    │   └── rdb/repoImpl/userRepository.test.ts
    └── helpers/testDb.ts
```

## ファイル命名規則（camelCase）

| 種類 | ファイル名 | 例 |
|-----|----------|-----|
| エンティティ | `xxx.entity.ts` | `user.entity.ts` |
| ファクトリー | `xxx.factory.ts` | `user.factory.ts` |
| リポジトリIF | `xxx.repository.ts` | `user.repository.ts` |
| Value Object | `xxxYyy.vo.ts` | `userId.vo.ts` |
| ドメインサービス | `xxxYyy.domainService.ts` | `isExistUser.domainService.ts` |

## CQRS

| ユースケース | パターン | 実装場所 |
|-------------|---------|---------|
| 一覧・検索（複雑） | QueryService | usecase/にIF → infra/rdb/queryImpl/に実装 |
| 詳細取得（単純） | Repository | Repositoryで完結する場合はRepositoryでOK |
| 作成・更新・削除 | Usecase + Domain + Repository | ドメインロジックを経由 |

## Domain層

### Value Object
```typescript
// domain/user/userId.vo.ts
export class UserId {
  private constructor(public readonly value: string) {}
  static create(): UserId { return new UserId(crypto.randomUUID()); }
  static reconstruct(value: string): UserId { return new UserId(value); }
  equals(other: UserId): boolean { return this.value === other.value; }
}
```

### エンティティ（コンストラクタはprivate、ファクトリーからのみ生成）
```typescript
// domain/user/user.entity.ts
interface UserProps { id: UserId; name: UserName; email: Email; createdAt: CreatedAt; }

export class User {
  private constructor(private readonly props: UserProps) {}  // private
  static _create(props: UserProps): User { return new User(props); }  // ファクトリー専用
  get id(): UserId { return this.props.id; }  // ゲッターは名詞のみ、セッターは必要時のみ
  get name(): UserName { return this.props.name; }
}
```

### ファクトリー（ユースケース別にpublic生成メソッド）
```typescript
// domain/user/user.factory.ts
export const genUserForCreate = (name: string, email: string): User => {
  const userName = UserName.create(name);  // バリデーションあり
  const emailVO = Email.create(email);
  return User._create({
    id: UserId.create(),
    name: userName,
    email: emailVO,
    createdAt: CreatedAt.create(),
  });
};

export const genUserForReconstruct = (
  id: string, name: string, email: string, createdAt: Date
): User => {
  return User._create({
    id: UserId.reconstruct(id),
    name: UserName.reconstruct(name),
    email: Email.reconstruct(email),
    createdAt: CreatedAt.reconstruct(createdAt),
  });
};
```

### リポジトリIF（取得系はドメインかnullを返す）
```typescript
// domain/user/user.repository.ts
export interface UserRepository {
  findById(id: UserId): Promise<User | null>;
  findByEmail(email: Email): Promise<User | null>;
  existsById(id: UserId): Promise<boolean>;  // プリミティブ返却もOK
  save(user: User): Promise<void>;
  delete(id: UserId): Promise<void>;
}
```

### ドメインサービス（DBを使うドメインロジック）
```typescript
// domain/user/isExistUser.domainService.ts
@injectable()
export class IsExistUserDomainService {
  constructor(@inject('UserRepository') private userRepo: UserRepository) {}
  async exec(email: Email): Promise<boolean> {  // Execメソッド1つのみ
    const user = await this.userRepo.findByEmail(email);
    return user !== null;
  }
}
// 命名例: IsExistXxx, CanXxx
```

### shared/の配置ルール
- **OK**: `email.vo.ts`, `createdAt.vo.ts`（世間一般で汎用的）
- **NG**: `userId.vo.ts`（userの所有物 → user/に置く。他ドメインからは`import { UserId } from '@/domain/user/userId.vo'`で参照）

## Usecase層

```typescript
// usecase/createUser/createUser.usecase.ts
@injectable()
export class CreateUserUsecase {
  constructor(
    @inject('UserRepository') private userRepo: UserRepository,
    @inject('IsExistUserDomainService') private isExistUserDS: IsExistUserDomainService,
  ) {}

  async exec(input: CreateUserInput): Promise<CreateUserOutput> {
    const exists = await this.isExistUserDS.exec(Email.create(input.email));
    if (exists) throw new Error('User already exists');

    const user = genUserForCreate(input.name, input.email);  // ファクトリー使用
    await this.userRepo.save(user);
    return { userId: user.id.value };
  }
}
```

## Infra層

### Prisma設定
```prisma
// prisma/schema.prisma
model User {
  id        String   @id
  name      String
  email     String   @unique
  createdAt DateTime @default(now()) @map("created_at")
  updatedAt DateTime @updatedAt @map("updated_at")
  @@map("users")
}
```

### Prismaコマンド
```bash
npx prisma migrate dev --name init  # マイグレーション作成・適用
npx prisma generate                 # クライアント生成
npx prisma db push                  # 開発用スキーマ反映
```

### リポジトリ実装
```typescript
// infra/rdb/repoImpl/userRepository.ts
@injectable()
export class UserRepositoryImpl implements UserRepository {
  constructor(@inject('PrismaClient') private prisma: PrismaClient) {}

  async findById(id: UserId): Promise<User | null> {
    const row = await this.prisma.user.findUnique({ where: { id: id.value } });
    if (!row) return null;
    return genUserForReconstruct(row.id, row.name, row.email, row.createdAt);  // ファクトリー使用
  }

  async save(user: User): Promise<void> {
    await this.prisma.user.upsert({
      where: { id: user.id.value },
      create: { id: user.id.value, name: user.name.value, email: user.email.value },
      update: { name: user.name.value, email: user.email.value },
    });
  }
}
```

## DI（tsyringe）

```typescript
// di/container.ts
import 'reflect-metadata';
import { container } from 'tsyringe';

container.register('PrismaClient', { useValue: prisma });
container.register('UserRepository', { useClass: UserRepositoryImpl });
container.register('IsExistUserDomainService', { useClass: IsExistUserDomainService });
container.register('UserQueryService', { useClass: UserQueryServiceImpl });
container.register('CreateUserUsecase', { useClass: CreateUserUsecase });

export { container };
```

## テスト（必須）※テストなしはマージ禁止

| 層 | 方法 | DB |
|----|------|-----|
| domain/usecase | Jest mock | 不要 |
| infra/rdb | SQLite直接アクセス | SQLite |
| infra/router | Jest + supertest | SQLite |
| infra/gateway | 都度判断（emailtrap等） | - |

### Jest設定
```javascript
// jest.config.js
module.exports = {
  preset: 'ts-jest',
  testEnvironment: 'node',
  roots: ['<rootDir>/tests'],
  testMatch: ['**/*.test.ts'],
  moduleNameMapper: { '^@/(.*)$': '<rootDir>/src/$1' },
  setupFilesAfterEnv: ['<rootDir>/tests/helpers/setup.ts'],
};
```

### テスト例（Jest mock）
```typescript
// tests/usecase/createUser/createUser.usecase.test.ts
describe('CreateUserUsecase', () => {
  it('should create a user', async () => {
    const mockRepo: jest.Mocked<UserRepository> = {
      findById: jest.fn(), findByEmail: jest.fn(), existsById: jest.fn(),
      save: jest.fn().mockResolvedValue(undefined), delete: jest.fn(),
    };
    const mockDS = { exec: jest.fn().mockResolvedValue(false) } as unknown as IsExistUserDomainService;

    const usecase = new CreateUserUsecase(mockRepo, mockDS);
    const result = await usecase.exec({ name: 'John', email: 'john@example.com' });

    expect(result.userId).toBeDefined();
    expect(mockRepo.save).toHaveBeenCalledTimes(1);
  });
});
```

### RDBテスト（SQLite）
```typescript
// tests/helpers/testDb.ts
export const createTestPrismaClient = (): PrismaClient => {
  return new PrismaClient({ datasources: { db: { url: 'file:./test.db' } } });
};

// tests/infra/rdb/repoImpl/userRepository.test.ts
const prisma = createTestPrismaClient();
const repo = new UserRepositoryImpl(prisma);

const user = genUserForCreate('John', 'john@example.com');
await repo.save(user);
const found = await repo.findById(user.id);
expect(found?.name.value).toBe('John');
```

### Routerテスト（supertest）
```typescript
const response = await request(app).get('/api/v1/users/1').expect(200);
expect(response.body).toHaveProperty('id');
```

## ESLint

```json
// .eslintrc.json
{
  "extends": ["plugin:@typescript-eslint/recommended"],
  "rules": {
    "@typescript-eslint/no-explicit-any": "error",
    "@typescript-eslint/explicit-function-return-type": "warn",
    "@typescript-eslint/strict-boolean-expressions": "error",
    "@typescript-eslint/no-floating-promises": "error",
    "@typescript-eslint/no-unused-vars": ["error", { "argsIgnorePattern": "^_" }],
    "import/order": ["error", { "alphabetize": { "order": "asc" } }]
  }
}
```

## tsconfig.json

```json
{
  "compilerOptions": {
    "target": "ES2022",
    "module": "NodeNext",
    "moduleResolution": "NodeNext",
    "strict": true,
    "esModuleInterop": true,
    "experimentalDecorators": true,
    "emitDecoratorMetadata": true,
    "baseUrl": ".",
    "paths": { "@/*": ["./src/*"] }
  }
}
```

## package.json scripts

```json
{
  "scripts": {
    "dev": "tsx watch src/main.ts",
    "build": "tsc",
    "lint": "eslint . --ext .ts",
    "lint:fix": "eslint . --ext .ts --fix",
    "test": "jest",
    "test:coverage": "jest --coverage",
    "prisma:migrate": "prisma migrate dev",
    "prisma:generate": "prisma generate"
  }
}
```

## パスalias

`@/*` → `./src/*`

## 禁止事項

- domain層からinfra層への依存
- usecase層からinterface層への依存
- 不要なセッター追加（必要なセッターはOK、ゲッターは名詞のみ）
- ディレクトリ名に `command/` `query/` を使用
- `userId.vo.ts` を shared/ に置く（ドメイン固有VOは各ドメインディレクトリに）
- `any` 型の使用（`unknown`か適切な型を使用）
- ESLintエラーを無視してコミット
- Prisma生成コードの手動編集
- テストなしでのマージ
