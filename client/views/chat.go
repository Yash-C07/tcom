package views

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func GetChatPage() *tview.Flex {

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
	return chatlayout

}
