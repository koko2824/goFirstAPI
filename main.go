package main

import (
	"test_members2_rest/db"
	"test_members2_rest/server"
)

func main() {
	db.Init()
	server.Init()
}