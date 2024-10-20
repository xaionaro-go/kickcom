package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/facebookincubator/go-belt"
	"github.com/facebookincubator/go-belt/tool/logger"
	xlogrus "github.com/facebookincubator/go-belt/tool/logger/implementation/logrus"
	"github.com/spf13/pflag"
	"github.com/xaionaro-go/kickcom/pkg/kickcom"
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

	msgs, err := k.GetChatMessagesV2(ctx, channel.ID, 0)
	assertNoError(err)

	for _, msg := range msgs.Data.Messages {
		spew.Dump(msg)
	}
}

func assertNoError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
