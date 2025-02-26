package main

import (
	"image/color"
	"math/rand"

	a "github.com/Setho0o/Journey/assets"
	_ "github.com/Setho0o/Journey/matrix"
)

func (g *Game)DrawBounds(color color.Color) {
  for _, e := range g.m {
		for _, c := range e {
		  for i := range c.Bounds.Dx() {
        g.s.Set(c.X + i, c.Y, color)
        g.s.Set(c.X + i, c.Y + c.Bounds.Dy(), color)
      }
      for i := range c.Bounds.Dy() {
        g.s.Set(c.Y, c.Y + i, color)
        g.s.Set(c.Y + c.Bounds.Dx(), c.Y + i, color)
      }
		}
	}
}

func (g *Game) DrawMatrix() {
	for _, e := range g.m {
		for _, c := range e {
			g.s.DrawImage(g.a[c.Id], c.Op)
		}
	}
}

func (g *Game) FillWater() {
	for _, e := range g.m {
		for _, c := range e {
			switch rand.Intn(3) {
			case 0:
				c.Id = a.WaterM1
			case 1:
				c.Id = a.WaterM2
			case 2:
				c.Id = a.WaterM3
			}
		}
	}
}
func (g *Game) FillLand() {
  for _, e := range g.m {
		for _, c := range e {
			switch rand.Intn(3) {
			case 0:
				c.Id = a.LandM1
			case 1:
				c.Id = a.LandM2
			case 2:
				c.Id = a.LandM3
			}
		}
	}
}
