package helpers

import (
	"errors"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type TemplateRegistry struct {
	Templates map[string]*template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.Templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}
