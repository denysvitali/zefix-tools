package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/denysvitali/documents-processor/pkg/zefix"
	"github.com/sirupsen/logrus"
)

var args struct {
	Dsn         string `arg:"-d,--dsn,env:ZEFIX_FIND_DSN,required"`
	Debug       *bool  `arg:"-D,--debug,env:ZEFIX_FIND_DEBUG"`
	CompanyName string `arg:"positional,required"`
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
	company, err := c.FindCompany(args.CompanyName)
	if err != nil {
		logger.Fatalf("unable to find company: %v", err)
	}

	if company == nil {
		logger.Errorf("company not found")
	} else {
		prettyPrintCompany(company)
	}

}

func prettyPrintCompany(company *zefix.Company) {
	fmt.Printf("Name: %s\n", company.Name)
	fmt.Printf("URI: %s\n", company.Uri)
}
