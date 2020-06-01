// Code generated by goa v3.1.3, DO NOT EDIT.
//
// hub HTTP client CLI support package
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	categoryc "github.com/tektoncd/hub/api/gen/http/category/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `category all
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` category all` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		categoryFlags = flag.NewFlagSet("category", flag.ContinueOnError)

		categoryAllFlags = flag.NewFlagSet("all", flag.ExitOnError)
	)
	categoryFlags.Usage = categoryUsage
	categoryAllFlags.Usage = categoryAllUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "category":
			svcf = categoryFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "category":
			switch epn {
			case "all":
				epf = categoryAllFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "category":
			c := categoryc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "all":
				endpoint = c.All()
				data = nil
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// categoryUsage displays the usage of the category command and its subcommands.
func categoryUsage() {
	fmt.Fprintf(os.Stderr, `The category service gives category details
Usage:
    %s [globalflags] category COMMAND [flags]

COMMAND:
    all: Get all Categories with their tags sorted by name

Additional help:
    %s category COMMAND --help
`, os.Args[0], os.Args[0])
}
func categoryAllUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] category all

Get all Categories with their tags sorted by name

Example:
    `+os.Args[0]+` category all
`, os.Args[0])
}
