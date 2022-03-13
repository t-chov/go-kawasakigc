# Kawasaki Garbage Collection CLI

---

川崎市のゴミ分別をCLIで検索するツールです。

CLI tool for searching information of garbage collection in Kawasaki City, Japan.

ごみのデータは [川崎市によるごみ分別オープンデータ](https://www.city.kawasaki.jp/300/page/0000075059.html#opendata_dataset_25) を [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/deed.ja) に基づいて利用しています。

Garbage data is based on [open data by Kawasaki city](https://www.city.kawasaki.jp/300/page/0000075059.html#opendata_dataset_25) under [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/legalcode) license.

## How to Build DB

`/resources` 以下にオープンデータのCSVと、GoのMapに変換するPythonスクリプトが入っています。

`python converter.py gc.csv` を実行し、 `db/db.go` の `InitDb` にコピペします。