package main

import (
	"fmt"
	"tcom/models"
	"tcom/views"

	"github.com/rivo/tview"
)

var loginReq models.LoginModel
var isLoggedin bool = false

func main() {

	app := tview.NewApplication()
	app.EnableMouse(true)
	pages := tview.NewPages()

	pages.AddPage("auth", views.GetAuthPage(pages), true, false)
	if !isLoggedin {
		pages.SwitchToPage("auth")
	}

	app.SetRoot(pages, true)
	err := app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
