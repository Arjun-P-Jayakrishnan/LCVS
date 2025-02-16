package ui

import (
	"gioui.org/layout"
	"gioui.org/widget"
)

type BorderProps struct {
	gtx    layout.Context
	border widget.Border
	child  layout.Widget
}

func Border(props BorderProps) layout.Dimensions{
	return props.border.Layout(props.gtx,props.child);
}