package main

import (
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/ttf"

	"github.com/veandco/go-sdl2/sdl"
)

func run() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)

	if err != nil {
		return fmt.Errorf("Couldn't iniitailize")
	}

	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		return fmt.Errorf("Couldn't initialize ttf:%v", err)
	}

	w, r, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)

	if err != nil {
		return fmt.Errorf("Error in creating and rendering: %v", err)
	}

	defer w.Destroy()
	err = drawTitle(r)

	if err != nil {
		return fmt.Errorf("Couldn't not draw title:%v", err)
	}

	running := true

	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			time.Sleep(3 * time.Second)
			fmt.Println("Quit")
			running = false
			break
		}
	}
	running = true

	s, err := newScene(r)

	if err != nil {
		return fmt.Errorf("Couldn't draw scene:%v", err)
	}

	defer s.destroy()
	if err := s.paint(r); err != nil {
		return fmt.Errorf("Could not paint scene:%v", err)
	}

	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
	}
	return nil
}

func drawTitle(r *sdl.Renderer) error {
	r.Clear()
	f, err := ttf.OpenFont("res/font/Flappy.ttf", 20)
	if err != nil {
		return fmt.Errorf("COULDN'T LOAD FONT: %v", err)
	}

	defer f.Close()
	color := sdl.Color{
		R: 255,
		G: 100,
		B: 0,
		A: 255,
	}
	s, err := f.RenderUTF8Solid("Funny Time", color)
	if err != nil {
		return fmt.Errorf("Couldn't not render title:%v", err)
	}

	defer s.Free()
	t, err := r.CreateTextureFromSurface(s)
	if err != nil {
		return fmt.Errorf("COuldn't create texture:%v", err)
	}
	if err := r.Copy(t, nil, nil); err != nil {
		fmt.Errorf("Couldn't copy texture:%v", err)
	}
	r.Present()
	defer t.Destroy()
	return nil
}
func main() {

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}

}
