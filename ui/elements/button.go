package elements

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type ButtonProps struct{
	Theme material.Theme
	Gtx layout.Context
	Label string
}

func Button(props ButtonProps) layout.Dimensions{
	btn :=material.Button(&props.Theme,&widget.Clickable{},props.Label)

	return btn.Layout(props.Gtx)
}