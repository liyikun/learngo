package main

import (
	"image"
	"fmt"
	"image/color"
)

type Image struct {

}

func (i Image) Bounds() image.Rectangle{
	return image.Rect(0,0,200,200)
}

func (i Image) ColorModel() color.Model{
	return color.RGBAModel
}

func (i Image) At(x,y int) color.Color {
	return color.RGBA{uint8(x),uint8(y),uint8(255),uint8(255)}
}

func main()  {
	m := image.NewRGBA(image.Rect(0,0,100,100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}