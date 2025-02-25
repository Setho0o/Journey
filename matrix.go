package main

import (
	"image"

	a "github.com/Setho0o/Journey/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Matrix [][]*Cell

func MatrixInit() Matrix {
	m := make(Matrix, Ly/TileSize)
	for i := range m {
		m[i] = make([]*Cell, Lx/TileSize)
	}
    
	for i, row := range m {
		for j, _ := range row {
      x := j * TileSize
      y := i * TileSize
      op := &ebiten.DrawImageOptions{}
      op.GeoM.Translate(float64(x),float64(y))
			m[i][j] = &Cell{
				x:  x,
				y:  y,
				id: a.LandM,
        op : op,
        bounds: image.Rect(x,y,x+TileSize,y+TileSize), // added for easy collison detection image/rect has a lot of useful functions 
			}
		}
	}
	return m
}

type Cell struct {
	x  int
	y  int
  id a.Id
  op  *ebiten.DrawImageOptions
  bounds image.Rectangle
}

func (g *Game) DrawMatrix() {
  for _, e := range g.m {
    for _, c := range e {
      g.s.DrawImage(g.a[c.id],c.op)
    }
  }
}
