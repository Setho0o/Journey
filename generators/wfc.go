package generators

import (
	"math/rand/v2"

	m "github.com/Setho0o/Journey/matrix"
	"github.com/Setho0o/Journey/utils"
)

type Stack []int

type Neighbor int

const (
  Tl Neighbor = iota
  T 
  TR
  L
  R
  BL
  B
  BR
)



func (s Stack) Push(i, j int) Stack {
	return append(s, i, j)
}

func (s Stack) Pop(index int) (int, int) {
	l := len(s)
	return s[l - index - 1], s[ l - index]
}



type Wfc struct {
  m m.Matrix
  s Stack

}

func (w *Wfc) Wfc(){
  x := rand.IntN(utils.Mx)
  y := rand.IntN(utils.My)
  w.Collapse(w.m[y][x]) 

	var done bool
	done = false
	for done == false {
		x, y, done = w.FindNextCell()
		w.Collapse(w.m[y][x])
	}

}

func (w *Wfc) Collapse(c *m.Cell) {
	w.s.Push(c.X, c.Y)

	//set tile
      
}

func (w *Wfc) FindNextCell() (int, int, bool) {



}


