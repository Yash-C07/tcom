package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"net/http"
	"bytes"
	"encoding/json"
)


type LoginModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var loginReq LoginModel
var isLoggedin bool = false

func main() {


	app := tview.NewApplication()
	app.EnableMouse(true)
	pages := tview.NewPages()

	var messages []string = make([]string, 0)
	// box := tview.NewBox().SetBorder(true).SetBackgroundColor(tcell.ColorGrey).SetTitle(" Home ")
	textbox := tview.NewInputField().SetPlaceholder("type something...")
	textbox.SetFieldBackgroundColor(tcell.ColorRed)
	chatview := tview.NewTextView()

	button := tview.NewButton("Click me!").SetSelectedFunc(func() {
		if textbox.GetText() == "" {
			return
		}
		messages = append(messages, textbox.GetText())
		fmt.Fprintln(chatview, ">>>", messages[len(messages)-1])
		textbox.SetText("")
	})
	messagebar := tview.NewFlex().AddItem(textbox, 0, 1, true).AddItem(button, 15, 1, true)

	chatview.SetBorder(true)
	textbox.SetDoneFunc(func(key tcell.Key) {
		if textbox.GetText() == "" {
			return
		}
		messages = append(messages, textbox.GetText())
		fmt.Fprintln(chatview, ">>>", messages[len(messages)-1])
		textbox.SetText("")
	})
	chatlayout := tview.NewFlex().AddItem(chatview, 0, 1, true).AddItem(messagebar, 1, 1, true).SetDirection(tview.FlexRow)

	pages.AddPage("chat", chatlayout, true, true)
	pages.AddPage("auth", GetAuthPage(pages), true, false)
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


func GetAuthPage(pages *tview.Pages) *tview.Flex {
	page := tview.NewFlex().SetDirection(tview.FlexRow)
	page.SetBorder(true).SetTitle(" TCOM Messaging ")

	headerText := tview.NewTextView().
		SetText("Choose authentication method!").
		SetTextAlign(tview.AlignCenter)

	headerText2 := tview.NewTextView().
		SetText("Login | Register").
		SetTextAlign(tview.AlignCenter)

	headerRow := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(headerText, 0, 1, false).
		AddItem(headerText2, 0, 1, true)


	loginModel := tview.NewForm().
		AddInputField("Username", "", 0, nil, func(text string) {
			loginReq.Username = text
		}).
		AddInputField("Password", "", 0, nil, func(text string) {
			loginReq.Password = text
		}).
		AddButton("Sign in", func(){

		 	readerData, er := json.Marshal(loginReq)
			reader := bytes.NewReader(readerData)

			if er != nil {
				fmt.Println(er)
			}

			response, err := http.Post("http://localhost:3000/login", "application/json", reader)
			defer response.Body.Close()
			if err != nil {
				fmt.Println(err)
			}

			if response.StatusCode == http.StatusOK {
				isLoggedin = true
				pages.SwitchToPage("chat")
			}
		})

	loginModel.SetBorder(true)

	registerModel := tview.NewForm().
		AddInputField("Username", "", 0, nil, func(text string) {
			loginReq.Username = text
		}).
		AddInputField("Password", "", 0, nil, func(text string) {
			loginReq.Password = text
		}).
		AddInputField("Confirm password", "", 0, nil, func(text string) {
			loginReq.Password = text
		}).

		AddButton("Sign up", func(){

			readerData, er := json.Marshal(loginReq)

			reader := bytes.NewReader(readerData)

			if er != nil {
				fmt.Println(er)
			}

			response, err := http.Post("http://localhost:3000/register", "application/json", reader)
			defer response.Body.Close()
			if err != nil {
				fmt.Println(err)
			}

			if response.StatusCode == http.StatusOK {
				isLoggedin = true
				pages.SwitchToPage("chat")
			} else {
				// SHOW ERROR MESSAGE
			}
		})		
	registerModel.SetBorder(true)


	authCol := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(loginModel, 0 , 1, true).
		AddItem(registerModel, 0, 1, true)

	page.
		AddItem(nil, 0, 1, false).       // Empty top spacer (pushes content down)
		AddItem(headerRow, 2, 2, false).  // Your actual content row (fixed height of 2 lines)
		AddItem(authCol, 0, 6, false)        // Empty bottom spacer (pushes content up)

	return page
}

