package ui

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

func Diagram(gtx layout.Context, th *material.Theme) layout.Dimensions{
	return layout.UniformInset(10).Layout(gtx,func(gtx layout.Context) layout.Dimensions {
		return material.Body1(th,"Diagram").Layout(gtx)
	})
}