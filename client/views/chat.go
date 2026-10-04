package views

import (
	"fmt"

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
	friendList := tview.NewList().
		AddItem("Yashwanth", "yash16", 'Y', func() { chatview.SetTitle(" yash16 ") }).
		AddItem("Kavin Charles", "kavincharles", 'K', func() { chatview.SetTitle(" kavincharles ") })
	friendList.SetBorder(true).SetTitle(" Friend List ")
	page.
		AddItem(friendList, 0, 2, true).
		AddItem(chatLayout, 0, 8, true)
	return page

}
