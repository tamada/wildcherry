package main

import (
	"errors"
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
	"github.com/tamada/wildcherry"
	"github.com/tamada/wildcherry/fs"
	"github.com/tamada/wildcherry/url"
)

type PrintOptions struct {
	humanize bool
	format   string
	help     bool
}

type options struct {
	readOpts   *wildcherry.Option
	completion bool
	printer    *PrintOptions
	logLevel   string
}

func helpMessage(command string) string {
	return fmt.Sprintf(`%s [OPTIONS] [FILEs|URLs|DIRs...]
	-b, --bytes              Count bytes
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
	opts := &options{readOpts: wildcherry.NewOption(), printer: &PrintOptions{}, logLevel: "info"}
	flags := flag.NewFlagSet("wildcherry", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage("wildcherry")) }
	flags.StringVarP(&opts.logLevel, "log", "L", "info", "Set the log level")
	flags.BoolVarP(&opts.readOpts.T.Bytes, "bytes", "b", false, "Count bytes")
	flags.BoolVarP(&opts.readOpts.T.Words, "words", "w", false, "Count words")
	flags.BoolVarP(&opts.readOpts.T.Line, "lines", "l", false, "Count lines")
	flags.BoolVarP(&opts.printer.humanize, "humanize", "H", false, "Humanize the output")
	flags.StringVarP(&opts.printer.format, "format", "f", "default", "Output format")
	flags.BoolVarP(&opts.completion, "generate-completions", "", false, "Generate shell completions")
	flags.BoolVarP(&opts.printer.help, "help", "h", false, "Print this message")

	return flags, opts
}

func hello() string {
	return "Welcome to WildCherry!"
}

func printResults(results []*wildcherry.Result, errs []error, opts *PrintOptions) error {
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	for _, r := range results {
		fmt.Printf(" %7d %7d %7d %s\n",
			r.Lines, r.Words, r.Bytes, r.S.Name())
	}
	return nil
}

func perform(opts *options, args []string) error {
	list, err := listTargets(args, opts.readOpts)
	if err != nil {
		return err
	}
	if len(list) == 0 {
		return errors.New("fatal: no sources provided")
	}
	errs := []error{}
	results := []*wildcherry.Result{}
	for _, source := range list {
		r, e := wildcherry.Count(source, opts.readOpts.T)
		if e != nil {
			errs = append(errs, e)
		} else {
			results = append(results, r)
		}
	}
	return printResults(results, errs, opts.printer)
}

func listTargets(args []string, opts *wildcherry.Option) ([]wildcherry.Source, error) {
	if len(args) == 0 {
		return []wildcherry.Source{wildcherry.NewStdinSource()}, nil
	}
	var sources []wildcherry.Source
	for _, arg := range args {
		list, err := NewSource(arg, opts)
		if err != nil {
			return nil, err
		}
		sources = append(sources, list...)
	}
	return sources, nil
}

func NewSource(path string, opts *wildcherry.Option) ([]wildcherry.Source, error) {
	if path == "" {
		return nil, fmt.Errorf("given path is empty")
	}
	if url.IsURL(path) {
		s, err := url.New(path)
		if err != nil {
			return nil, err
		}
		return []wildcherry.Source{s}, nil
	}
	return fs.New(path, opts)
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
}

func updateOptions(opts *options) {
	r := opts.readOpts
	if !r.T.Bytes && !r.T.Words && !r.T.Line {
		r.T.Bytes = true
		r.T.Words = true
		r.T.Line = true
	}
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
	if opts.completion {
		if err := generateCompletions(flags, "wildcherry"); err != nil {
			printError(err)
			return 1
		}
		return 0
	}
	updateOptions(opts)
	if err := perform(opts, flags.Args()[1:]); err != nil {
		printError(err)
		return 1
	}
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
