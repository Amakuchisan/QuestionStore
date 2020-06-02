package main

import (
	"github.com/Amakuchisan/QuestionStore/route"
)

func main() {
	route.Echo.Logger.Fatal(route.Echo.Start(":1323"))
}
