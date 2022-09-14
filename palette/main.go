package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Color struct {
	r uint8
	g uint8
	b uint8
	a uint8
}

func checkColor(colors []string) ([]Color, error) {
	error := false
	ret := []Color{}
	for _, v := range colors {
		if len(v) != 3 && len(v) != 6 {
			error = true
			break
		}
		v = strings.ToLower(v)
		if !regexp.MustCompile(`^[a-f0-9]+$`).MatchString(v) {
			error = true
			break
		}
		if len(v) == 3 {
			v = v[0:1] + v[0:1] + v[1:2] + v[1:2] + v[2:3] + v[2:3]
		}

		r, _ := strconv.ParseInt(v[0:2], 16, 64)
		g, _ := strconv.ParseInt(v[2:4], 16, 64)
		b, _ := strconv.ParseInt(v[4:6], 16, 64)
		a := 255
		ret = append(ret, Color{uint8(r), uint8(g), uint8(b), uint8(a)})
	}

	if error {
		return nil, fmt.Errorf("An invalid color string is specified.")
	} else {
		return ret, nil
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: palette [Hex Color 1] [Hex Color 2] [Hex Color 3] ...")
		fmt.Println("Usage: palette 111 111222 123456 abcdef")
		return
	}

	colors, err := checkColor(os.Args[1:])
	if err != nil {
		panic(err)
	}

	size := 100

	img := image.NewRGBA(image.Rect(0, 0, len(colors) * size, size))
	for i, c := range colors {
		draw.Draw(img, image.Rect(i * size, 0, i * size + size, size), &image.Uniform{color.RGBA{c.r, c.g, c.b, c.a}}, image.Point{0, 0}, draw.Src)
	}

	out, err := os.Create("palette.png")
	if err != nil {
		panic(err)
	}

	png.Encode(out, img)
}