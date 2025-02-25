package main

import (
	"image"
	"image/color"
	"math/rand"

	a "github.com/Setho0o/Journey/assets"
	u "github.com/Setho0o/Journey/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Matrix [][]*Cell

func MatrixInit() Matrix {
	m := make(Matrix, u.Ly/u.TileSize+1)
	for i := range m {
		m[i] = make([]*Cell, u.Lx/u.TileSize)
	}
	for i, row := range m {
		for j, _ := range row {
			x := j * u.TileSize
			y := i * u.TileSize
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			m[i][j] = &Cell{
				x:      x,
				y:      y,
				id:     a.Blank,
				op:     op,
				bounds: image.Rect(x, y, x+u.TileSize, y+u.TileSize), // added for easy collison detection image/rect has a lot of useful functions
			}
		}
	}
	return m
}

type Cell struct {
	x      int
	y      int
	id     a.Id
	op     *ebiten.DrawImageOptions
	bounds image.Rectangle
} 

func (g *Game)DrawBounds(color color.Color) {
  for _, e := range g.m {
		for _, c := range e {
		  for i := range c.bounds.Dx() {
        g.s.Set(c.x + i, c.y, color)
        g.s.Set(c.x + i, c.y + c.bounds.Dy(), color)
      }
      for i := range c.bounds.Dy() {
        g.s.Set(c.x, c.y + i, color)
        g.s.Set(c.x + c.bounds.Dx(), c.y + i, color)
      }
		}
	}   
}

func (g *Game) DrawMatrix() {
	for _, e := range g.m {
		for _, c := range e {
			g.s.DrawImage(g.a[c.id], c.op)
		}
	}
}

func (g *Game) FillWater() {
	for _, e := range g.m {
		for _, c := range e {
			switch rand.Intn(3) {
			case 0:
				c.id = a.WaterM1
			case 1:
				c.id = a.WaterM2
			case 2:
				c.id = a.WaterM3
			}
		}
	}
}
func (g *Game) FillLand() {

}
