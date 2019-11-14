package main

import (
	"image"
	"math"
	"math/rand"
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
	rndpt := 0
	for {
		rndpt = int(rand.Int31n(1))
		x := int(math.Sin(thita)*radx) + midptx + rndpt
		y := int(math.Cos(thita)*rady) + midpty + rndpt
		thita += 0.001
		// fmt.Printf("%v %v \n", x, y)
		shared.mu.Lock()
		shared.mouseEvents = append(shared.mouseEvents, image.Point{x, y})
		shared.mu.Unlock()
		time.Sleep(time.Millisecond)
	}
}

func tsim() {
	for {
		time.Sleep(time.Second)
		t++
	}
}
