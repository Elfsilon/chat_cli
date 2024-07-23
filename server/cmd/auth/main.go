package main

import app "server/internal/auth"

func main() {
	authAPI := app.NewApp()
	authAPI.Run()
}
