package ui

import (
	//"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)



type SidepaneProps struct{
	//Context for the sidepane
	gtx   layout.Context;
	//The global theme
	theme material.Theme;
	//if any child widgets need to be rendered
	child layout.Widget;
}

func Sidepane(props SidepaneProps) layout.Dimensions {

	return  Border(BorderProps{
		gtx: props.gtx,
		border: widget.Border{
			Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
			CornerRadius: unit.Dp(0),
			Width:        unit.Dp(1),
		},
		child: func(gtx1 layout.Context) layout.Dimensions {
			return layout.Flex{
				Axis:      layout.Vertical,
				Alignment: layout.Middle,
				Spacing:   layout.SpaceEvenly,
			}.Layout(gtx1,
	
				layout.Rigid(func(_gtx layout.Context) layout.Dimensions {
	
					return material.H3(&props.theme, "Sidepane").Layout(_gtx)
	
					//return material.H1(globalContext.Theme,"Sidepane").Layout(*globalContext.gtx)
				}),
				layout.Rigid(func(_gtx layout.Context) layout.Dimensions {
					//RenderImage("./assets/gamer.png", _gtx.Ops)
	
					return material.H1(&props.theme, "Sidepane").Layout(_gtx)
				}),
			)
		},
	})

}

/*
* renders the side pane used for file navigation or ticket
 */
func RenderNavigationPane(globalContext Context,) {

	


	
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
