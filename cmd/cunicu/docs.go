package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/stv0g/cunicu/pkg/util/buildinfo"
	"gopkg.in/yaml.v3"
)

var (
	outputDir       string
	withFrontMatter bool

	docsCmd = &cobra.Command{
		Use:    "docs",
		Short:  "Generate documentation for the cunīcu commands",
		Long:   `When used without a sub-command, both the Markdown documentation and Man-pages will be generated.`,
		Hidden: true,
		Args:   cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := docsMarkdown(cmd, args); err != nil {
				return err
			}

			return docsManpage(cmd, args)
		},
	}

	docsMarkdownCmd = &cobra.Command{
		Use:   "markdown",
		Short: "Generate markdown docs",
		RunE:  docsMarkdown,
		Args:  cobra.NoArgs,
	}

	docsManpageCmd = &cobra.Command{
		Use:   "man",
		Short: "Generate manpages",
		RunE:  docsManpage,
		Args:  cobra.NoArgs,
	}
)

func init() {
	rootCmd.AddCommand(docsCmd)

	docsCmd.AddCommand(docsManpageCmd)
	docsCmd.AddCommand(docsMarkdownCmd)

	pf := docsCmd.PersistentFlags()
	pf.StringVar(&outputDir, "output-dir", "./docs/usage", "Output directory of generated documentation")
	pf.BoolVar(&withFrontMatter, "with-frontmatter", false, "Prepend a frontmatter to the generated Markdown files as used by our static website generator")
}

func docsMarkdown(cmd *cobra.Command, args []string) error {
	dir := filepath.Join(outputDir, "md")

	//#nosec G301 -- Doc directories must be world readable
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	filePrepender := func(path string) string {
		if !withFrontMatter {
			return ""
		}

		filename := filepath.Base(path)
		filename = strings.TrimSuffix(filename, ".md")
		parts := strings.Split(filename, "_")

		fm := struct {
			Title            string   `yaml:"title"`
			SidebarLabel     string   `yaml:"sidebar_label,omitempty"`
			SidebarClassName string   `yaml:"sidebar_class_name"`
			Slug             string   `yaml:"slug"`
			HideTitle        bool     `yaml:"hide_title"`
			Keywords         []string `yaml:"keywords"`
		}{
			Title:            strings.Join(parts, " "),
			SidebarClassName: "command-name",
			Slug:             fmt.Sprintf("/usage/man/%s", strings.Join(parts[1:], "/")),
			HideTitle:        true,
			Keywords:         []string{"manpage"},
		}

		if len(parts) > 1 {
			fm.SidebarLabel = strings.Join(parts[1:], " ")
		}

		fmYaml, err := yaml.Marshal(fm)
		if err != nil {
			return ""
		}

		return fmt.Sprintf("---\n%s---\n\n", fmYaml)
	}

	// The linkHandler can be used to customize the rendered internal links to the commands, given a filename:
	linkHandler := func(name string) string {
		return name
	}

	return doc.GenMarkdownTreeCustom(rootCmd, dir, filePrepender, linkHandler)
}

func docsManpage(cmd *cobra.Command, args []string) error {
	dir := filepath.Join(outputDir, "man")

	//#nosec G301 -- Doc directories must be world readable
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	header := &doc.GenManHeader{
		Title:  "cunīcu",
		Source: "https://github.com/stv0g/cunicu",
		Date:   buildinfo.Date,
	}

	return doc.GenManTree(rootCmd, header, dir)
}