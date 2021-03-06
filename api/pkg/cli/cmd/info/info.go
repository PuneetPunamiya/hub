// Copyright © 2020 The Tekton Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package info

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/tektoncd/hub/api/pkg/cli/app"
	"github.com/tektoncd/hub/api/pkg/cli/flag"
	"github.com/tektoncd/hub/api/pkg/cli/formatter"
	"github.com/tektoncd/hub/api/pkg/cli/hub"
	"github.com/tektoncd/hub/api/pkg/cli/printer"
	so "github.com/tektoncd/hub/api/pkg/cli/select_options"
)

var cmdExamples string = `
Display info of a %S of name 'foo':

    tkn hub info %s foo

or

Display info of a %S of name 'foo' of version '0.3':

    tkn hub info %s foo --version 0.3
`

const resTemplate = `{{ icon "name" }}Name: {{ .Resource.Name }}
{{ $n := .ResVersion.DisplayName }}{{ if ne (default $n "") "" }}
{{ icon "displayName" }}Display Name: {{ $n }}
{{ end }}
{{ icon "version" }}Version: {{ formatVersion .ResVersion.Version .Latest }}

{{ icon "description" }}Description: {{ formatDesc .ResVersion.Description 80 16 }}

{{ icon "minPipelineVersion" }}Minimum Pipeline Version: {{ .ResVersion.MinPipelinesVersion }}

{{ icon "rating" }}Rating: {{ .Resource.Rating }}

{{ $t := len .Resource.Tags }}{{ if ne $t 0 }}
{{- icon "tags" }}Tags
 {{- range $p := .Resource.Tags }}
  {{ icon "bullet" }}{{ $p.Name }}
 {{- end }}
{{- end }}

{{ icon "install" }}Install Command:
  {{ formatInstallCMD .Resource .ResVersion .Latest }}
`

var (
	funcMap = template.FuncMap{
		"icon":             formatter.Icon,
		"formatDesc":       formatter.WrapText,
		"formatVersion":    formatter.FormatVersion,
		"formatInstallCMD": formatter.FormatInstallCMD,
		"default":          formatter.DefaultValue,
		"lower":            strings.ToLower,
	}
	tmpl = template.Must(template.New("Resource Info").Funcs(funcMap).Parse(resTemplate))
)

type templateData struct {
	Resource   *hub.ResourceData
	ResVersion *hub.ResourceWithVersionData
	Latest     bool
}

type options struct {
	cli          app.CLI
	from         string
	version      string
	kind         string
	args         []string
	hubClient    hub.Client
	selectOption so.Options
}

func Command(cli app.CLI) *cobra.Command {

	opts := &options{cli: cli}

	cmd := &cobra.Command{
		Use:   "info",
		Short: "Display info of resource by its name, kind, catalog, and version",
		Long:  ``,
		Annotations: map[string]string{
			"commandType": "main",
		},
		SilenceUsage: true,
	}
	cmd.AddCommand(
		commandForKind("task", opts),
	)

	cmd.PersistentFlags().StringVar(&opts.from, "from", "", "Name of Catalog to which resource belongs.")
	cmd.PersistentFlags().StringVar(&opts.version, "version", "", "Version of Resource")

	return cmd
}

// commandForKind creates a cobra.Command that when run sets
// opts.Kind and opts.Args and invokes opts.run
func commandForKind(kind string, opts *options) *cobra.Command {

	return &cobra.Command{
		Use:          kind,
		Short:        "Display info of " + strings.Title(kind) + " by its name, catalog and version",
		Long:         ``,
		SilenceUsage: true,
		Example:      examples(kind),
		Annotations: map[string]string{
			"commandType": "main",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.kind = kind
			opts.args = args
			return opts.run()
		},
	}
}

func (opts *options) run() error {

	if err := opts.validate(); err != nil {
		return err
	}

	opts.hubClient = opts.cli.Hub()
	var err error

	// Check if catalog name is passed else
	// ask the user to select the catalog
	if opts.from == "" {
		opts.from, err = opts.AskCatalogName()
		if err != nil {
			return err
		}
	}

	name := opts.name()

	// Check if resource name is passed else
	// ask the user to select the resource
	if name == "" {
		name, err = opts.AskResourceName()
		if err != nil {
			return err
		}
	}

	// Check if version of the resource is passed else
	// ask the user to select the version of resource
	if opts.version == "" {
		opts.version, err = opts.AskVersion(name)
		if err != nil {
			return err
		}
	}

	res := opts.hubClient.GetResource(hub.ResourceOption{
		Name:    name,
		Catalog: opts.from,
		Kind:    opts.kind,
		Version: opts.version,
	})

	resource, err := res.Resource()
	if err != nil {
		return err
	}

	out := opts.cli.Stream().Out

	resVersion := resource.(hub.ResourceWithVersionData)
	tmplData := templateData{
		ResVersion: &resVersion,
		Resource:   resVersion.Resource,
		Latest:     false,
	}
	return printer.New(out).Tabbed(tmpl, tmplData)
}

func (opts *options) AskCatalogName() (string, error) {
	// Get all Catalogs
	catalog, err := opts.hubClient.GetCatalogsList()
	if err != nil {
		return "", err
	}
	switch len(catalog) {
	case 0:
		return "", fmt.Errorf("No catalogs found")
	case 1:
		return catalog[0], nil
	default:
		// Ask the catalog
		err = opts.selectOption.Ask("catalog", catalog)
		if err != nil {
			return "", err
		}
		return opts.selectOption.Catalog, nil
	}

}

func (opts *options) AskResourceName() (string, error) {
	// Get all resources from the Catalog selected
	resources, err := opts.hubClient.GetResourcesList(hub.SearchOption{
		Kinds:   []string{opts.kind},
		Catalog: opts.from,
	})
	if err != nil {
		return "", err
	}
	switch len(resources) {
	case 0:
		return "", fmt.Errorf("No resources found")
	case 1:
		return resources[0], nil
	default:
		// Ask the resource name
		err = opts.selectOption.Ask(opts.kind, resources)
		if err != nil {
			return "", err
		}
		return opts.selectOption.Name, nil
	}
}

func (opts *options) AskVersion(name string) (string, error) {
	// Get all the versions of the resource selected
	ver, err := opts.hubClient.GetResourceVersionslist(hub.ResourceOption{
		Name:    name,
		Kind:    opts.kind,
		Catalog: opts.from,
	})
	if err != nil {
		return "", err
	}
	switch len(ver) {
	case 1:
		return ver[0], nil
	default:
		latestVersion := ver[0]
		ver[0] = ver[0] + " (latest)"

		// Ask the version
		err = opts.selectOption.Ask("version", ver)
		if err != nil {
			return "", err
		}

		if strings.Contains(opts.selectOption.Version, "(latest)") {
			opts.selectOption.Version = latestVersion
		}
		return opts.selectOption.Version, nil
	}
}

func (opts *options) validate() error {
	return flag.ValidateVersion(opts.version)
}

func (opts *options) name() string {
	if len(opts.args) == 0 {
		return ""
	}
	return strings.TrimSpace(opts.args[0])
}

func examples(kind string) string {
	replacer := strings.NewReplacer("%s", kind, "%S", strings.Title(kind))
	return replacer.Replace(cmdExamples)
}
