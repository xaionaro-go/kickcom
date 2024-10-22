package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/facebookincubator/go-belt"
	"github.com/facebookincubator/go-belt/tool/logger"
	xlogrus "github.com/facebookincubator/go-belt/tool/logger/implementation/logrus"
	"github.com/spf13/pflag"
	"github.com/xaionaro-go/kickcom"
)

func main() {
	logLevel := logger.LevelInfo
	pflag.Var(&logLevel, "log-level", "")
	pflag.Parse()

	if pflag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "expected 1 argument\n")
		os.Exit(1)
	}

	channelSlug := pflag.Arg(0)

	ctx := logger.CtxWithLogger(context.Background(), xlogrus.Default().WithLevel(logLevel))
	logger.Default = func() logger.Logger {
		return logger.FromCtx(ctx)
	}
	defer belt.Flush(ctx)

	k, err := kickcom.New()
	assertNoError(err)

	channel, err := k.GetChannelV1(ctx, channelSlug)
	assertNoError(err)

	currentCursor := uint64(0)

	t := time.NewTicker(time.Second)
	defer t.Stop()
	for ; ; <-t.C {
		reply, err := k.GetChatMessagesV2(ctx, channel.ID, currentCursor)
		assertNoError(err)

		if reply.Status.Code != 200 {
			if reply.Status.Message == "No messages found" {
				continue
			}
		}

		for _, msg := range reply.Data.Messages {
			spew.Dump(msg)
		}

		break // the cursor mechanism does not work, so let's stop here :(

		newCursor, err := strconv.ParseUint(reply.Data.Cursor, 10, 64)
		assertNoError(err)

		currentCursor = newCursor
	}
}

func assertNoError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
