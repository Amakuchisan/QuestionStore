package main

import (
	//"fmt"
	"database/sql"
	"encoding/json"
	"github.com/Amakuchisan/QuestionBox/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	DB_DRIVER = "mysql"

	// TODO: read from environment values
	DATA_SOURCE = "tts:tts@tcp(mysql-container:3306)/tts?parseTime=true"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

var db *sql.DB

type User struct {
	ID        uint64
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func init() {
	var err error
	db, err = sql.Open(DB_DRIVER, DATA_SOURCE)

	if err != nil {
		log.Fatal("failed to connect db", err)
	}
}

func main() {
	e := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer
	e.GET("/", handler.MainPage).Name = "main"
	e.GET("/users", func(c echo.Context) error {
		rows, err := db.Query("SELECT * from users;")
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		users := []User{}

		for rows.Next() {
			user := User{}
			err := rows.Scan(&user.ID, &user.Name, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
			if err != nil {
				panic(err)
			}
			users = append(users, user)
		}
		bytes, err := json.Marshal(users)
		return c.JSON(http.StatusOK, string(bytes))
	})
	defer db.Close()
	e.Logger.Fatal(e.Start(":1323"))
}
