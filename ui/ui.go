package ui

import (
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/widget"
)

func runApp(){

    a :=app.New()
    w :=a.NewWindow("Hello World")

    w.SetContent(widget.NewLabel("Hello World"))

    w.ShowAndRun()
}


