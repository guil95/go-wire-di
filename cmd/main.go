package main

import "github.com/guil95/go-wire-di/cmd/di"

func main() {
	app := di.SetupApp()

	app.Run()
}
