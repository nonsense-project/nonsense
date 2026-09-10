# wire

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://choosealicense.com/licenses/isc/)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/nonsense-project/nonsense/wire)

Package wire implements the nonsense wire protocol.

## Nonsense Message Overview

The nonsense protocol consists of exchanging messages between peers.
Each message is preceded by a header which identifies information
about it such as which nonsense network it is a part of, its type, how
big it is, and a checksum to verify validity. All encoding and
decoding of message headers is handled by this package.

To accomplish this, there is a generic interface for nonsense messages
named `Message` which allows messages of any type to be read, written,
or passed around through channels, functions, etc. In addition,
concrete implementations of most all nonsense messages are provided.
All of the details of marshalling and unmarshalling to and from the
wire using nonsense encoding are handled so the  caller doesn't have
to concern themselves with the specifics.

## Reading Messages Example

In order to unmarshal nonsense messages from the wire, use the
`ReadMessage` function. It accepts any `io.Reader`, but typically
this will be a `net.Conn` to a remote node running a nonsense peer.
Example syntax is:

```Go
// Use the most recent protocol version supported by the package and the
// main nonsense network.
pver := wire.ProtocolVersion
nonsensenet := wire.Mainnet

// Reads and validates the next nonsense message from conn using the
// protocol version pver and the nonsense network nonsensenet. The returns
// are a appmessage.Message, a []byte which contains the unmarshalled
// raw payload, and a possible error.
msg, rawPayload, err := wire.ReadMessage(conn, pver, nonsensenet)
if err != nil {
	// Log and handle the error
}
```

See the package documentation for details on determining the message
type.

## Writing Messages Example

In order to marshal nonsense messages to the wire, use the
`WriteMessage` function. It accepts any `io.Writer`, but typically
this will be a `net.Conn` to a remote node running a nonsense peer.
Example syntax to request addresses from a remote peer is:

```Go
// Use the most recent protocol version supported by the package and the
// main bitcoin network.
pver := wire.ProtocolVersion
nonsensenet := wire.Mainnet

// Create a new getaddr nonsense message.
msg := wire.NewMsgGetAddr()

// Writes a nonsense message msg to conn using the protocol version
// pver, and the nonsense network nonsensenet. The return is a possible
// error.
err := wire.WriteMessage(conn, msg, pver, nonsensenet)
if err != nil {
	// Log and handle the error
}
```
