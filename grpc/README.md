# Golang + gRPC の学習

## 要約

### RPC とは

- Remote Procedure Call の略でローカルからリモートの関数を呼び出す発送のこと

### gRPC とは

- RPC を具現化するための方式の一つが gRPC
- 実現するために HTTP/2 + Protocol Buffers を使う
  - HTTP/2 の POST リクエストとレスポンスを使う
    - 呼び出す関数の情報: リクエストパスに含める
    - 呼び出す関数に渡す引数: HTTP リクエストボディに含める
    - 呼び出した関数の戻り値: HTTP レスポンスボディに含める
  - Protocol Buffers はシリアライズ方式のこと

### Protocol Buffer Language とは

スキーマ定義言語のこと

NOTE: swagger と似たようなニュアンスか？

## 参考

[作ってわかる!はじめての gRPC](https://zenn.dev/hsaki/books/golang-grpc-starting)
[proto ファイルで定義できる型の公式](https://developers.google.com/protocol-buffers/docs/proto3)
[組み込み以外の便利な型](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf)
[メッセージ型の Go への変換先](https://developers.google.com/protocol-buffers/docs/reference/go-generated)
[メソッド定義がどういうコードに変換されるのか](https://grpc.io/docs/languages/go/generated-code/)
