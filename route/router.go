package route

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Amakuchisan/QuestionStore/database"
	"github.com/Amakuchisan/QuestionStore/handler"
	"github.com/Amakuchisan/QuestionStore/repository"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
)

// TemplateRenderer -- custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render -- Rendering templates
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Echo -- instance for the echo web framework
var Echo *echo.Echo

func init() {
	e := echo.New()
	e.Debug = true

	err := setupOAuth()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

  g := e.Group("", authCheckMiddleware())

	e.Use(middleware.Logger())
	e.Static("/static", "static")

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

	Echo = e
}

func setupOAuth() error {
	if os.Getenv("QS_ENV") == "development" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}

	host := os.Getenv("QS_HOST")
	googleCallbackURL := fmt.Sprintf("http://%s/auth/callback/google", host)

	// setup gomniauth
	gomniauth.SetSecurityKey(os.Getenv("SECURITY_KEY"))
	gomniauth.WithProviders(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_CLIENT_SECRET"),
			googleCallbackURL,
		),
	)
	return nil
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
