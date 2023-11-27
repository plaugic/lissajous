package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {

}

func lissajous(out io.Writer) {
	const (
		cycles = 5 // number of complete oscillator revolutions
		res = 0.001
		size = 100
		nframes = 64

	)
}
