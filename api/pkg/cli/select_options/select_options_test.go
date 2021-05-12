package select_options

import (
	"testing"

	"github.com/AlecAivazis/survey/v2/terminal"
	goexpect "github.com/Netflix/go-expect"
	"github.com/tektoncd/cli/test/prompt"
)

func TestOptions_Ask(t *testing.T) {
	options := []string{
		"buildah",
		"git-clone",
		"tkn",
	}

	options1 := []string{
		"foo",
		"baar",
	}

	options2 := []string{
		"0.1",
		"0.2",
	}

	testParams := []struct {
		name     string
		resource string
		prompt   prompt.Prompt
		options  []string
		want     Options
	}{
		{
			name:     "select task name",
			resource: "task",
			prompt: prompt.Prompt{
				CmdArgs: []string{},
				Procedure: func(c *goexpect.Console) error {
					if _, err := c.ExpectString("Select task:"); err != nil {
						return err
					}
					if _, err := c.SendLine(options[0]); err != nil {
						return err
					}
					return nil
				},
			},
			options: options,
			want: Options{
				Name:    "buildah",
				Catalog: "",
				Version: "",
			},
		},
		{
			name:     "select catalog",
			resource: "catalog",
			prompt: prompt.Prompt{
				CmdArgs: []string{},
				Procedure: func(c *goexpect.Console) error {
					if _, err := c.ExpectString("Select catalog:"); err != nil {
						return err
					}
					if _, err := c.SendLine(options1[0]); err != nil {
						return err
					}
					return nil
				},
			},
			options: options1,
			want: Options{
				Name:    "",
				Catalog: "foo",
				Version: "",
			},
		},
		{
			name:     "select version",
			resource: "version",
			prompt: prompt.Prompt{
				CmdArgs: []string{},
				Procedure: func(c *goexpect.Console) error {
					if _, err := c.ExpectString("Select version:"); err != nil {
						return err
					}
					if _, err := c.SendLine(options2[0]); err != nil {
						return err
					}
					return nil
				},
			},
			options: options2,
			want: Options{
				Name:    "",
				Catalog: "",
				Version: "0.1",
			},
		},
	}

	for _, tp := range testParams {
		t.Run(tp.name, func(t *testing.T) {
			opts := &Options{}
			tp.prompt.RunTest(t, tp.prompt.Procedure, func(stdio terminal.Stdio) error {
				opts.AskOpts = prompt.WithStdio(stdio)
				return opts.Ask(tp.resource, tp.options)
			})
			if opts.Name != tp.want.Name {
				t.Errorf("Unexpected Task Name")
			}
			if opts.Catalog != tp.want.Catalog {
				t.Errorf("Unexpected Catalog Name")
			}
			if opts.Version != tp.want.Version {
				t.Errorf("Unexpected Version")
			}
		})
	}
}
