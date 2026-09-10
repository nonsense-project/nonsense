# Protobuf wire types

The `.proto` files define the node P2P and RPC messages. Generated Go files are
committed so building a node or wallet does not require protoc.

To regenerate them, install the Protocol Buffers compiler and the Go protobuf
and gRPC generators matching the source headers, then run `go generate .` in
this directory. Review generated changes for wire compatibility.

`rpc.md` documents the RPC schema. The canonical message definitions are in
`rpc.proto`, `p2p.proto`, and `messages.proto`.
