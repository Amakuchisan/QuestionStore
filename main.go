package main

import (
	"errors"
	"github.com/Amakuchisan/QuestionBox/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
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

	templates := make(map[string]*template.Template)
	templates["index.html"] = template.Must(template.ParseFiles("templates/index.html", "templates/layout.html"))
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}
	e.GET("/", handler.MainPage).Name = "main"
	e.GET("/users", handler.UsersPage)
	e.Logger.Fatal(e.Start(":1323"))
}
