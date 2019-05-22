package args

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)

func Args(args []string) *string {
	parser := argparse.NewParser("flags", "")

	m := parser.String("m", "message", &argparse.Options{Default: nil, Help: "commit message"})

	err := parser.Parse(args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	return m
}
