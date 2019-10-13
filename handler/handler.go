package handler

import (
	"database/sql"
	// MySQL Driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
	"log"
	"net/http"
	"time"
)

const (
	dbDriver = "mysql"

	// TODO: read from environment values
	dataSource = "tts:tts@tcp(mysql-container:3306)/tts?parseTime=true"
)

var db *sql.DB

// User -- This is user model
type User struct {
	ID        uint64
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func init() {
	var err error
	db, err = sql.Open(dbDriver, dataSource)

	if err != nil {
		log.Fatal("failed to connect db", err)
	}
}

// MainPage -- top page
func MainPage(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

// UsersPage -- user page
func UsersPage(c echo.Context) error {
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

	return c.JSON(http.StatusOK, users)
}

// LoginHandler -- Login to each provider
func LoginHandler(c echo.Context) error {
	provider, err := gomniauth.Provider(c.Param("provider"))
	if err != nil {
		return err
	}
	state := gomniauth.NewState("after", "success")
	authURL, err := provider.GetBeginAuthURL(state, nil)

	if err != nil {
		return err
	}
	return c.Redirect(http.StatusMovedPermanently, authURL)

}

// CallbackHandler -- Provider called this handler after login
func CallbackHandler(c echo.Context) error {
	provider, err := gomniauth.Provider(c.Param("provider"))
	if err != nil {
		return err
	}

	omap, err := objx.FromURLQuery(c.QueryString())
	if err != nil {
		return err
	}

	creds, err := provider.CompleteAuth(omap)
	if err != nil {
		return err
	}

	user, err := provider.GetUser(creds)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
