package views

import (
	"tcom/api"
	"tcom/types"

	"net/http"

	"github.com/rivo/tview"
)

var loginReq types.LoginModel
var registerReq types.RegisterModel

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
		AddButton("Sign in", func() {

			statuscode := api.Login(loginReq)
			if statuscode == http.StatusOK {

				pages.SwitchToPage("chat")
			}
		})

	loginModel.SetBorder(true)

	registerModel := tview.NewForm().
		AddInputField("Username", "", 0, nil, func(text string) {
			registerReq.Username = text
		}).
		AddInputField("Password", "", 0, nil, func(text string) {
			registerReq.Password = text
		}).
		AddInputField("Confirm password", "", 0, nil, func(text string) {
			registerReq.ConfirmPassword = text
		}).
		AddButton("Sign up", func() {
			var statuscode int
			if registerReq.ConfirmPassword == registerReq.Password {
				statuscode = api.Register(registerReq)
			} else {
				statuscode = http.StatusBadRequest
			}

			if statuscode == http.StatusOK {
				pages.SwitchToPage("chat")
			} else {
				// SHOW ERROR MESSAGE
			}
		})
	registerModel.SetBorder(true)

	authCol := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(loginModel, 0, 1, true).
		AddItem(registerModel, 0, 1, true)

	page.
		AddItem(nil, 0, 1, false).       // Empty top spacer (pushes content down)
		AddItem(headerRow, 2, 2, false). // Your actual content row (fixed height of 2 lines)
		AddItem(authCol, 0, 6, false)    // Empty bottom spacer (pushes content up)

	return page
}
