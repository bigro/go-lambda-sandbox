# go-lambda-sandbox
以下の記事を写経した。  
[go modules のつかいかたメモ - castaneaiのブログ](https://castaneai.hatenablog.com/entry/2019/02/22/151213)

## デプロイ方法
```
go build
zip go-lambda-sandbox.zip go-lambda-sandbox
```
出来上がったzipファイルをlambdaにアップする。
