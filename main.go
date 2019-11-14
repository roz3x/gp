package main

import (
	"image"
	"image/color"
	"image/draw"
	"log"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

func main() {
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(&screen.NewWindowOptions{
			Title: "cfd-application",
		})
		if err != nil {
			log.Fatal(err)
		}
		buf, tex := screen.Buffer(nil), screen.Texture(nil)
		defer func() {
			if buf != nil {
				tex.Release()
				buf.Release()
			}
			w.Release()
		}()

		go simulate(w)
		// go circle()
		var (
			buttonDown bool
			sz         size.Event
		)
		for {
			publish := false

			switch e := w.NextEvent().(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}

				switch e.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					pauseChan <- play
					var err error
					buf, err = s.NewBuffer(image.Point{N, N})
					if err != nil {
						log.Fatal(err)
					}
					tex, err = s.NewTexture(image.Point{N, N})
					if err != nil {
						log.Fatal(err)
					}
					tex.Fill(tex.Bounds(), color.White, draw.Src)

				case lifecycle.CrossOff:
					pauseChan <- pause
					tex.Release()
					tex = nil
					buf.Release()
					buf = nil
				}

			case mouse.Event:
				if e.Button == mouse.ButtonLeft {
					buttonDown = e.Direction == mouse.DirPress
				}
				if !buttonDown {
					break
				}
				z := sz.Size()
				x := int(e.X) * N / z.X
				y := int(e.Y) * N / z.Y

				if x < 0 || N <= x || y < 0 || N <= y {
					break
				}

				shared.mu.Lock()
				shared.mouseEvents = append(shared.mouseEvents, image.Point{x, y})
				shared.mu.Unlock()

			case paint.Event:
				publish = buf != nil

			case size.Event:
				sz = e

			case uploadEvent:
				shared.mu.Lock()
				if buf != nil {
					copy(buf.RGBA().Pix, shared.pix)
					publish = true
				}
				shared.uploadEventSent = false
				shared.mu.Unlock()

				if publish {
					tex.Upload(image.Point{}, buf, buf.Bounds())
				}

			case error:
				log.Print(e)
			}

			if publish {
				w.Scale(sz.Bounds(), tex, tex.Bounds(), draw.Src, nil)
				w.Publish()
			}
		}
	})
}
