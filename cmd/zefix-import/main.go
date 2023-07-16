package main

import (
	"github.com/alexflint/go-arg"
	"github.com/denysvitali/documents-processor/pkg/zefix"
	"github.com/sirupsen/logrus"
)

var args struct {
	Dsn       string `arg:"-d,--dsn,env:ZEFIX_IMPORT_DSN,required"`
	InputFile string `arg:"-i,--input,env:ZEFIX_IMPORT_INPUT,required"`
	Debug     *bool  `arg:"-D,--debug,env:ZEFIX_IMPORT_DEBUG"`
}

var logger = logrus.New()

func main() {
	arg.MustParse(&args)

	if args.Debug != nil && *args.Debug {
		logger.SetLevel(logrus.DebugLevel)
	}

	c, err := zefix.New(args.Dsn)
	if err != nil {
		logger.Fatalf("unable to create zefix client: %v", err)
	}

	c.SetLogger(logger)
	err = c.Import(args.InputFile)
	if err != nil {
		logger.Fatalf("unable to import: %v", err)
	}
}
