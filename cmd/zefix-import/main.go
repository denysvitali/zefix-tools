package main

import (
	"github.com/alexflint/go-arg"
	"github.com/sirupsen/logrus"

	"github.com/denysvitali/documents-processor/pkg/zefix"
)

var args struct {
	Dsn       string `arg:"-d,--dsn,env:ZEFIX_DSN,required"`
	LogLevel  string `arg:"--log-level,env:LOG_LEVEL" default:"info"`
	InputFile string `arg:"-i,--input,env:INPUT_FILE,required"`
}

var logger = logrus.New()

func main() {
	arg.MustParse(&args)
	lvl, err := logrus.ParseLevel(args.LogLevel)
	if err != nil {
		lvl = logrus.InfoLevel
	}
	logger.SetLevel(lvl)

	c, err := zefix.New(args.Dsn)
	if err != nil {
		logger.Fatalf("unable to create zefix client: %v", err)
	}

	err = c.Import(args.InputFile)
	if err != nil {
		logger.Fatalf("unable to import: %v", err)
	}
}
