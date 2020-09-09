# trash

ゴミ箱

## build

`GOOS=linux GOARCH=amd64 go build main.go `

## 使い方

### ゴミ箱に入れる

`./main put test.txt`

or

`./main test.txt`

### ゴミ箱から削除する

`./main remove test.txt`

or

全て削除
`./main removeall`

### ゴミ箱から復元する

デフォルトの場所に復元する
`./main restore test.txt`

or

~/ABC に復元する 
`./main restore test.txt --restore-path ~/ABC`

### ゴミ箱の一覧を確認する

`./main list`

### ゴミ箱の設定

`./main config`

設定する
`./main config set --delete-time 10` (デフォルト 30日)
`./main config set --restore-path ~/ABC` (デフォルト $HOME)

### help
`./main help`

or

`./main -h`