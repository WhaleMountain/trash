# trash

ゴミ箱

## install

setting GOPATH

```bashrc
export GOPATH=$HOME/go
export PATH=PATH:$GOPATH/bin
```

```shell
$ cd trash
$ go install
```

## 使い方

### ゴミ箱に入れる

* `trash put test.txt`

or

* `trash test.txt`

### ゴミ箱から削除する

* `trash remove test.txt`

or

全て削除
* `trash removeall`

### ゴミ箱から復元する

デフォルトの場所に復元する
* `trash restore test.txt`

or

~/ABC に復元する 
* `trash restore test.txt --restore-path ~/ABC`

### ゴミ箱の一覧を確認する

* `trash list`

### ゴミ箱の設定

* `trash config`

設定する
* `trash config set --delete-time 10` (default 30 day)
* `trash config set --restore-path ~/ABC` (default $HOME)

### help
* `trash help`

or

* `trash -h`