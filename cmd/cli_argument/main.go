package main

import (
	"github.com/gocarina/gocsv"
	cliArgument "github.com/sadmansakib/gractical/cli_argument"
	"github.com/sadmansakib/gractical/common"
	"os"
	"strings"
)

func main() {
	var runners []cliArgument.CliRunner

	err := parseArgs(&runners)

	common.Check(err)

	cliArgument.Execute(runners)
}

func parseArgs(templates *[]cliArgument.CliRunner) error {

	var args = strings.Join(os.Args[1:], " ")
	err := gocsv.UnmarshalString(args, templates)

	common.Check(err)
	return nil
}

