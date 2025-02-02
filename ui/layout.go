package ui

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Widget func(gtx layout.Context) layout.Dimensions

/*
* renders the side pane used for file navigation or ticket
 */
func RenderNavigationPane(globalContext Context) layout.Dimensions {
	//RenderImage("./assets/gamer.png",gtx.Ops)
	border := widget.Border{
		Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
		CornerRadius: unit.Dp(3),
		Width:        unit.Dp(1),
	}

	return layout.Flex{
		Axis:      layout.Vertical,
		Alignment: layout.Middle,
	}.Layout(*globalContext.gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return border.Layout(*globalContext.gtx,

				func(gtx layout.Context) layout.Dimensions {
					return material.H1(globalContext.Theme, "Sidepane").Layout(*globalContext.gtx)
				})
			//return material.H1(globalContext.Theme,"Sidepane").Layout(*globalContext.gtx)
		}),
	)

}

/*
Code pane is where all code related viewing is to be done
*/
func RenderCodePane(gtx layout.Context) {

}

/*
Description Pane is where you can add your description of code or
provide a markdown editor to showcase your breakdown of problem
*/
func RenderDescriptionPane(gtx layout.Context) {

}
