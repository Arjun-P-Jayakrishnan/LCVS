package internal

import (
	"fmt"
	"log"

	"gioui.org/widget/material"
	"github.com/Arjun-P-Jayakrishnan/LCVS/data"
)

// App state holds the global state of the application
type AppState struct {
	ActiveView  ViewType
	FileManager *FileManagerState
	Editor      *EditorState
	Diagram     *DiagramState
	Journal     *VCState
	UI          *UIState
	Layout      *LayoutState
	FileStorage *data.FileDB
}

// child structs
type FileManagerState struct {
	CurrentDirectory string
	SelectedFile     string
	Nodes            []*data.FileNode
}

type EditorState struct {
	OpenFilePath string
	Content      string
	IsDirty      bool
	FileType     string //.md .go etc
}

type DiagramState struct {
	Markdown    string
	DiagramData string
}

type VCState struct {
	Commits        []Commit
	SelectedCommit *Commit
}

type Commit struct {
	ID      string
	Message string
	Time    string
}

type UIState struct {
	Theme         *material.Theme
	ModalOpen     bool
	FocusedPane   string
	PanelsVisible map[string]bool
}

type LayoutState struct {
	LeftPane   []View
	CenterPane []View
	RightPane  []View
}

type View struct {
	ID       string
	Type     ViewType
	StateRef any
}

type ViewType int

const (
	ViewFileTree ViewType = iota
	ViewEditor
	ViewDiagram
	ViewJournal
)

func NewAppState() *AppState {
	fileNodes, err := data.BuildFileTree("D://")

	if err != nil {
		fmt.Print("File structure error")
		log.Fatal("File structure creation error")
	}

	fmt.Print(len(fileNodes))

	return &AppState{
		UI: &UIState{
			PanelsVisible: make(map[string]bool),
			Theme:         material.NewTheme(),
		},
		FileManager: &FileManagerState{
			Nodes: fileNodes,
		},
		ActiveView:  ViewDiagram,
		FileStorage: &data.FileDB{},
	}
}
