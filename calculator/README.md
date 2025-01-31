# Calculator Demo

[Buf CLI](https://buf.build/docs/tutorials/getting-started-with-buf-cli/)

Install local tools:

```bash
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Generate code:

```bash
buf generate
```
