# Kawasaki Garbage Collection CLI

川崎市のゴミ分別をCLIで検索するツールです。

CLI tool for searching information of garbage collection in Kawasaki City, Japan.

ごみのデータは [川崎市によるごみ分別オープンデータ](https://www.city.kawasaki.jp/300/page/0000075059.html#opendata_dataset_25) を [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/deed.ja) に基づいて利用しています。

Garbage data is based on [open data by Kawasaki city](https://www.city.kawasaki.jp/300/page/0000075059.html#opendata_dataset_25) under [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/legalcode) license.

## How to Use

```
$ kgc "IH調理器"
小物金属(粗大ごみ)
最長辺30cm以上のものは粗大ごみとして出してください。/長辺30cm未満で、30cm×15cmの回収ボックスの投入口に入るものは小型家電としても出すことができます。
URL: http://www.city.kawasaki.jp/kurashi/category/24-1-23-1-1-6-9-0-0-0.html
```

## Installation

```
$ go get github.com/t-chov/kgc
```

## How to Build DB

`/resources` 以下にオープンデータのCSVと、GoのMapに変換するPythonスクリプトが入っています。

`python converter.py original.csv` を実行し、実行結果を `go_gc.csv` と置き換えてください( `go:embed` しています)