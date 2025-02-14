package ui

import (
	"image/color"
	"gioui.org/text"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"fmt"
	"gioui.org/widget"          // widget contains state for different widgets

)

var (
	red = color.NRGBA{R:0xC0,G:0x40,B:0x40,A:0xFF}
	titleColor = color.NRGBA{R: 127, G: 0, B: 0, A: 255}
)
// Title creates a center aligned H1.
func title(th *material.Theme, caption string) material.LabelStyle {
	label := material.H1(th, caption)
	label.Color = titleColor
	label.Alignment = text.Middle
	return label
}

type Context struct{
	Theme *material.Theme
	gtx *layout.Context
}

var context = Context{
	Theme:material.NewTheme(),
}

var editor widget.Editor

var inset = layout.UniformInset(8)
var border = widget.Border{
	Color: color.NRGBA{R: 0x88, G: 0x88, B: 0x88, A: 0xFF},
	Width: 1,
}

func Center(label material.LabelStyle) material.LabelStyle {
	label.Alignment = text.Middle
	return label
}
/*
	Logic for the rendering is written here
*/
func Layout(gtx *layout.Context) layout.Dimensions{
	context.gtx=gtx

	return inset.Layout(*gtx,
		func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{
				Axis:      layout.Horizontal,
				Alignment: layout.Middle,
			}.Layout(gtx,
				//layout.Rigid(RenderNavigationPane(context)),
				layout.Flexed(1,func(gtx layout.Context) layout.Dimensions {
					return RenderNavigationPane(context)
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return inset.Layout(gtx, material.Editor(context.Theme, &editor, "").Layout)
					})
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					line, col := editor.CaretPos()
					s := fmt.Sprintf("line:%d col:%d", line, col)
					return Center(material.Body1(context.Theme, s)).Layout(gtx)
				}),
			)
		})

	
}



