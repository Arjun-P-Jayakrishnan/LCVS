package elements

import (
	"gioui.org/layout"
	"gioui.org/widget"
)

type BorderProps struct {
	//the graphical context
	Gtx    layout.Context
	//border props
	Border widget.Border
	//child widget
	Child  layout.Widget
}

func Border(props BorderProps) layout.Dimensions{
	return props.Border.Layout(props.Gtx,props.Child);
}