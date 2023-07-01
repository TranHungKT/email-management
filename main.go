package main

import (
	"github.com/TranHungKT/email_management/database"
	"github.com/TranHungKT/email_management/router"
)

func main() {
	database.InitDB()
	router.InitGin()
}
