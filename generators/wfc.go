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

func (s Stack) Pop() (Stack, int, int) {
	l := len(s)
	return s[:l-2], s[l-2], s[l-1]
}



type Wfc struct {
  m m.Matrix
  s Stack

}
func (w *Wfc) Wfc(){
  x := rand.IntN(utils.Mx)
  y := rand.IntN(utils.My)
  w.Collapse(w.m[y][x]) 

}

func (w *Wfc) Collapse(c *m.Cell) {
  
  w.CheckNeighbors(c)

  // Figure out which cell have what entropy
  // add Cells to stack 
  // Collapse lowest cell 
  // Update entropy 
      
}

func (w *Wfc) CheckNeighbors(c *m.Cell)  {
  

}


