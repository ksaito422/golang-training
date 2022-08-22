# Golang + gRPC の学習

## 要約

### RPC とは

Remote Procedure Call の略でローカルからリモートの関数を呼び出す発送のこと

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

### gRPC で可能な 4 種類の通信方式

gRPC では「1 リクエスト-1 レスポンス」以外にもストリーミングで N:M になるような通信方式も採用している。

- Unary RPC
  - 1req - 1res の通信方式
- Server streaming RPC
  - 1req - 複数 res の通信方式
  - クライアントがサーバー側からプッシュ通知を受け取る場合とかに採用する
- Client streaming RPC
  - 複数 req - 1res の通信方式
  - クライアント側から複数回に分けてデータをアップロードして、全て受け取った段階でサーバーが 1 回だけ OK と返すような時に採用する
- Bidirectional streaming RPC
  - サーバー・クライアント共に任意のタイミングで req/res を送ることができる通信方式
  - WebSocket のような双方向通信をイメージすると良い

### gRPC のストリーミングを支える技術

HTTP/2 のプロトコル上で実現されているため  
HTTP/2 のフレームフォーマットの中でもタイプフィールドとフラグフィールドが肝となる

## 参考

[作ってわかる!はじめての gRPC](https://zenn.dev/hsaki/books/golang-grpc-starting)
[proto ファイルで定義できる型の公式](https://developers.google.com/protocol-buffers/docs/proto3)
[組み込み以外の便利な型](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf)
[メッセージ型の Go への変換先](https://developers.google.com/protocol-buffers/docs/reference/go-generated)
[メソッド定義がどういうコードに変換されるのか](https://grpc.io/docs/languages/go/generated-code/)
[gRPC のサーバーリフレクションについて](https://github.com/grpc/grpc/blob/master/doc/server-reflection.md)
[gRPC のステータスコード](https://grpc.io/docs/guides/error/#error-status-codes)
