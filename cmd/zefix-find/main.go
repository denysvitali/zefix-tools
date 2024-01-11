package main

import (
	"fmt"

	"github.com/alexflint/go-arg"
	"github.com/sirupsen/logrus"

	"github.com/denysvitali/zefix-tools/pkg/zefix"
)

var args struct {
	Dsn         string `arg:"-d,--dsn,env:ZEFIX_DSN,required"`
	LogLevel    string `arg:"--log-level,env:LOG_LEVEL" default:"info"`
	CompanyName string `arg:"positional,required"`
}

var logger = logrus.StandardLogger()

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
