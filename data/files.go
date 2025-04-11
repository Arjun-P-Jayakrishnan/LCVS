package data

import (
	"os"
	"path/filepath"
	"sort"

	"gioui.org/widget"
)
 
type FileNode struct {
	Name     string 			// 	Name of file / folder
	Path     string 			// 	Root Path of file / folder
	FileType string 			// 	Deprecated: not in use
	IsDir    bool   			// 	Shows whether the item is a file or folder
	Children []*FileNode 		//	if folder may have sub folders and files
	Expanded bool  				//  directory icon cabe in opened and closed states
	Button   widget.Clickable 	//	make them clickable tiles
}


//BuildFileTree builds the filenodes which consists of the details to render the file tree
func BuildFileTree(rootPath string) ([]*FileNode, error) {

	//obtain data from os
	entries, err := os.ReadDir(rootPath)

	if err != nil {
		return nil, err
	}

	//sort folder then files format
	sort.Slice(entries, func(i, j int) bool {
		a, b := entries[i], entries[j]

		if a.IsDir() != b.IsDir() {
			return a.IsDir()
		}

		return a.Name() < b.Name()
	})

	//marshall data
	var nodes []*FileNode

	for _, entry := range entries {

		info, err := entry.Info()

		if err != nil {
			continue
		}


		//try accessing files first for any permission issues 
		if info.IsDir(){
			if _,err := os.Open(filepath.Join(rootPath,entry.Name())); err!=nil{
				continue
			}
		}

		//initialize the data
		node := FileNode{
			Name:     entry.Name(),
			IsDir:    info.IsDir(),
			Children: nil,
			Expanded: false,
			Path:     filepath.Join(rootPath, entry.Name()),
		}

		nodes = append(nodes, &node)
	}

	return nodes, nil
}
