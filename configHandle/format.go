package confighandle

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log/slog"
)

// RenderTemplate Renders the template and returns the result as a string
func RenderTemplate(tmpl *template.Template, data interface{}) (string, error) {
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	if err != nil {
		slog.Error("Could not format template", "name", tmpl.Name(), "error", err)
		return "", fmt.Errorf("Unable to format template '%s'", tmpl.Name())
	}
	content, err := io.ReadAll(&buf)
	if err != nil {
		slog.Error("Could read from buffer", "error", err)
		return "", fmt.Errorf("Unable to format template '%s'", tmpl.Name())
	}
	return string(content), nil
}

// PrintItem Prints an item that can be formatted to the CLI
func PrintItem(i PrintableItem) {
	fmt.Print(i.FormatString())
}

func (m *Metadata) FormatString() string {
	templ, err := template.New("metadata").Parse(metaTemplate)
	if err != nil {
		panic(err)
	}
	rendered, err := RenderTemplate(templ, m)
	if err != nil {
		panic(err)
	}
	return rendered
}

func (v *Variable) FormatString() string {
	templ, err := template.New("variable").Parse(varTemplate)
	if err != nil {
		panic(err)
	}
	rendered, err := RenderTemplate(templ, v)
	if err != nil {
		panic(err)
	}
	return rendered
}

func (p *Project) FormatString() string {
	templ, err := template.New("project").Parse(projectTemplate)
	if err != nil {
		panic(err)
	}
	rendered, err := RenderTemplate(templ, p)
	if err != nil {
		panic(err)
	}
	return rendered
}
