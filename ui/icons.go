package ui

import (
	"strings"
)

func FileIcon(name string, isDir bool, expanded bool) string {

	icon := "📄"

	if isDir {
		if expanded {
			icon = "📂"

		} else {
			icon = "📁"
		}
		return icon
	}

	switch {
	case strings.HasSuffix(name, ".go"):
		icon = "🐹"
	case strings.HasSuffix(name, ".md"):
		icon = "🇲🇩"
	case strings.HasSuffix(name, ".json"):
		icon = "{ }"
	default:
		icon = "📄"
	}

	return icon

}
