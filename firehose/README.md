# Bluesky Firehose

Package **firehose** implements a client for the [Bluesky Firehose API](https://docs.bsky.app/docs/advanced-guides/firehose), for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-bsky/firehose

[![GoDoc](https://godoc.org/github.com/reiver/go-bsky/firehose?status.svg)](https://godoc.org/github.com/reiver/go-bsky/firehose)

## Examples

```golang
package main

import (
	"fmt"
	"os"

	"github.com/gorilla/websocket"
	"github.com/reiver/go-bsky/firehose"
)

func main() {

	// The Bluesky Firehose API use WebSockets.
	//
	// This is the URL that we will connect to it.
	const uri string = firehose.WebSocketURI

	// Connect to the WebSocket.
	conn, _, err := websocket.DefaultDialer.Dial(uri, http.Header{})
	if nil != err {
		fmt.Fprintf(os.Stderr, "ERROR: could not connect to Bluesky Firehose API at %q: %s \n", uri, err)
		return
	}
	defer conn.Close() // <-- we need to eventually close the WebSocket, so that we don't have a resource leak.

	// A WebSocket returns a series of messages.
	//
	// Here we loop, read each message from the WebSocket one-by-one.
	for {
		// Here we are just getting the raw binary data.
		// Later we decode it.
		wsMessageType, wsMessage, err := conn.ReadMessage()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: problem reading from WebSocket for the connection to the Bluesky Firehose API at %q: %s \n", uri, err)
			return
		}

		// Technically a WebSocket message can either be 'text' message, a 'binary' message, or a few other control messages.
		// We expect the Bluesky Firehose API to only return binary messages.
		//
		// If we receive a message from the WebSocket that is not binary, then we will just ignore it.
		if websocket.BinaryMessage != wsMessageType {
			continue
		}

		// Here we turn the WebSocket message into a Firehose Message.
		var message firehose.Message = firehose.Message(wsMessage)

		// Here we decode the message.
		var header firehose.MessageHeader
		var payload firehose.MessagePayload = firehose.MessagePayload{}
		err := message.Decode(&header, &paylooad)
		if nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: problem decoding WebSocket message from the connection to the Bluesky Firehose API at %q: %s \n", uri, err)
			continue
		}

		//@TODO: Do whatever you want to do with decode message-header and message-payload..
	}
}
```

## Import

To import package **firehose** use `import` code like the follownig:
```
import "github.com/reiver/go-bsky/firehose"
```

## Installation

To install package **firehose** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-bsky/firehose
```

## Author

DecodePackage **firehose** was written by [Charles Iliya Krempeaux](http://reiver.link)
