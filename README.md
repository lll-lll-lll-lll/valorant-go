
### なぜ作ろうと思ったか
- valorant apiを叩いてデータ分析してみたいと思ったのがきっかけ
- いろいろ調べていくと、プロダクションとして何かしら作っていないとデータが取れないみたいなので、プレビューとして最低限Valorant APIを
叩けるようにした。自分のデータ分析プロジェクトの案がプロダクションとして認められたら継続して開発していこうと思う。

### 意識したこと
- できる限りGoらしい書き方を日頃から意識している。振り返るときに統一された書き方をしているとコードが読みやすい。他の人のコードを読む時に
読みやすさを実感する。
    - [google best practice](https://google.github.io/styleguide/go/#about)
    - [uber](https://github.com/uber-go/guide)
- http.Client.Doのhttp.Responseをbindする際に内部で条件分岐して内部で生成した変数にバインドして返してる。
バインドメソッドの引数の値が変わるのは避けたかったので、内部で定義してバインドさせた変数を返すようにした。

### HOW TO GET API KEY
URL: https://developer.riotgames.com/

### HOW TO USE

```sh
export VALORANT_API_KEY=VALORANT_API_KEY
go run example/main.go
```