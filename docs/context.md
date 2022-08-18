# context パッケージについて

## 概要

### Context 型の役割とは？

- 処理の締め切りを伝達
- キャンセル信号の伝搬
- リクエストスコープ値の伝達

### Context の意義

- 一つの処理が複数のゴールーチンをまたいで行われる場合に役立つ

### context の定義

```
// context.Context型

type Context interface {
	Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key any) any
}
```

4 つのメソッドが確認できる。

## 参考

[go context](https://pkg.go.dev/context#section-documentation)
[go routine についてわかりやすそうな記事](https://zenn.dev/mikankitten/articles/6344d71f4f4920)
