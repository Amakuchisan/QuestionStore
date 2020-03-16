package route

import (
	"html/template"
	"io"

	"github.com/Amakuchisan/QuestionStore/database"
	"github.com/Amakuchisan/QuestionStore/handler"
	"github.com/Amakuchisan/QuestionStore/repository"
	"github.com/labstack/echo"
)

// TemplateRenderer -- custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render -- Rendering templates
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Init : routerの初期化
func Init() *echo.Echo {
	e := echo.New()
	e.Debug = true

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	e.GET("/", handler.MainPage)
	e.GET("/questions/form", handler.QuestionFormHandler)
	e.GET("/auth/login/:provider", handler.LoginHandler)

	userHandler := handler.NewUserHandler(repository.NewUserModel(database.DB))
	e.GET("/auth/callback/:provider", userHandler.CallbackHandler)
	e.GET("/users", userHandler.UserAll)
	// e.GET("/users/:id", userHandler.DetailUser)
	// e.DELETE("/users/:id", userHandler.DeleteUser)
	questionHandler := handler.NewQuestionHandler(repository.NewQuestionModel(database.DB))
	e.POST("/questions", questionHandler.PostQuestion)
	e.GET("/questions", questionHandler.QuestionsTitleList)
	e.GET("/questions/:id", questionHandler.Question)

	return e
}
