# golang-trainning-frontend

エンドユーザ向けの参照専用 API サーバ。CQRS の Query 側として機能する。

---

## アーキテクチャ概要

Clean Architecture + CQRS (Query Side) を採用。

```
HTTPリクエスト
     ↓
[ Router ]
     ↓
[ Controller ]        adapter層
     ↓
[ Usecase (Inputport / Interactor) ]   usecase層
     ↓
[ Repository (Outputport) ]            usecase層（インターフェース定義）
     ↓
[ Repository 実装 ]                    adapter層
     ↓
[ DB (MySQL) ]
```

---

## CQRS における位置づけ

このシステムは**読み取り専用 (Query Side)** に特化している。

| | Command Side | Query Side（本システム）|
|---|---|---|
| 役割 | データの書き込み・更新・削除 | データの参照・返却 |
| ドメインロジック | あり | なし |
| モデル | Domain Entity | Query Model |

書き込みは別サービス（admin 等）が同一 DB に対して行い、本サービスはその DB を参照するだけのため、サービス間通信は発生しない。

```
admin-service  ──write──→ DB ←──read── golang-trainning-frontend (本システム)
```

---

## パッケージ構成

```
pkg/
├── querymodel/          # Query Model（読み取り専用モデル）
│   ├── interface.go     # WomanQueryModel / StoreQueryModel / BlogQueryModel インターフェース
│   ├── woman.go
│   ├── woman_store.go
│   ├── woman_image.go
│   ├── store.go
│   ├── blog.go
│   ├── photo.go
│   └── valueobject/     # 値オブジェクト（BusinessType など）
├── collection/          # ジェネリックコレクション型
├── usecase/
│   ├── input/           # ユースケース入力値の構造体
│   ├── inputport/       # ユースケースのインターフェース定義
│   ├── interactor/      # ユースケースの実装
│   ├── outputport/      # リポジトリのインターフェース定義
│   └── query/           # クエリ条件型
├── adapter/
│   ├── controller/      # HTTP ハンドラ
│   ├── repository/      # DB アクセス実装（Raw SQL + GORM）
│   ├── mapper/          # DB rows → Query Model 変換
│   └── response/        # API レスポンス構造体
├── infrastructure/
│   ├── datastore/       # DB 接続
│   └── router/          # ルーティング設定
├── registry/            # 依存性の組み立て（DI）
├── config/              # 設定値
├── helper/              # 型変換ユーティリティ
└── apperror/            # アプリケーションエラー定義
```

---

## Query Model について

> **Query Model ≒ DTO (Data Transfer Object)**
>
> 一般的に「DTO」と呼ばれる概念と同義。本システムでは CQRS の文脈を明示するため `QueryModel` という名称を採用している。
> ビジネスロジックを持たない読み取り専用のデータ構造であり、画面・API の返却に最適化されている。
> 集約をまたいだデータ（例: `WomanQueryModel` が `WomanStore` を持つ）も自然に表現できる。

---

## DB アクセス方針

- ORM は GORM を使用しているが、複雑なクエリは **Raw SQL** で記述する
- 女性一覧など 1:N JOIN を伴うページングは**サブクエリ**で woman レベルの件数を先に絞り、外側クエリで詳細を取得する
- blood_type / age_range の絞り込みは `pkg/adapter/repository/woman_filter.go` に共通化

---

## API エンドポイント

| Method | Path | 説明 |
|---|---|---|
| GET | `/frontend/stores/:id` | 店舗詳細 |
| GET | `/frontend/stores/:id/women` | 店舗の女性一覧 |
| GET | `/frontend/women` | 女性一覧 |
| GET | `/frontend/women/:id` | 女性詳細 |
| GET | `/frontend/districts/:id/women` | 地区の女性一覧 |
| GET | `/frontend/districts/:id/search-woman-count` | 地区の女性絞り込み件数 |

---

## フロントエンド構成（web/）

### 技術スタック

| 技術 | バージョン | 役割 |
|---|---|---|
| Next.js | 16.2.4 | App Router / SSR / ルーティング |
| React | 19.2.4 | UI ライブラリ |
| TypeScript | ^5 | 型安全 |
| TanStack Query | ^5.99.2 | サーバー状態管理・キャッシュ |
| TanStack Query Devtools | ^5.99.2 | キャッシュデバッグ |

### ディレクトリ構成

```
web/src/
├── app/                          # Next.js App Router
│   ├── layout.tsx                # ルートレイアウト（Providers をラップ）
│   ├── providers.tsx             # QueryClientProvider + ReactQueryDevtools
│   └── districts/
│       └── [id]/
│           └── women/
│               ├── page.tsx              # Server Component（SSR・HydrationBoundary）
│               ├── WomenPageClient.tsx   # Client Component（URL・ルーティング管理）
│               ├── FilterPanel.tsx       # フィルター状態管理・件数リアルタイム表示
│               ├── WomenList.tsx         # 女性一覧表示・ページング
│               └── useDistrictWomen.ts   # useQuery カスタムフック
├── api/
│   ├── base.ts                   # API ベース URL 設定
│   └── woman.ts                  # 女性関連 API 関数
├── components/
│   ├── conditions/
│   │   ├── BloodTypeFilter.tsx   # 血液型チェックボックス
│   │   ├── AgeFilter.tsx         # 年齢チェックボックス
│   │   └── WomanFilterPanel.tsx  # フィルターパネル（表示専用）
│   └── pagination/
│       └── usePageParam.ts       # URL クエリパラメータのページ管理フック
├── hooks/
│   └── useDebounce.ts            # デバウンス汎用フック
└── interfaces/
    ├── district/
    │   ├── womanList.ts          # DistrictWomenResponse 型
    │   └── womanCount.ts         # DistrictWomanCountResponse 型
    └── woman/
        └── list.ts               # WomanListItem など共通型
```

### データフロー（district/[id]/women）

```
page.tsx（Server Component）
  prefetchQuery → dehydrate → HydrationBoundary
    ↓ キャッシュをクライアントへ渡す
WomenPageClient（Client Component）
  URL の blood_type / age_range / page を読み取り props に変換
    ├── FilterPanel
    │     useDebounce(500ms) → useQuery（districtWomanCount）
    │     件数リアルタイム表示 / 検索ボタンで URL 更新
    └── WomenList
          useDistrictWomen（useQuery: districtWomen）
          女性一覧表示 / ページング
```

### キャッシュ戦略

- `staleTime: 30 * 1000`：取得後30秒間は再フェッチしない
- `next: { revalidate: 30 }`：サーバー側 fetch キャッシュで Go サーバーへのリクエストを30秒抑制
- queryKey の配列を sort してフィルター選択順序に関係なくキャッシュを一致させる
