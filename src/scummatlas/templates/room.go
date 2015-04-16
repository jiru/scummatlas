package templates

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"scummatlas"
	l "scummatlas/condlog"
)

type roomData struct {
	Index      string
	Title      string
	Background string
	Boxes      [][4]scummatlas.Point
	scummatlas.Room
}

func (self roomData) PaletteHex() []string {
	var hexes []string
	hexes = make([]string, len(self.Palette))
	for i, color := range self.Palette {
		r, g, b, _ := color.RGBA()
		hexes[i] = fmt.Sprintf("%02x%02x%02x", uint8(r), uint8(g), uint8(b))
	}
	return hexes
}

func (self roomData) DoubleHeight() int {
	return self.Height * 2
}

func (self roomData) ViewBox() string {
	return fmt.Sprintf("-10 -10 %v %v", self.Width+10, self.Height+10)
}

func (self roomData) SvgWidth() int {
	return self.Width * 2
}

func (self roomData) SvgHeight() int {
	return self.Height * 2
}

func WriteRoom(room scummatlas.Room, index int, outputdir string) {

	roomTpl, err := ioutil.ReadFile("./templates/room.html")
	if err != nil {
		panic("No index.html in the templates directory")
	}

	t := template.Must(template.New("index").Parse(string(roomTpl)))

	bgPath := fmt.Sprintf("./img_bg/room%02d_bg.png", index)
	htmlPath := fmt.Sprintf("%v/room%02d.html", outputdir, index)
	file, err := os.Create(htmlPath)
	l.Log("template", "Create "+htmlPath)
	if err != nil {
		panic("Can't create room file")
	}

	var boxes [][4]scummatlas.Point

	for _, v := range room.Boxes {
		boxes = append(boxes, v.Corners())
	}

	data := roomData{
		fmt.Sprintf("%02d", index),
		"A room",
		bgPath,
		boxes,
		room,
	}
	t.Execute(file, data)
}
