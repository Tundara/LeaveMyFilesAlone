package gui

import (
	"LetMyFiles/windows"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func StartGui() {

	a := app.New()
	w := a.NewWindow("LeaveMyFiles")
	w.Resize(fyne.NewSize(800, 300))

	TxtFile := widget.NewLabel("File Choosen : ")
	TxtFile.Wrapping = fyne.TextWrapWord
	TxtFile.Alignment = fyne.TextAlignCenter

	Choosen := widget.NewLabel("")
	Choosen.Wrapping = fyne.TextWrapWord
	Choosen.Alignment = fyne.TextAlignCenter
	Choosen.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	button := widget.NewButton("Choose File", func() {
		info, file := windows.WinStart()
		Choosen.SetText(file)
		go windows.Process(info, file)
	})

	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), button, layout.NewSpacer())
	text := container.New(layout.NewVBoxLayout(), TxtFile, Choosen, centered, widget.NewSeparator())

	w.SetContent(centered)
	w.SetContent(text)

	w.ShowAndRun()
}
