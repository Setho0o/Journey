package utils

const (
	Lscale   = 2
	Wscale   = 1
	Lx       = Wx / Lscale // useable pixels
	Ly       = Wy / Lscale
	Wx       = 1920 / Wscale // window sizw
	Wy       = 1080 / Wscale
	TileSize = 16
	Mx       = Lx/TileSize - 1 // size of the matrix array
	My       = Ly / TileSize
)
