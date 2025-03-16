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

	if pflag.NArg() != 0 {
		fmt.Fprintf(os.Stderr, "expected zero arguments\n")
		os.Exit(1)
	}

	ctx := logger.CtxWithLogger(context.Background(), xlogrus.Default().WithLevel(logLevel))
	logger.Default = func() logger.Logger {
		return logger.FromCtx(ctx)
	}
	defer belt.Flush(ctx)

	k, err := kickcom.New()
	assertNoError(err)

	subcategories, err := k.GetSubcategoriesV1(ctx)
	assertNoError(err)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	err = enc.Encode(subcategories)
	assertNoError(err)
}

func assertNoError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
