第1回大阪Go勉強会
============

## 告知
[Zusaar](http://www.zusaar.com/event/1010007)


## 資料
1. [第1回大阪Go勉強会 Go紹介](http://www.slideshare.net/AtsushiKanaya/1go)
1. [速習A tour of Go](http://www.slideshare.net/AtsushiKanaya/a-tour-of-go-26263016)
1. GDG Kobeでの活動内容の紹介(特にGo言語関連)


## ハンズオン課題

### FizzBuzz
おなじみのFizzBuzzをGoで解いてみてください。
* 1から順番に数を表示する
* その数が3で割り切れるなら数字の代わりにFizzと表示する
* その数が5で割り切れるなら数字の代わりにBuzzと表示する
* その数が3でも5でも割り切れるなら数字の代わりにFizzBuzzと表示する

速習A tour of Goの11ページまでで解けると思います。


### ワードカウント
標準入力から英文テキストファイルを読み込み、各英単語の出現回数を表示してください。

ヒント: 標準入力から1行毎に入力を得る方法は下記サンプルコードをご参照ください。

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		//TODO write your logic here
		fmt.Println(text)
	}
}
```

ヒント2: 文字列操作関連のパッケージは[stringsパッケージ](http://golang.org/pkg/strings/)です。

速習A tour of Goの17ページまでで解けると思います。
合わせて、[bufioパッケージ](http://golang.org/pkg/bufio/)、[osパッケージ](http://golang.org/pkg/os/)もご参照ください。


### 指定日がその年の何日目かを計算するプログラム
与えた年月日が、その年の何日目か計算するプログラムを作成してください。
また、その際に、テストコードも書いてください。テストドリブン推奨。

* 平年の場合は、1月1日を1日目、12月31日を365日目とする
* うるう年の場合は、2月29日があるので、12月31日が366日目となる

テストの書き方は速習A tour of Goの最終ページにありますが、
1つだけうるう年判定のサンプルを置きます(dayofyear_test.go)。
```go
package main

import (
	"testing"
)

func TestIsLeapYear(t *testing.T) {
	if IsLeapYear(2013) {
		t.Error("2013 is not a leap-year!")
	}
	if IsLeapYear(2100) {
		t.Error("2100 is not a leap-year!")
	}
	if !IsLeapYear(2012) {
		t.Error("2012 is a leap-year!")
	}
	if !IsLeapYear(2000) {
		t.Error("2000 is a leap-year!")
	}
}
```


### HTTPでJSONデータの取得
ISBNコードを引数にして、オライリーのサイトから書籍情報を取得して表示してください。
例えば、「入門 機械学習」の書籍情報を取得するためのURLは下記になります。
下記URLのうち13桁の数字がISBNコードになります。
http://www.oreilly.co.jp/books/9784873115948/biblio.json

ヒント1: JSONからGoのstructにマッピングさせる方法は下記をご参照ください。
[GoでJSONシリアライズ・デシリアライズ](http://qiita.com/todogzm/items/b2250aa2acc12cdd50f7)

ヒント2: HTTPクライアントは[net/httpパッケージ](http://golang.org/pkg/net/http/)をご参照ください。
特にExamplesでJSONへのデシリアライズに必要な[]bytesの取得までが書かれています。

上記ヒントと、速習A tour of Goの17ページまでで解けると思います。


### 上記コードの発展・複数ISBNコードを指定して総ページ数・総額を計算するプログラム
上記コードができたら、次はISBNコードを複数件(サーバ負荷を考慮し1回あたり5件以内推奨)指定し、
指定コードの書籍の総ページ数・総額を計算して表示してください。
また、複数件取得する部分を並行処理化してください。

テストとして、下記データをご使用いただき、総ページ数・総額をご確認ください。
* 9784873115948 : 344ページ, 3360円
* 9784873116402 : 310ページ, 2940円
* 9784873114750 : 392ページ, 3570円

ヒント1: mainのgoroutineは、自分の処理が終了次第、他goroutineが動いていてもプログラムを終了します。


## 参考までに解答例
解答例をアップしました。
https://github.com/todoa2c/hango/tree/master/vol01

