package internal

type EventHandler struct {
	State *AppState
}

func NewEventHandler(appState *AppState) *EventHandler {
	return &EventHandler{State: appState}
}

func (h *EventHandler) SwicthView(view ViewType) {
	h.State.ActiveView = view
}
