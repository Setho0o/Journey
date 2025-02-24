package main


type Matrix [][]*Cell

func MatrixInit() Matrix {
	m := make(Matrix, Ly/TileSize)
	for i := range m {
		m[i] = make([]*Cell, Lx/TileSize)
	}

	for i, row := range m {
		for j, _ := range row {
			m[i][j] = &Cell {
				x:     j * TileSize,
				y:     i * TileSize,
			}
		}
	}
	return m
}


type Cell struct {
   x int 
   y int
}
