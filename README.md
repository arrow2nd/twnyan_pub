# twnyan

[![arrow2nd](https://circleci.com/gh/arrow2nd/twnyan.svg?style=shield)](https://circleci.com/gh/arrow2nd/twnyan/tree/main)
[![Go Report Card](https://goreportcard.com/badge/github.com/arrow2nd/twnyan)](https://goreportcard.com/report/github.com/arrow2nd/twnyan)
[![GitHub license](https://img.shields.io/github/license/arrow2nd/twnyan)](https://github.com/arrow2nd/twnyan/blob/main/LICENSE.txt)

> **[English](README_EN.md)**

いつでも「にゃーん」したいねこのための Twitter クライアント

![twnyan](https://user-images.githubusercontent.com/44780846/106699506-612c0f80-6626-11eb-803e-332512822789.gif)

## 特徴

- にゃーん
- 疑似 UserStream モード
- マルチバイト文字対応
- 対話モード対応
- とってもカラフル 🎨

## 動作条件

- Windows / macOS / Linux
- 絵文字が表示可能なターミナル

### 備考

- WSL 環境で実行する場合、xdg-open が使用できる必要があります（~微妙な~[解決策](https://qiita.com/arrow2nd/items/5c02a8cdf8197ae15cb7)）

## インストール

### Homebrew

```sh
brew tap arrow2nd/tap
brew install arrow2nd/tap/twnyan
```

### Go

```sh
go install github.com/arrow2nd/twnyan@latest
```

### それ以外

[Releases](https://github.com/arrow2nd/twnyan/releases) からお使いの環境にあった、最新のバイナリファイルをダウンロードしてください。

## アンインストール

### Homebrew

```sh
brew uninstall twnyan
rm -rf ~/.twnyan # 設定ファイル群を削除
```

### Go

```sh
go clean -i github.com/arrow2nd/twnyan
rm -rf $GOPATH/src/github.com/arrow2nd/twnyan
rm -rf ~/.twnyan
```

## 初期設定

![auth](https://user-images.githubusercontent.com/44780846/106747441-4a59dd00-6667-11eb-8248-3468cb39f7d1.png)

1. 初回起動時に表示される認証ページの URL にブラウザでアクセス
2. 画面に従って手順を進め、表示される PIN コードを twnyan に入力
3. 完了！ 😺

## 使い方

### コマンドラインモード

![cmdline](https://user-images.githubusercontent.com/44780846/106699170-b287cf00-6625-11eb-8374-8565286db3e2.gif)

`twnyan [コマンド] [引数]`

一部コマンドはコマンドラインモードでは使用できません。

### にゃーん

![nyaan](https://user-images.githubusercontent.com/44780846/106699001-558c1900-6625-11eb-948e-6212ab0cba40.gif)

`twnyan tw`

ツイート文を省略すると「にゃーん」とツイートされます。

リプライや引用リツイートの時でも同じ動きをします。

画像が添付されている場合はにゃーんしません。

### 対話モード

`twnyan`

タイムライン上のツイートへいいね・RT などの操作を行うことができます。

それぞれツイートの先頭に表示されている番号（tweetnumber）を指定して行います。

## コマンド一覧

<details>
<summary>開く</summary>

## tweet

`エイリアス: tw`

### tweet [テキスト] [画像ファイル]

ツイートを投稿します。

| 引数         | ヒント                                                     | 例                               |
| ------------ | ---------------------------------------------------------- | -------------------------------- |
| テキスト     | テキストと画像ファイルが無い場合「にゃーん」と投稿されます | `tweet`                          |
| 画像ファイル | 複数ある場合は半角スペースで区切って下さい                 | `tweet 🍣 sushi1.png sushi2.png` |

- テキストを省略して、画像のみの投稿も可能です。(e.g. `tweet cat.png`)

### tweet multi [画像ファイル]

`エイリアス: ml`

複数行のツイートを投稿します。

| 引数         | ヒント                                     | 例                            |
| ------------ | ------------------------------------------ | ----------------------------- |
| 画像ファイル | 複数ある場合は半角スペースで区切って下さい | `tweet multi dog.png cat.png` |

- 入力を終了する場合、セミコロン `;` を文末に入力してください
- キャンセルする場合、`:exit` を入力してください

### tweet remove [<ツイート番号>]...

`エイリアス: rm`

ツイートを削除します。

| 引数         | ヒント                                     | 例                 |
| ------------ | ------------------------------------------ | ------------------ |
| ツイート番号 | 複数ある場合は半角スペースで区切って下さい | `tweet remove 2 5` |

## timeline

`エイリアス: tl`

### timeline [取得件数]

ホームタイムラインを表示します。

| 引数     | ヒント                                                   | 例            |
| -------- | -------------------------------------------------------- | ------------- |
| 取得件数 | 省略した場合、設定ファイル内のデフォルト値が指定されます | `timeline 39` |

## stream

`エイリアス: st`

最初に 1 分間ホームタイムラインのツイートを蓄積した後、UserStream API のように 1 分遅れのツイートを表示します。

**Ctrl+C** で終了します。

## mention

`エイリアス: mt`

### mention [取得件数]

自分宛てのメンションを表示します。

| 引数     | ヒント                                                   | 例           |
| -------- | -------------------------------------------------------- | ------------ |
| 取得件数 | 省略した場合、設定ファイル内のデフォルト値が指定されます | `mention 20` |

## list

`エイリアス: ls`

### list [<リスト名>] [取得件数]

リストのタイムラインを表示します。

| 引数     | ヒント                                                   | 例                     |
| -------- | -------------------------------------------------------- | ---------------------- |
| リスト名 | 対話モードで起動している場合、Tab キーで補完が可能です   | `list ねこたち`        |
| 取得件数 | 省略した場合、設定ファイル内のデフォルト値が指定されます | `list "ねこ集会 Ⅱ" 30` |

## user

`エイリアス: ur`

### user [<ユーザー名/ツイート番号>] [取得件数]

指定したユーザーのタイムラインを表示します。

| 引数                    | ヒント                                                   | 例                        |
| ----------------------- | -------------------------------------------------------- | ------------------------- |
| ユーザー名/ツイート番号 | どちらかが指定できます<br>ユーザー名の'@'は省略可能です  | `user github`<br>`user 1` |
| 取得件数                | 省略した場合、設定ファイル内のデフォルト値が指定されます | `user twitter 15`         |

### user own [取得件数]

自分のタイムラインを表示します。

| 引数     | ヒント                                                   | 例            |
| -------- | -------------------------------------------------------- | ------------- |
| 取得件数 | 省略した場合、設定ファイル内のデフォルト値が指定されます | `user own 50` |

## search

`エイリアス: sh`

### search [<キーワード>] [取得件数]

過去 7 日間のツイートを検索します。

| 引数       | ヒント                                                                     | 例                 |
| ---------- | -------------------------------------------------------------------------- | ------------------ |
| キーワード | 先頭が記号、またはスペースを含む場合はダブルクォーテーションで囲んで下さい | `search "cat dog"` |
| 取得件数   | 省略した場合、設定ファイル内のデフォルト値が指定されます                   | `search sushi 5`   |

## like

`エイリアス: lk, fv`

### like [<ツイート番号>]

ツイートにいいね！します。

| 引数         | ヒント                                     | 例         |
| ------------ | ------------------------------------------ | ---------- |
| ツイート番号 | 複数ある場合は半角スペースで区切って下さい | `like 1 2` |

### like remove [<ツイート番号>]

`エイリアス: rm`

ツイートのいいね！を取り消します。

| 引数         | ヒント                                     | 例                |
| ------------ | ------------------------------------------ | ----------------- |
| ツイート番号 | 複数ある場合は半角スペースで区切って下さい | `like remove 1 2` |

## retweet

`エイリアス: rt`

### retweet [<ツイート番号>]...

ツイートをリツイートします。

| 引数         | ヒント                                     | 例            |
| ------------ | ------------------------------------------ | ------------- |
| ツイート番号 | 複数ある場合は半角スペースで区切って下さい | `retweet 1 5` |

### retweet remove [<ツイート番号>]...

`エイリアス: rm`

リツイートを取り消します。

| 引数         | ヒント                                     | 例                   |
| ------------ | ------------------------------------------ | -------------------- |
| ツイート番号 | 複数ある場合は半角スペースで区切って下さい | `retweet remove 1 5` |

## quote

`エイリアス: qt`

### quote [<ツイート番号>] [テキスト] [画像ファイル]

ツイートを引用リツイートします。

| 引数         | ヒント                                                     | 例                      |
| ------------ | ---------------------------------------------------------- | ----------------------- |
| ツイート番号 | 引用するツイートの番号を指定してください                   | `quote 1 これすき`      |
| テキスト     | テキストと画像ファイルが無い場合「にゃーん」と投稿されます | `quote 1`               |
| 画像ファイル | 複数ある場合は半角スペースで区切って下さい                 | `quote 1 🍣 sushi1.png` |

### quote multi

`エイリアス: ml`

複数行の引用リツイートを投稿します。

| 引数         | ヒント                                     | 例                    |
| ------------ | ------------------------------------------ | --------------------- |
| 画像ファイル | 複数ある場合は半角スペースで区切って下さい | `quote multi cat.png` |

- 入力を終了する場合、セミコロン `;` を文末に入力してください
- キャンセルする場合、`:exit` を入力してください

## reply

`エイリアス: rp`

### reply [<ツイート番号>] [テキスト] [画像ファイル]

リプライを投稿します。

| 引数         | ヒント                                                     | 例                                       |
| ------------ | ---------------------------------------------------------- | ---------------------------------------- |
| ツイート番号 | リプライ先のツイートの番号を指定してください               | `reply 1 ねこだ！！！`                   |
| テキスト     | テキストと画像ファイルが無い場合「にゃーん」と投稿されます | `reply 1`                                |
| 画像ファイル | 複数ある場合は半角スペースで区切って下さい                 | `reply 2 寿司みて sushi1.png sushi2.png` |

- テキストを省略して、画像のみの投稿も可能です (e.g. `reply dog.png`)

### reply multi

`エイリアス: ml`

複数行のリプライを投稿します。

| 引数         | ヒント                                     | 例                    |
| ------------ | ------------------------------------------ | --------------------- |
| 画像ファイル | 複数ある場合は半角スペースで区切って下さい | `reply multi dog.png` |

- 入力を終了する場合、セミコロン `;` を文末に入力してください
- キャンセルする場合、`:exit` を入力してください

## follow

`エイリアス: fw`

### follow [<ユーザー名/ツイート番号>]

ユーザーをフォローします。

| 引数                    | ヒント                                                  | 例                            |
| ----------------------- | ------------------------------------------------------- | ----------------------------- |
| ユーザー名/ツイート番号 | どちらかが指定できます<br>ユーザー名の'@'は省略可能です | `follow github`<br>`follow 1` |

### follow remove [<ユーザー名/ツイート番号>]

`エイリアス: rm`

ユーザーのフォローを解除します。

| 引数                    | ヒント                                                  | 例                                             |
| ----------------------- | ------------------------------------------------------- | ---------------------------------------------- |
| ユーザー名/ツイート番号 | どちらかが指定できます<br>ユーザー名の'@'は省略可能です | `follow remove arrow_2nd`<br>`follow remove 1` |

## block

`エイリアス: bk`

### block [<ユーザー名/ツイート番号>]

ユーザーをブロックします。

| 引数                    | ヒント                                                  | 例                             |
| ----------------------- | ------------------------------------------------------- | ------------------------------ |
| ユーザー名/ツイート番号 | どちらかが指定できます<br>ユーザー名の'@'は省略可能です | `block arrow_2nd`<br>`block 1` |

### block remove [<ユーザー名/ツイート番号>]

`エイリアス: rm`

ユーザーのブロックを解除します。

| 引数                    | ヒント                                                  | 例                                           |
| ----------------------- | ------------------------------------------------------- | -------------------------------------------- |
| ユーザー名/ツイート番号 | どちらかが指定できます<br>ユーザー名の'@'は省略可能です | `block remove arrow_2nd`<br>`block remove 1` |

## mute

`エイリアス: mu`

### mute [<ユーザー名/ツイート番号>]

ユーザーをミュートします。

| 引数                    | ヒント                                                  | 例                           |
| ----------------------- | ------------------------------------------------------- | ---------------------------- |
| ユーザー名/ツイート番号 | どちらかが指定できます<br>ユーザー名の'@'は省略可能です | `mute arrow_2nd`<br>`mute 1` |

### mute remove [<ユーザー名/ツイート番号>]

`エイリアス: rm`

ユーザーのミュートを解除します。

| 引数                    | ヒント                                                  | 例                                         |
| ----------------------- | ------------------------------------------------------- | ------------------------------------------ |
| ユーザー名/ツイート番号 | どちらかが指定できます<br>ユーザー名の'@'は省略可能です | `mute remove arrow_2nd`<br>`mute remove 1` |

## open

`エイリアス: op`

### open [<ツイート番号>]

指定したツイートをブラウザで表示します。

| 引数         | ヒント                                             | 例       |
| ------------ | -------------------------------------------------- | -------- |
| ツイート番号 | ブラウザで表示するツイートの番号を指定してください | `open 2` |

## clear

画面を初期化します。

## help

ヘルプを表示します。

また、 `[コマンド] help` とするとコマンドのヘルプが表示されます。

## exit

対話モードを終了します。

</details>

## 設定ディレクトリ

設定ディレクトリはホームディレクトリ直下に`.twnyan`として作成されます。

<details>
<summary>開く</summary>

### .cred.yaml

認証情報のファイルです。

### option.yaml

オプション設定のファイルです。

| 名前       | 説明                   |
| ---------- | ---------------------- |
| ConfigDir  | 設定ディレクトリのパス |
| Counts     | デフォルトの取得件数   |
| DateFormat | 日付のフォーマット     |
| TimeFormat | 時刻のフォーマット     |

- 日付、時刻のフォーマットは[time パッケージのフォーマット文字列](https://golang.org/pkg/time/#pkg-constants)と同じ書式です

### color.yaml

色設定のファイルです。

| 名前         | 説明                   |
| ------------ | ---------------------- |
| Accent1      | アクセント１           |
| Accent2      | アクセント２           |
| Accent3      | アクセント３           |
| Error        | エラーメッセージ背景色 |
| BoxForground | 反転時の文字色         |
| Separator    | セパレータ             |
| UserName     | ユーザー名             |
| ScreenName   | スクリーンネーム       |
| Reply        | リプライ表示           |
| Hashtag      | ハッシュタグ           |
| Favorite     | いいね表示             |
| Retweet      | リツイート表示         |
| Verified     | 認証済みアカウント     |
| Protected    | 鍵アカウント           |
| Following    | フォロー中表示         |
| FollowedBy   | 被フォロー表示         |
| Block        | ブロック表示           |
| Mute         | ミュート表示           |

 </details>
