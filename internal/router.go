package internal

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type OrganiseFunc func(gtx layout.Context, theme *material.Theme, state *AppState) layout.Dimensions

type Router struct {
	State    *AppState
	Handler  *EventHandler
	layoutFn OrganiseFunc
}

// Creates new Router
func NewRouter(s *AppState, h *EventHandler, fn OrganiseFunc) *Router {
	return &Router{
		State:    s,
		Handler:  h,
		layoutFn: fn,
	}
}

// Layout renders the current active view
func (r *Router) Layout(gtx layout.Context) layout.Dimensions {

	theme := r.State.UI.Theme

	return r.layoutFn(gtx, theme, r.State)

	switch r.State.ActiveView {
	case ViewFileTree:
		return material.Body1(theme, "File Tree View").Layout(gtx)
	case ViewEditor:
		return material.Body1(theme, "View Editor").Layout(gtx)
	case ViewDiagram:
		return material.Body1(theme, "View Diagram").Layout(gtx)
	case ViewJournal:
		return material.Body1(theme, "View Journal").Layout(gtx)
	default:
		return material.Body1(theme, "Unkown View").Layout(gtx)
	}
}
