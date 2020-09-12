package main

import "fmt"

const (
	Empty int = 0
	Alive int = 1
)

type Grid struct {
	height int
	width  int
	rows   [][]int
}

func (g *Grid) New(height int, width int) {
	g.height = height
	g.width = width

	g.rows = make([][]int, height)
	for row := 0; row < height; row++ {
		rawSlice := make([]int, width)
		for col := 0; col < width; col++ {
			rawSlice[col] = Empty
		}
		g.rows[row] = rawSlice
	}
}

func (g *Grid) get(y int, x int) int {
	if y < 0 {
		y += g.height
	}
	if x < 0 {
		x += g.width
	}
	return g.rows[y%g.height][x%g.width]
}

func (g *Grid) set(y int, x int, state int) {
	g.rows[y%g.height][x%g.width] = state
	return
}

func (g *Grid) countNeighbors(y int, x int) int {
	nn := g.get(y-1, x+0) // North
	ne := g.get(y-1, x+1) // Northeast
	ee := g.get(y-0, x+1) // East
	se := g.get(y+1, x+1) // Southeast
	ss := g.get(y+1, x+0) // South
	sw := g.get(y+1, x-1) // Southwest
	ww := g.get(y-0, x-1) // West
	nw := g.get(y-1, x-1) // Northwest

	count := 0
	for _, state := range []int{nn, ne, ee, se, ss, sw, ww, nw} {

		if state == Alive {
			count++
		}
	}
	return count
}

type lifeGame interface {
	set(y int, x int, state int)
	get(y int, x int) int
	countNeighbors(y int, x int) int
}

func gameLogic(state int, neighbors int) int {
	if state == Alive {
		if neighbors == 3 || neighbors == 2 {
			return Alive
		}
		return Empty
	}

	if neighbors == 3 {
		return Alive
	}

	return state
}

func stepCell(y int, x int, game lifeGame, nextGame lifeGame) {

	state := game.get(y, x)
	neighbors := game.countNeighbors(y, x)
	nextState := gameLogic(state, neighbors)
	nextGame.set(y, x, nextState)
}

func (g *Grid) simulate() (nextGrid *Grid) {

	nextGrid = &Grid{}
	nextGrid.New(g.height, g.width)
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			go stepCell(y, x, g, nextGrid)
		}
	}
	return nextGrid
}

func (g *Grid) print() {
	var m = map[int]string{
		Alive: "*",
		Empty: "-",
	}
	for _, row := range g.rows {
		rowPrint := ""
		for _, col := range row {
			rowPrint += m[col]
		}
		rowPrint += "\n"
		fmt.Println(rowPrint)
	}
}

func main() {
	grid := &Grid{}
	grid.New(5, 9)
	grid.set(0, 3, Alive)
	grid.set(1, 4, Alive)
	grid.set(2, 2, Alive)
	grid.set(2, 3, Alive)
	grid.set(2, 4, Alive)

	for i := 0; i < 5; i++ {
		fmt.Println("Step: ", i)
		grid.print()
		grid = grid.simulate()
	}
}
