package handler

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"net/http"
	"time"
	"log"
)

const (
	DB_DRIVER = "mysql"

	// TODO: read from environment values
	DATA_SOURCE = "tts:tts@tcp(mysql-container:3306)/tts?parseTime=true"
)

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

func MainPage(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

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

func CallbackHandler(c echo.Context) error {
	provider, err := gomniauth.Provider(c.Param("provider"))
	if err != nil {
		return err
	}

	omap, err := objx.FromURLQuery(c.QueryString())
	creds, err := provider.CompleteAuth(omap)

	if err != nil {
		return err
	}

	user, err := provider.GetUser(creds)
	if err != nil {
		return err
	}
	// Debug
	// 	fmt.Printf("%v", user)
	// 	fmt.Printf("%s, %s, %s", user.Nickname(), user.Email(), user.AvatarURL())
	log.Println(user.Email())
	return c.JSON(http.StatusOK, user)
}
