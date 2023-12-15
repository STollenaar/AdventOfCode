package golang

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"text/template"
)

//go:embed golang.tmpl
var golangTmpl string

type TemplateData struct {
}

func Create(dayName string, force bool) error {
	_, err := os.Getwd() // os.Getenv("GOPACKAGE") not available since this is not run with go generate
	if err != nil {
		return fmt.Errorf("error reading working directory: %s", err)
	}

	// servicePackage := filepath.Base(wd)

	if dayName == "" {
		return fmt.Errorf("error checking: no name given")
	}

	templateData := TemplateData{}

	tmpl := golangTmpl

	f := fmt.Sprintf("%s/main.go", dayName)
	if err = writeTemplate("newds", f, tmpl, force, templateData); err != nil {
		return fmt.Errorf("writing golang template: %w", err)
	}
	return nil
}

func writeTemplate(templateName, filename, tmpl string, force bool, td TemplateData) error {
	if _, err := os.Stat(filename); !errors.Is(err, fs.ErrNotExist) && !force {
		return fmt.Errorf("file (%s) already exists and force is not set", filename)
	}
	dir := strings.Join(strings.Split(filename, "/")[:2], "/")
	os.MkdirAll(dir, 0770)

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error opening file (%s): %s", filename, err)
	}

	tplate, err := template.New(templateName).Parse(tmpl)
	if err != nil {
		return fmt.Errorf("error parsing template: %s", err)
	}

	var buffer bytes.Buffer
	err = tplate.Execute(&buffer, td)
	if err != nil {
		return fmt.Errorf("error executing template: %s", err)
	}

	if _, err := f.Write(buffer.Bytes()); err != nil {
		f.Close() // ignore error; Write error takes precedence
		return fmt.Errorf("error writing to file (%s): %s", filename, err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("error closing file (%s): %s", filename, err)
	}
	os.Create(dir + "/input.txt")

	return nil
}
