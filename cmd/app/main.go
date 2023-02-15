package main

import (
	server "project/internal/server"
	store "project/internal/store"
)

// Точка входа
func main() {
	mainDB := store.ConnectDB()
	mainServer := server.NewServer(store.NewStore(mainDB))
	mainServer.Start()
}
