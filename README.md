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
