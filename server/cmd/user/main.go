package main

import app "server/internal/user"

func main() {
	userAPI := app.NewApp()
	userAPI.Run()
}
