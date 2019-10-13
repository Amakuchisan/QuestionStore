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
