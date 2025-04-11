package ui

import (
	"log"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/Arjun-P-Jayakrishnan/LCVS/data"
	"github.com/Arjun-P-Jayakrishnan/LCVS/internal"
)

// Type coersion
type FileNode = data.FileNode

// List of clickable tiles
var list = widget.List{List: layout.List{Axis: layout.Vertical}}

// InitFileNav renders the file tree structure from internal state.
//
//	`gtx`	Graphical Context
//	`th`	Theme
//	`state`	FileManagerState
func InitFileNav(
	gtx layout.Context, 		
	th *material.Theme,			
	state *internal.FileManagerState,
) layout.Dimensions {

	//have a gap at sides
	return layout.Inset{
		Top:   10,
		Left:  10,
		Right: 10,
	}.Layout(
		gtx,
		func(gtx layout.Context) layout.Dimensions {
			return material.List(th, &list).Layout(
				gtx,
				len(state.Nodes),
				func(gtx layout.Context, index int) layout.Dimensions {
					//base level folders and files
					return renderNode(gtx, th, state.Nodes[index], 0, "./", state)
				})
		})

}

// renderNode represents an internal function that creates the clickable list tiles.
//
//	`gtx`		Graphical Context
//	`th`		Theme
//	`filenode`	Reference to the File Meta Data
//	`depth`		Depth is the level from root, root having depth of 0
//	`path`		Path to the file or folder
//	`state`		FileManagerState reference
func renderNode(
	gtx layout.Context, 	// 	
	th *material.Theme, 	//	theme
	fileNode *FileNode,		//	node that contains details of the current file
	depth int,				//	from root having depth of 0 what is the depth
	path string,			//	path to file relative or absolute
	state *internal.FileManagerState, // app state reference
) layout.Dimensions {

	inset := layout.Inset{Left: unit.Dp(float32(depth) * 16)}

	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				//Button Click check
				if fileNode.Button.Clicked(gtx) {
					//If Directory
					if fileNode.IsDir {
						//Toggle state
						fileNode.Expanded = !fileNode.Expanded

						//Lazy Load files (imporve later)
						if fileNode.Expanded && fileNode.Children == nil {
							children, err := data.BuildFileTree(fileNode.Path)

							if err != nil {
								log.Fatal("Error on opening folder", err)
							}

							if err == nil {
								fileNode.Children = children
							}

						} else if !fileNode.Expanded {
							fileNode.Children = nil
						}
					} else {
						state.SelectedFile = path + fileNode.Name

					}
				}
				// Create the list tile
				return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					icon := FileIcon(fileNode.Name, fileNode.IsDir, fileNode.Expanded)

					return material.Clickable(gtx, &fileNode.Button, func(gtx layout.Context) layout.Dimensions {
						return material.Body1(th, icon+" "+fileNode.Name+" ").Layout(gtx)
					})
				})
			}),
		//If directory and child elements are present we create sub renderNodes
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if fileNode.Expanded && len(fileNode.Children) > 0 {
				return layout.Flex{Axis: layout.Vertical}.Layout(
					gtx, 
					renderChildNodes( th, fileNode.Children, depth+1, path+fileNode.Name+"/", state)...
				)
			}

			return layout.Dimensions{}
		}),
	)

}

//	renderChildNodes is an internal function thet adds child nodes for a folder when expanded.
//
//		`gtx`		Graphical Context
//		`th`		Theme
//		`children`	Child Nodes
//		`depth`		Depth from root , root having 0 depth
//		`path`		Path to the file or folder
//		`state`		Reference to FileManagerState
func renderChildNodes(
	th *material.Theme,		
	children []*FileNode,	
	depth int,				
	path string,			
	state *internal.FileManagerState,	
) []layout.FlexChild {

	var items []layout.FlexChild

	for _, child := range children {

		node := child
		items = append(items, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return renderNode(gtx, th, node, depth, path, state)
		}))
	}

	return items
}
