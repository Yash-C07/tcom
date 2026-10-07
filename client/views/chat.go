package views

import (
	"fmt"
	"tcom/api"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func GetChatPage() *tview.Flex {

	var messages []string = make([]string, 0)
	page := tview.NewFlex().SetDirection(tview.FlexColumn)
	page.SetBorder(true).SetTitle("TCOM")
	// box := tview.NewBox().SetBorder(true).SetBackgroundColor(tcell.ColorGrey).SetTitle(" Home ")
	textbox := tview.NewInputField().SetPlaceholder("type something...")
	textbox.SetFieldBackgroundColor(tcell.ColorBlack)

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
	chatLayout := tview.NewFlex().
		AddItem(chatview, 0, 1, true).
		AddItem(messagebar, 1, 1, true).
		SetDirection(tview.FlexRow)
	friendListview := tview.NewList()
	//	AddItem("Yashwanth", "yash16", 'Y', func() { chatview.SetTitle(" yash16 ") }).
	friendlist := api.GetFriends()
	for u, d := range friendlist {
		friendListview.AddItem(d, u, 'o', func() { chatview.SetTitle(d) })
	}
	friendListview.SetBorder(true).SetTitle(" Friend List ")
	requestlistview := tview.NewList()
	requestlist := api.GetFriendReqs()
	for _, d := range requestlist {
		requestlistview.AddItem(d, "", 'o', func() { chatview.SetTitle(d) })
	}

	requestlistview.SetBorder(true).SetTitle(" Requests ")

	friendCol := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(friendListview, 0, 1, true).
		AddItem(requestlistview, 0, 1, true)

	page.
		AddItem(friendCol, 0, 2, true).
		AddItem(chatLayout, 0, 8, true)
	return page

}
