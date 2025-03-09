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
	btn :=material.Button(&props.Theme,new(widget.Clickable),props.Label)
	props.Gtx.Constraints.Max.X*=6;
	return btn.Layout(props.Gtx)
}

