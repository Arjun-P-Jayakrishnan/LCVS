package ui

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/Arjun-P-Jayakrishnan/LCVS/internal"
)

func Layout(gtx layout.Context,theme *material.Theme,state *internal.AppState) layout.Dimensions{

	

	return layout.Flex{Axis:layout.Horizontal}.Layout(gtx,
	

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			width := gtx.Constraints.Max.X/5
			gtx.Constraints.Min.X=width
			gtx.Constraints.Max.X=width

			return InitFileNav(gtx,theme,state.FileManager)
		}),
		
		layout.Flexed(0.5,func(gtx layout.Context) layout.Dimensions {
			return Editor(gtx,theme)
		}),

		layout.Flexed(0.3,func(gtx layout.Context) layout.Dimensions {
			return Diagram(gtx,theme)
		}),
	)

}