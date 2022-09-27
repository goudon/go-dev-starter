# 目標

- Goをとりあえず触ってみる（FizzBuzz）
- GoでFile I/O
- Goでスクレイピングしてみる

Reference： https://go.dev/doc/tutorial/getting-started



# 学習メモ

## 1. hello world
go install のインストールとHelloWorld

Ref:
- https://go.dev/dl/
- https://www.tohoho-web.com/ex/golang.html#hello-world

module インストール
go mod init example/hello

実行
go run . (もしくは、 go run hello.go)


(コンパイルファイルの掃除)
go clean -i -n

## 2. FizzBuzz

0から100まで続く連番をコンソールに出力するプログラムを作成する。
ただし、
- 3の倍数のときは、Fizz。
-5 の倍数のときは、Buzz。
-3,5の倍数のときはFizzBuzz。
と表示する。

条件分岐・ループ処理の学習

メモ：Goには、Whileがない
