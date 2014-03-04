package util

type Float32Grid struct {
	Xs []float32
	Height, Width int32
}

func (g *Float32Grid) Get(x, y int32) float32 {
	return g.Xs[x + y * g.Width]
}

func (g *Float32Grid) Set(x, y int32, val float32) {
	g.Xs[x + y * g.Width] = val
}

func (g *Float32Grid) Up(y int32) int32 {
	return (y + 1) % g.Height
}

func (g *Float32Grid) Down(y int32) int32 {
	return (y - 1 + g.Height) % g.Height
}

func (g *Float32Grid) Right(x int32) int32 {
	return (x + 1) % g.Width
}

func (g *Float32Grid) Left(x int32) int32 {
	return (x - 1 + g.Width) % g.Width
}

type Int32Grid struct {
	Xs []int32
	Height, Width int32
}

func (g *Int32Grid) Get(x, y int32) int32 {
	return g.Xs[x + y * g.Width]
}

func (g *Int32Grid) Set(x, y int32, val int32) {
	g.Xs[x + y * g.Width] = val
}

func (g *Int32Grid) Up(y int32) int32 {
	return (y + 1) % g.Height
}

func (g *Int32Grid) Down(y int32) int32 {
	return (y - 1 + g.Height) % g.Height
}

func (g *Int32Grid) Right(x int32) int32 {
	return (x + 1) % g.Width
}

func (g *Int32Grid) Left(x int32) int32 {
	return (x - 1 + g.Width) % g.Width
}

type Float64Grid struct {
	Xs []float64
	Height, Width int32
}

func (g *Float64Grid) Get(x, y int32) float64 {
	return g.Xs[x + y * g.Width]
}

func (g *Float64Grid) Set(x, y int32, val float64) {
	g.Xs[x + y * g.Width] = val
}

func (g *Float64Grid) Up(y int32) int32 {
	return (y + 1) % g.Height
}

func (g *Float64Grid) Down(y int32) int32 {
	return (y - 1 + g.Height) % g.Height
}

func (g *Float64Grid) Right(x int32) int32 {
	return (x + 1) % g.Width
}

func (g *Float64Grid) Left(x int32) int32 {
	return (x - 1 + g.Width) % g.Width
}

type Int64Grid struct {
	Xs []int64
	Height, Width int32
}

func (g *Int64Grid) Get(x, y int32) int64 {
	return g.Xs[x + y * g.Width]
}

func (g *Int64Grid) Set(x, y int32, val int64) {
	g.Xs[x + y * g.Width] = val
}

func (g *Int64Grid) Up(y int32) int32 {
	return (y + 1) % g.Height
}

func (g *Int64Grid) Down(y int32) int32 {
	return (y - 1 + g.Height) % g.Height
}

func (g *Int64Grid) Right(x int32) int32 {
	return (x + 1) % g.Width
}

func (g *Int64Grid) Left(x int32) int32 {
	return (x - 1 + g.Width) % g.Width
}
