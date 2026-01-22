---
name: front-architect
description: Next.js + TypeScript フロントエンド設計ガイドライン。Atomic Design（atoms/organismsのみ、moleculesなし）、styled-components + Tailwind、ESLint、Playwright。Figmaからのコード生成後のリファクタリング、コンポーネント作成、プロジェクトセットアップ時に使用。
---

# Next.js Frontend Guidelines

## 技術スタック

Next.js (App Router) / TypeScript (strict) / styled-components + Tailwind / ESLint / Playwright

## ディレクトリ構造

```
src/
├── app/                      # App Router (default export必須)
├── components/
│   ├── atoms/                # 最小単位: Button, Input, Label, Icon
│   │   └── Button/
│   │       ├── index.tsx
│   │       └── Button.styles.ts
│   └── organisms/            # atomsの組み合わせ: Header, Footer, UserCard
│       └── Header/
│           ├── index.tsx
│           ├── Header.styles.ts
│           └── useHeader.ts  # organism専用Hook
├── hooks/                    # 再利用Hooksのみ
├── lib/                      # helpers.ts, api.ts に集約（utilフォルダ禁止）
├── types/
└── tests/e2e/                # Playwright
```

## Atomic Design

- **atoms**: 単一責務の最小UI部品
- **organisms**: atomsを組み合わせた機能単位
- **molecules**: 作成しない（分類が曖昧になるため）

## コンポーネント規約

各コンポーネントは**ディレクトリで管理**：`ComponentName/index.tsx` + `ComponentName.styles.ts`

| 項目 | ルール |
|------|--------|
| export | named export（page.tsx/layout.tsxのみdefault） |
| Props型 | `interface ComponentNameProps` |
| ファイル名 | PascalCase |
| 関数 | アロー関数 |
| any | **禁止**（unknownか適切な型を使用） |
| styled-components props | `$` prefix（transient props） |

## Hooks配置

- organism専用 → コンポーネントディレクトリ内（`useHeader.ts`）
- 複数箇所で再利用 → `src/hooks/`

## ESLint

`@typescript-eslint/no-explicit-any: "error"` を必須設定。
一括修正: `npx eslint . --ext .ts,.tsx --fix`

## E2Eテスト

Playwright使用。テストは `src/tests/e2e/` に配置。

## パスalias

`@/*` → `./src/*`
