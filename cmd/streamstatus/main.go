package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

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

	streamInfo, err := k.GetLivestreamV2(ctx, channelSlug)
	assertNoError(err)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	err = enc.Encode(streamInfo)
	assertNoError(err)
}

func assertNoError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
