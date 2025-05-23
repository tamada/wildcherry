package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
	"github.com/tamada/wildcherry"
)

type PrintOptions struct {
	humanize bool
	format   string
	help     bool
}

type options struct {
	targets  *wildcherry.CountingTargets
	printer  *PrintOptions
	logLevel string
}

func helpMessage(command string) string {
	return fmt.Sprintf(`%s [OPTIONS] [FILEs|URLs|DIRs...]
    -b, --bytes              Count bytes
    -c, --characters         Count characters
    -w, --words              Count words
    -l, --lines              Count lines

    -L, --log <LEVEL>        Set the log level (default "info")
                             Available: trace, debug, info, warn, error
    -f, --format <FORMAT>    Output format (default "default")
                             Available: default, json, xml
    -H, --humanize           Humanize the output
    -h, --help               Print this message`, command)
}

func buildFlagSet() (*flag.FlagSet, *options) {
	opts := &options{targets: &wildcherry.CountingTargets{}, printer: &PrintOptions{}, logLevel: "info"}
	flags := flag.NewFlagSet("wildcherry", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage("wildcherry")) }
	flags.StringVarP(&opts.logLevel, "log", "L", "info", "Set the log level")
	flags.BoolVarP(&opts.targets.Bytes, "bytes", "b", false, "Count bytes")
	flags.BoolVarP(&opts.targets.Characters, "characters", "c", false, "Count characters")
	flags.BoolVarP(&opts.targets.Words, "words", "w", false, "Count words")
	flags.BoolVarP(&opts.targets.Line, "lines", "l", false, "Count lines")
	flags.BoolVarP(&opts.printer.humanize, "humanize", "H", false, "Humanize the output")
	flags.StringVarP(&opts.printer.format, "format", "f", "default", "Output format")
	flags.BoolVarP(&opts.printer.help, "help", "h", false, "Print this message")

	return flags, opts
}

func hello() string {
	return "Welcome to WildCherry!"
}

func perform(_ *options, _ []string) error {
	fmt.Println(hello())
	return nil
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
}

func goMain(args []string) int {
	flags, opts := buildFlagSet()
	if err := flags.Parse(args); err != nil {
		printError(err)
		return 1
	}
	if opts.printer.help {
		flags.Usage()
		return 0
	}
	if err := perform(opts, flags.Args()); err != nil {
		printError(err)
		return 1
	}
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
