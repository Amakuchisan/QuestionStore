package main

import (
	"errors"
	"github.com/Amakuchisan/QuestionBox/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"path/filepath"
)

type TemplateRegistry struct {
	templates map[string]*template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "layout.html", data)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/static", "static")

	templates := make(map[string]*template.Template)
	templates["index.html"] = makeTemplate("index.html")
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}
	e.GET("/", handler.MainPage)
	e.GET("/users", handler.UsersPage)
	e.Logger.Fatal(e.Start(":1323"))
}

const (
	baseTemplate = "templates/layout.html"
	templatesDir = "templates"
)

func makeTemplate(html string) *template.Template {
	templateFile := filepath.Join(templatesDir, html)
	return template.Must(template.ParseFiles(baseTemplate, templateFile))
}
