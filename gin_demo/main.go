package main

import (
	db "lovehome/gin_demo/database"
)

func main() {

	defer db.SqlDB.Close()

	router := initRouter()

	router.Run(":8000")
}
