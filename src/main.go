package main

import (
	"github.com/ygorrodriguesmeli/gorestapi/src/database"
	"github.com/ygorrodriguesmeli/gorestapi/src/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()
	server.Run()
}
