package main

import (
	"fmt"
	"tcom/types"
	"tcom/views"

	"github.com/rivo/tview"
)

var loginReq types.LoginModel
var isLoggedin bool = false

func main() {

	app := tview.NewApplication()
	app.EnableMouse(true)
	pages := tview.NewPages()
	pages.AddPage("chat", views.GetChatPage(), true, true)
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

//checking my git
//AGAIN
