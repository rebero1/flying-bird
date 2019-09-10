package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type scene struct {
	time  int
	bg    *sdl.Texture
	birds []*sdl.Texture
}

func newScene(r *sdl.Renderer) (*scene, error) {
	bg, err := img.LoadTexture(r, "res/imgs/background.png")

	if err != nil {
		return nil, fmt.Errorf("Couldn't load background image:%v", err)
	}

	var birds []*sdl.Texture

	for i := 1; i < 5; i++ {

		path := fmt.Sprintf("res/imgs/bird_frame_%d.png", i)
		bird, err := img.LoadTexture(r, path)

		if err != nil {
			return nil, fmt.Errorf("Couldn't load background image:%v", err)
		}

		birds = append(birds, bird)
	}

	return &scene{
		bg:    bg,
		birds: birds,
	}, nil
}

func (s *scene) paint(r *sdl.Renderer) error {
	var xl int32
	xl = 10
	for {

		time.Sleep(100 * time.Millisecond)
		s.time++
		r.Clear()
		if err := r.Copy(s.bg, nil, nil); err != nil {
			return fmt.Errorf("ERROR COULDN'T COPY:%v ", err)
		}

		rect := &sdl.Rect{
			X: xl,
			Y: 300 - 43/2,
			W: 50,
			H: 43,
		}
		i := s.time % len(s.birds)
		fmt.Print(i)

		if err := r.Copy(s.birds[i], nil, rect); err != nil {
			return fmt.Errorf("ERROR COULDN'T COPY:%v ", err)
		}

		r.Present()
		xl = (xl + 
			5) % 800
	}

	return nil

}

func (s *scene) destroy() {
	s.bg.Destroy()
}
