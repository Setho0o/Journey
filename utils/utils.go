package utils

const (
	scale   = 3 //1,2,3,6
	Lx       = Wx / scale // useable pixels
	Ly       = Wy / scale
	Wx       = 1920  // window sizw
	Wy       = 1080 
	TileSize = 20
	Mx       = Lx/TileSize - 1 // size of the matrix array
	My       = Ly / TileSize
)
