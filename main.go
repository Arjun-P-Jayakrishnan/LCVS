package main

import (
	"github.com/Arjun-P-Jayakrishnan/LCVS/data"
	"github.com/Arjun-P-Jayakrishnan/LCVS/internal"
	"github.com/Arjun-P-Jayakrishnan/LCVS/ui"
)

func main() {
	//UI state
	appState := internal.NewAppState()

	//Event Handler
	eventHandler := &internal.EventHandler{
		State: appState,
	}

	//Router
	appRouter := internal.NewRouter(
		appState,
		eventHandler,
		ui.Layout,
	)

	appState.FileStorage=data.NewFilesDB()

	app := internal.NewApp(appRouter.Layout)

	app.Run()

	data.CloseConnection(appState.FileStorage)
}
