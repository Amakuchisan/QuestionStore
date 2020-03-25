package route

import (
	"html/template"
	"io"
	"net/http"

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

	g := e.Group("", authCheckMiddleware())

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	g.GET("/", handler.MainPage)
	g.GET("/questions/form", handler.QuestionFormHandler)
	e.GET("/auth/login/:provider", handler.LoginHandler)

	userHandler := handler.NewUserHandler(repository.NewUserModel(database.DB))
	e.GET("/auth/callback/:provider", userHandler.CallbackHandler)
	g.GET("/users", userHandler.UserAll)
	// e.GET("/users/:id", userHandler.DetailUser)
	// e.DELETE("/users/:id", userHandler.DeleteUser)
	questionHandler := handler.NewQuestionHandler(repository.NewQuestionModel(database.DB))
	g.POST("/questions", questionHandler.PostQuestion)
	g.GET("/questions", questionHandler.QuestionsTitleList)
	g.GET("/questions/:id", questionHandler.QuestionDetail)

	return e
}

func authCheckMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			_, err := c.Cookie("auth")
			if err != nil {
				return c.Redirect(http.StatusTemporaryRedirect, "/auth/login/google")
			}

			return next(c)
		}
	}
}
