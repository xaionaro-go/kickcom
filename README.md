# About

[![Go Reference](https://godoc.org/github.com/xaionaro-go/kickcom?status.svg)](https://godoc.org/github.com/xaionaro-go/kickcom)
[![Go Report Card](https://goreportcard.com/badge/github.com/xaionaro-go/kickcom?branch=main)](https://goreportcard.com/report/github.com/xaionaro-go/kickcom)

This is package for Go that implements a client to Kick.com API.

Currently, it supports only a couple of features, but feel free to extend it.

# Backstory

Unfortunately, Kick stuff does not respond me for more than a year on a request to provide docs to their API, and on questions "when?" they respond with:

> [...] we do not have an estimated time of when our relevant team will respond to your request. We may only advise you to remain patient and they will respond to your request in due time.

So I had to start reverse engineering how it works. This package is the result of this effort.

# How to use

An example how to get the current chat messages in the chat, given the channel's slug:
```go
package main

import "github.com/davecgh/go-spew/spew"
import "github.com/xaionaro-go/kickcom"

func main() {
    ...

	k, err := kickcom.New()
	if err != nil { ... }

	channel, err := k.GetChannelV1(ctx, channelSlug)
	if err != nil { ... }

	reply, err := k.GetChatMessagesV2(ctx, channel.ID, 0)
	if err != nil { ... }

	for _, msg := range reply.Data.Messages {
		spew.Dump(msg)
	}

    ...
}
```
