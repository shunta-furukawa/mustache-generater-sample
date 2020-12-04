# mustache-generater-sample

このリポジトリは [CyberAgent PTA Advent Calendar](https://adventar.org/calendars/5741) 2020の24日目の記事である、
[YAML+Mustache+go-generate で go の メタプログラミング をする]() のために作られた サンプルコードです。 

## 事前準備 

- [go-yaml](https://github.com/go-yaml/yaml)
- [ Mustache Template Engine for Go ](https://github.com/cbroglie/mustache)
 
## 実行

### コード生成 

``` 
go generate ./... 
``` 

### 実行 

``` 
go run cmd/main.go 
``` 

### 構成

``` 
.
├── LICENSE
├── README.md
├── app
│   ├── gift
│   │   ├── gift.go  -- Gift インターフェース
│   │   └── sportscar.go 
│   └── kid
│       ├── kid.go -- Kid インターフェース
│       └── taro.go
├── cmd
│   └── main.go -- 実行ファイル
└── gen
    ├── app
    │   └── model - YAML を Goで扱うための 構造体
    │       ├── christmas.go
    │       ├── gift.go
    │       └── kid.go
    ├── main.go -- ジェネレータの実行ファイル
    ├── templates -- ジェネレータで用いるテンプレートファイル
    │   ├── christmas.mustache
    │   ├── gift.mustache
    │   └── kid.mustache
    └── variables -- ジェネレータで用いるデータ
        ├── gifts.yml
        ├── gifts.yml.sample
        ├── kids.yml
        └── kids.yml.sample
``` 