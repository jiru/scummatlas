package image

import (
	"fmt"
	"image"
	"image/color"
	b "scummatlas/binaryutils"
	l "scummatlas/condlog"
)

func ParsePalette(data []byte) color.Palette {
	var palette = make(color.Palette, 0, 256)
	for i := 0; i < len(data); i += 3 {
		color := color.RGBA{
			data[i],
			data[i+1],
			data[i+2],
			255,
		}
		palette = append(palette, color)
	}

	return palette
}

func ParseImage(data []byte, zBuffers int, width int, height int, pal color.Palette, transpIndex uint8) (image *image.RGBA, zplane *image.RGBA) {
	blockName := string(data[8:12])

	if blockName == "BOMP" {
		l.Log("image", "BOMP not implemented yet")
		return nil, nil
	}
	if blockName != "SMAP" {
		panic("No stripe table found, " + blockName + " found instead")
	}
	blockSize := b.BE32(data, 12)
	l.Log("image", "SmapSize %v", blockSize)

	stripeCount := width / 8
	offsets := parseStripeTable(data, 16, stripeCount, 4)
	//fmt.Println("offsets", offsets)
	image = parseStripesIntoImage(data, offsets, 8, width, height, pal, transpIndex)

	if b.FourCharString(data, 8+blockSize) == "ZP01" {
		fmt.Println("ZP01 found")
	}
	offsets = parseStripeTable(data, 16+blockSize, stripeCount, 2)
	fmt.Println("ZP01 offsets")
	zplane = parseZplaneStripesIntoImage(data, offsets, 8+blockSize, height)

	l.Log("image", "image decoded\n")
	return
}

func parseZplaneStripesIntoImage(data []byte, offsets []int, initialOffset int, height int) *image.RGBA {
	width := len(offsets) * 8
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	img.Set(10, 10, color.RGBA{255, 0, 0, 1})

	for i, offset := range offsets {
		fmt.Printf("%d\t%x\n", i, offset)
	}

	return img
}

func parseStripesIntoImage(data []byte, offsets []int, initialOffset int, width int, height int, pal color.Palette, transpIndex uint8) *image.RGBA {
	l.Log("image", "Decoding image")
	l.Log("image", "Stripes information")
	l.Log("image", "\n#ID\tCode\tMethod\tDirect\tPalInSz\tTransp")

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < len(offsets); i++ {
		offset := offsets[i]
		size := len(data) - offset
		if i < len(offsets)-1 {
			size = offsets[i+1] - offsets[i]
		}
		if l.Logflags["image"] {
			printStripeInfo(i, data[initialOffset+offset])
		}
		if len(data) < initialOffset+offset+size {
			return img
		}
		drawStripe(
			img,
			i,
			data[initialOffset+offset:initialOffset+offset+size],
			pal,
			transpIndex)
	}
	return img
}

func parseStripeTable(data []byte, offset int, stripeCount int, offsetSize int) []int {
	offsets := make([]int, 0, stripeCount)
	stripeOffset := 0
	for i := 0; i < stripeCount; i++ {
		if offsetSize == 4 {
			stripeOffset = b.LE32(data, offset+4*i)
		} else {
			stripeOffset = b.LE16(data, offset+2*i)
		}
		offsets = append(offsets, stripeOffset)
	}
	return offsets
}
