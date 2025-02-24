package main

import "github.com/hajimehoshi/ebiten/v2"

type Matrix [][]Cell
type id int 
type Cell struct {
  img *ebiten.Image
}
