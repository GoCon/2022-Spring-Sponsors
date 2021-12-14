package main

import (
	"embed"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/gostaticanalysis/skeletonkit"
)

var (
	//go:embed _template/*
	tmplFS embed.FS
	tmpl   = template.Must(skeletonkit.ParseTemplate(tmplFS, "makecompanies", "_template"))
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	data, err := readCompaniesAll()
	if err != nil {
		return err
	}

	fsys, err := skeletonkit.ExecuteTemplate(tmpl, data)
	if err != nil {
		return err
	}

	prompt := &skeletonkit.Prompt{
		Output:    io.Discard,
		ErrOutput: os.Stderr,
		Input:     strings.NewReader("3\ny\ny\n"),
	}

	if err := skeletonkit.CreateDir(prompt, ".", fsys); err != nil {
		return err
	}

	return nil
}
