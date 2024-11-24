package ui

import (
	"bytes"
	_ "embed"
	"fmt"
	"image/png"
	"log"

	"gioui.org/op"       // used for recording different events
	"gioui.org/op/paint" // paint contains operations for coloring
	"github.com/Arjun-P-Jayakrishnan/LCVS/handlers"
)

//store image as byte array


func imageOp(imageData []byte) paint.ImageOp {

	m,err := png.Decode(bytes.NewReader(imageData))

	if err!= nil {
		fmt.Println(err)
		panic(err)
		log.Fatal(err)
	}

	return paint.NewImageOp(m)
}

func RenderImage(path string,ops *op.Ops ) {
	imageData:=handlers.ReadFromFileAsBytes(path)
	imageOp(imageData).Add(ops)
	paint.PaintOp{}.Add(ops)
}

