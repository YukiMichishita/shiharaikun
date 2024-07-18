## 仕様
https://github.com/upsidr/coding-test/blob/main/web-api-language-agnostic/README.ja.md

## ローカル環境構築

### .envファイルの作成
```
cp .env.example .env.local
```

### DB起動
```
make run_mysql
```

### DB初期化
```
make init_db
```

### DBの状態からコード生成
```
make gorm_gen
```

### サーバー起動
```
make run_server
```


## ディレクトリ構成
書籍「[Clean Architecture](https://tatsu-zine.com/books/clean-architecture)」に記載の同心円状のアーキテクチャ例を念頭に置いて、以下のような構成にしています。

| ディレクトリ  | 対応する層                    |
|:--------|:-------------------------|
| domain | Enterprise Business Rule |
| usecase | Application Business Rule |
| adapter | Interface Adapters | 
| infrastructure | Frameworks & Drivers | 
