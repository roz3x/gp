package main

import (
	"image"
	"sync"
	"time"
)

const (
	//N is buffer size
	N = 300

	tickDuration = time.Second / 60
	iterations   = 20
	dt           = 0.1
	diff         = 0
	visc         = 0.01
	forcex       = 1
	forcey       = 10
	source       = 100
	fade         = 0.9
)

const (
	pause = false
	play  = true
)

var pauseChan = make(chan bool, 100)

type array [N + 2][N + 2]float32
type uploadEvent struct{}

var shared = struct {
	mu              sync.Mutex
	uploadEventSent bool
	mouseEvents     []image.Point
	pix             []byte
}{
	pix: make([]byte, 4*N*N),
}
