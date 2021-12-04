package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	p := make([][]uint8, dy)
	for x := range p {
		p[x] = make([]uint8, dx)
		for y := range p[x] {
			//p[x][y] = uint8((x + y) / 2)
			//p[x][y] = uint8(x * y)
			p[x][y] = uint8(x * x * y * y)
		}
	}
	return p
}

func main() {
	pic.Show(Pic)
}
