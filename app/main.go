package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func basic() string {
	return "Hello World!"
}

func main() {
	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:     985,
		Height:    615,
		Title:     "MockGopher",
		JS:        js,
		CSS:       css,
		Colour:    "#000",
		Resizable: true,
	})
	app.Bind(basic)
	app.Run()
}
