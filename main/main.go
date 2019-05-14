package main

import (
	"attendance/router"
)

func main() {
	//defer database.SqlDB.Close()
	router.InitRouter(":9999")
}