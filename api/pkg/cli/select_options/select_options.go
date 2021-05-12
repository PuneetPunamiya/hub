package select_options

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

type Options struct {
	AskOpts survey.AskOpt
	Name    string
	Catalog string
	Version string
}

func (opts *Options) Ask(resourceInfo string, options []string) error {
	var ans string
	var qs = []*survey.Question{
		{
			Name: resourceInfo,
			Prompt: &survey.Select{
				Message: fmt.Sprintf("Select %s:", resourceInfo),
				Options: options,
			},
		},
	}

	if err := survey.Ask(qs, &ans, opts.AskOpts); err != nil {
		return err
	}

	switch resourceInfo {
	case "catalog":
		opts.Catalog = ans
	case "task":
		opts.Name = ans
	case "version":
		opts.Version = ans
	}

	return nil
}
