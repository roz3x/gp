package main

import (
	"image"
	"math"
	"time"
)

var (
	midptx = N / 2
	midpty = N / 2
	radx   = float64(50)
	rady   = float64(70)
)

func circle() {
	thita := 0.
	for {
		x := int(math.Sin(thita)*radx) + midptx
		y := int(math.Cos(thita)*rady) + midpty
		thita += 0.001
		// fmt.Printf("%v %v \n", x, y)
		shared.mu.Lock()
		shared.mouseEvents = append(shared.mouseEvents, image.Point{x, y})
		shared.mu.Unlock()
		time.Sleep(time.Millisecond * 10)
	}
}

func tsim() {
	for {
		time.Sleep(time.Second)
		t++
	}
}
