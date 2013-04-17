package display

import "fmt"

func getGlyphPositions() (g map[string][2]int) {
	g[" "] = [2]int{0, 0}
	g["face"] = [2]int{1, 0}
	g["angry"] = [2]int{2, 0}
	g["heart"] = [2]int{3, 0}
	g["diamond"] = [2]int{4, 0}
	g["clubs"] = [2]int{5, 0}
	g["spades"] = [2]int{6, 0}
	g["oval"] = [2]int{7, 0}
	g["circle"] = [2]int{9, 0}
	g["female"] = [2]int{11, 0}
	g["male"] = [2]int{12, 0}
	g["note"] = [2]int{13, 0}
	g["notes"] = [2]int{14, 0}
	g["star"] = [2]int{15, 0}
	g["thickleft"] = [2]int{0, 1}
	g["thickright"] = [2]int{1, 1}
	g["updown"] = [2]int{2, 1}
	g["!!"] = [2]int{3, 1}
	g["nspc"] = [2]int{4, 1}
	g["ampersand"] = [2]int{5, 1}
	g["underscore"] = [2]int{6, 1}
	g["updownline"] = [2]int{7, 1}
	g["up"] = [2]int{8, 1}
	g["down"] = [2]int{9, 1}
	g["right"] = [2]int{10, 1}
	g["left"] = [2]int{11, 1}
	g["hook"] = [2]int{12, 1}
	g["leftright"] = [2]int{13, 1}
	g["thickup"] = [2]int{14, 1}
	g["thickdown"] = [2]int{15, 1}

	g["!"] = [2]int{1, 2}
	g["\""] = [2]int{2, 2}
	g["#"] = [2]int{3, 2}
	g["$"] = [2]int{4, 2}
	g["%"] = [2]int{5, 2}
	g["&"] = [2]int{6, 2}
	g["'"] = [2]int{7, 2}
	g["("] = [2]int{8, 2}
	g[")"] = [2]int{9, 2}
	g["*"] = [2]int{10, 2}
	g["+"] = [2]int{11, 2}
	g[","] = [2]int{12, 2}
	g["-"] = [2]int{13, 2}
	g["."] = [2]int{14, 2}
	g["/"] = [2]int{15, 2}

	for i := 0; i < 10; i++ {
		g[fmt.Sprintf("%v", i)] = [2]int{i, 3}
	}
	g[":"] = [2]int{10, 3}
	g[";"] = [2]int{11, 3}
	g["<"] = [2]int{12, 3}
	g["="] = [2]int{13, 3}
	g[">"] = [2]int{14, 3}
	g["?"] = [2]int{15, 3}
	
	for i, a := range "@ABCDEFGHIJKLMNO" {
		g[fmt.Sprintf("%v", a)] = [2]int{i, 4}
	}

	for i, a := range "PQRSTUVWXYZ[\\]^_" {
		g[fmt.Sprintf("%v", a)] = [2]int{i, 5}
	}

	for i, a := range "`abcdefghijlkmno" {
		g[fmt.Sprintf("%v", a)] = [2]int{i, 6}
	}

	for i, a := range "pqrstuvwxyz{|}~" {
		g[fmt.Sprintf("%v", a)] = [2]int{i, 7}
	}
	g["triangle"] = [2]int{15, 7}

	return g
}