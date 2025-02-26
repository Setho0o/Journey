package matrix

import (
	"image"

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
				X:      x,
				Y:      y,
				Id:     a.Blank,
				Op:     op,
				Bounds: image.Rect(x, y, x+u.TileSize, y+u.TileSize), // added for easy collison detection image/rect has a lot of useful functions
	      Entropy: a.Entropy,	
      }
		}
	}
	return m
}

type Cell struct {
	X      int
	Y      int
	Id     a.Id
	Op     *ebiten.DrawImageOptions
	Bounds image.Rectangle
  Entropy int
} 

