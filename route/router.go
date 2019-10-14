package route

import (
	"errors"
	"github.com/Amakuchisan/QuestionBox/database"
	"github.com/Amakuchisan/QuestionBox/handler"
	"github.com/Amakuchisan/QuestionBox/model"
	"github.com/labstack/echo"
	"html/template"
	"io"
	"path/filepath"
)

// TemplateRegistry -- This have all templates
type TemplateRegistry struct {
	templates map[string]*template.Template
}

// Render -- Rendering templates
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "layout.html", data)
}

// Init : routerの初期化
func Init() *echo.Echo {
	e := echo.New()
	e.Debug = true

	templates := make(map[string]*template.Template)
	templates["index.html"] = makeTemplate("index.html")
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	e.GET("/", handler.MainPage)
	e.GET("/auth/login/:provider", handler.LoginHandler)
	e.GET("/auth/callback/:provider", handler.CallbackHandler)

	userHandler := handler.NewUserHandler(model.NewUserModel(database.DB))
	e.GET("/users", userHandler.UserAll)
	// e.POST("/users", userHandler.CreateUser)
	// e.GET("/users/:id", userHandler.DetailUser)
	// e.DELETE("/users/:id", userHandler.DeleteUser)
	return e
}

const (
	baseTemplate = "templates/layout.html"
	templatesDir = "templates"
)

func makeTemplate(html string) *template.Template {
	templateFile := filepath.Join(templatesDir, html)
	return template.Must(template.ParseFiles(baseTemplate, templateFile))
}
