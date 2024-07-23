package main

import app "server/internal/chat"

func main() {
	chatAPI := app.NewApp()
	chatAPI.Run()
}
