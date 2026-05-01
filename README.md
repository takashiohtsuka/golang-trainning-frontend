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
│   ├── interface.go
│   ├── woman.go
│   ├── woman_store.go
│   ├── woman_image.go
│   ├── store.go
│   ├── blog.go
│   ├── photo.go
│   ├── prefecture.go
│   ├── district.go
│   ├── business_type.go
│   ├── immediate_available_woman.go
│   └── valueobject/     # 値オブジェクト（BusinessType など）
├── collection/          # ジェネリックコレクション型
├── usecase/
│   ├── input/           # ユースケース入力値の構造体
│   ├── inputport/       # ユースケースのインターフェース定義
│   ├── interactor/      # ユースケースの実装
│   ├── outputport/      # リポジトリのインターフェース定義
│   └── query/           # クエリ条件型（Condition / KindWhereBetweenOr など）
├── adapter/
│   ├── controller/      # HTTP ハンドラ
│   ├── repository/      # DB アクセス実装（Raw SQL + GORM）
│   ├── mapper/          # DB rows → Query Model 変換
│   ├── request/         # リクエスト構造体
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
>
> 単一アイテムを返す用途では **NilObject パターン** を採用し、`nil` チェックを排除している。

---

## DB アクセス方針

- ORM は GORM を使用しているが、複雑なクエリは **Raw SQL** で記述する
- 女性一覧など 1:N JOIN を伴うページングはサブクエリで woman レベルの件数を先に絞り、外側クエリで詳細を取得する
- 絞り込み条件は `[]query.Condition` で統一し、`buildWhereClause` で SQL に変換する
- 年齢などの OR グループ BETWEEN 条件は `KindWhereBetweenOr` で表現する

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
| GET | `/frontend/immediate_available_women` | 即予約女性一覧 |
| GET | `/frontend/prefectures` | 都道府県一覧 |
| GET | `/frontend/prefectures/:id/districts` | 都道府県に属するエリア一覧 |
| GET | `/frontend/business_types` | 業種一覧 |

---

## フロントエンド構成（web/）

### 技術スタック

| 技術 | バージョン | 役割 |
|---|---|---|
| Next.js | 16.2.4 | App Router / SSR / ルーティング |
| React | 19.2.4 | UI ライブラリ |
| TypeScript | ^5 | 型安全 |
| TanStack Query | ^5.99.2 | サーバー状態管理・キャッシュ |

### ディレクトリ構成

```
web/src/
├── api/                              # バックエンドへの fetch 関数
│   ├── base.ts                       # BASE_URL・fetchJSON 共通ヘルパー
│   ├── woman.ts                      # 女性一覧・件数取得
│   ├── immediateAvailableWoman.ts    # 即予約女性一覧取得
│   ├── prefecture.ts                 # 都道府県一覧取得
│   ├── district.ts                   # エリア一覧取得
│   └── businessType.ts               # 業種一覧取得
├── interfaces/                       # API レスポンスの型定義
│   ├── woman/
│   ├── district/
│   ├── prefecture/
│   ├── immediateAvailableWoman/
│   └── businessType/
├── components/
│   ├── conditions/                   # フィルター系共通コンポーネント
│   │   ├── FilterPanelShell.tsx      # HTML 構造（children + 検索ボタン）
│   │   ├── PrefectureSelect.tsx      # 都道府県セレクト
│   │   ├── DistrictSelect.tsx        # エリアセレクト（prefecture_id 連動 fetch）
│   │   ├── BusinessTypeFilter.tsx    # 業種チェックボックス
│   │   ├── BloodTypeFilter.tsx       # 血液型チェックボックス（parseBloodTypes をエクスポート）
│   │   └── AgeFilter.tsx             # 年齢チェックボックス（parseAgeRanges をエクスポート）
│   └── pagination/
│       └── usePageParam.ts           # URL クエリパラメータのページ管理フック
├── hooks/
│   └── useDebounce.ts
└── app/
    ├── layout.tsx                    # ルートレイアウト（QueryClientProvider）
    ├── districts/[id]/women/         # エリア別女性一覧ページ
    │   ├── page.tsx                  # Server Component（SSR プリフェッチ）
    │   ├── WomenPageClient.tsx       # URL state 管理
    │   ├── FilterPanel.tsx           # フィルター（血液型・年齢・件数リアルタイム表示）
    │   ├── WomenList.tsx             # 一覧表示・ページネーション
    │   └── useDistrictWomen.ts       # React Query hook
    └── immediate_available_women/    # 即予約女性一覧ページ
        ├── page.tsx                  # Server Component（force-dynamic）
        ├── ImmediateAvailableWomenPageClient.tsx  # URL state 管理
        ├── FilterPanel.tsx           # フィルター（都道府県・エリア・業種・血液型・年齢）
        ├── ImmediateAvailableWomenList.tsx        # 一覧表示・ページネーション
        └── useImmediateAvailableWomen.ts          # React Query hook
```

### フィルターコンポーネント設計方針

各フィルターを独立した単一責任のコンポーネントとして実装している。`FilterPanelShell` が HTML 構造を担い、各ページの `FilterPanel` が必要なコンポーネントを組み合わせる。

```tsx
<FilterPanelShell onSearch={handleSearch}>
  <PrefectureSelect ... />
  <DistrictSelect ... />   {/* prefectureId に連動して districts を fetch */}
  <BusinessTypeFilter ... />
  <BloodTypeFilter ... />
  <AgeFilter ... />
</FilterPanelShell>
```

### キャッシュ戦略

- `staleTime: 30 * 1000`：取得後30秒間は再フェッチしない
- `next: { revalidate: 30 }`：サーバー側 fetch キャッシュで Go サーバーへのリクエストを抑制
- queryKey の配列を sort してフィルター選択順序に関係なくキャッシュを一致させる
