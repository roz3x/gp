package main

import (
	"image"
	"image/color"

	"github.com/aquilax/go-perlin"
)

var (
	p   *perlin.Perlin
	img *image.RGBA
)

func init() {
	p = perlin.NewPerlin(2, 2, 3, 10)
}

func cloud() {
	img := image.NewRGBA(image.Rect(0, 0, N, N))
	z := 0.
	for {
		for i := 0.; i < N; i++ {
			for j := 0.; j < N; j++ {
				c := p.Noise3D(i/10, j/10, z)
				img.Set(int(i), int(j), color.NRGBA{
					R: uint8(c * 255),
					G: uint8(c * 255),
					B: uint8(c * 255),
				})
			}
			z += .1
		}
	}
}
