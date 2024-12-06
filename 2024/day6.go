package aoc2024

import (
	"aoc/utils"
	"bufio"
)

const (
	EMPTY = 0
	OBSTICLE = 1
	GUARD = 2
)

type Guard struct {
	position Position
	direction int
	isAboutToExit bool
}

type AreaMap struct {
	width int
	height int
	data []int
}

type Object struct {
	position Position
	objType int
}

func GetGuardVisitedAreaCount(scanner *bufio.Scanner) int {
	
	mapData := readMap(scanner)

	guard := generateGuard(mapData)
	var path []int
	var visited int
	isAboutToExit := false
	for i := 0; i < mapData.width * mapData.height; i++ {
		visited, isAboutToExit = findPathToExit(mapData, &guard)
		if visited > -1 && !isAboutToExit && !utils.IntContains(path, visited) {
			path = append(path, visited)
		}
		if isAboutToExit {
			path = append(path, visited)
			break
		}
	}
	return len(path)
}

func readMap(scanner *bufio.Scanner) AreaMap {
	var mapData AreaMap

	width := 0
	height := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		width = len(line)
		height++
		for _, char := range line {
			if char == '.' {
				mapData.data = append(mapData.data, EMPTY)
			} else if char == '#' {
				mapData.data = append(mapData.data, OBSTICLE)
			} else if char == '^' {
				mapData.data = append(mapData.data, GUARD)
			}
		}
	}

	mapData.width = width
	mapData.height = height

	return mapData
}

func generateGuard(mapData AreaMap) Guard {
	var pos Position
	for position, obj := range mapData.data {
		if obj == GUARD {
			pos.y = position / mapData.width
			pos.x = position % mapData.width
		}
	}
	isAboutToExit := isGuardAboutToExit(pos, TOP, mapData.width, mapData.height)
	return Guard{position: pos, direction: TOP, isAboutToExit: isAboutToExit}
}

func isGuardAboutToExit(pos Position, direction int, width int, height int) bool {
	isAboutToExit := false
	if direction == TOP {
		isAboutToExit = pos.y == 0
	} else if direction == BOTTOM {
		isAboutToExit = pos.y == height - 1
	} else if direction == LEFT {
		isAboutToExit = pos.x == 0
	} else if direction == RIGHT {
		isAboutToExit = pos.x == width - 1
	}
	return isAboutToExit
}

func findPathToExit(mapData AreaMap, guard *Guard) (int, bool) {
	var visited int
	isAboutToExit := false
	twoDMap := generate2DMap(mapData)

	if guard.direction == TOP {
		position := Position{x: guard.position.x, y: guard.position.y - 1}
		visited, isAboutToExit = traverse(guard, position, twoDMap, mapData.width, mapData.height)
		if isAboutToExit {
			return visited, true
		}
	} else

	if guard.direction == LEFT {
		position := Position{x: guard.position.x - 1, y: guard.position.y}
		visited, isAboutToExit = traverse(guard, position, twoDMap, mapData.width, mapData.height)
		if isAboutToExit {
			return visited, true
		}
	} else

	if guard.direction == BOTTOM {
		position := Position{x: guard.position.x, y: guard.position.y + 1}
		visited, isAboutToExit = traverse(guard, position, twoDMap, mapData.width, mapData.height)
		if isAboutToExit {
			return visited, true
		}
	} else

	if guard.direction == RIGHT {
		position := Position{x: guard.position.x + 1, y: guard.position.y}
		visited, isAboutToExit = traverse(guard, position, twoDMap, mapData.width, mapData.height)
		if isAboutToExit {
			return visited, true
		}
	}

	return visited, false
}

func findNextTile(mapData []Object, position Position) Object {
	var found Object
	for _, obj := range mapData {
		if obj.position.x == position.x && obj.position.y == position.y {
			found = obj
		}
	}

	return found
}

func generate2DMap(mapData AreaMap) []Object {
	var twoDMap []Object
	for index, obj := range mapData.data {
		position := Position{x: index % mapData.width, y: index / mapData.width}
		object := Object{position: position, objType: obj}

		twoDMap = append(twoDMap, object)
	}
	return twoDMap
}

func getNextDirection(direction int) int {
	var newDirection int
	if direction == TOP {
		newDirection = RIGHT
	} else if direction == RIGHT {
		newDirection = BOTTOM
	} else if direction == BOTTOM {
		newDirection = LEFT
	} else if direction == LEFT {
		newDirection = TOP
	}
	return newDirection
}

func traverse(guard *Guard, position Position, twoDmap []Object, width int, height int) (int, bool) {
	var visited int
	var isAboutToExit bool
	if !isGuardAboutToExit(guard.position, guard.direction, width, height) {
		object := findNextTile(twoDmap, position)
		if object.objType == EMPTY || object.objType == GUARD {
			visited = guard.position.y * width + guard.position.x
			isAboutToExit = false
			guard.position = position
		} else if object.objType == OBSTICLE {
			guard.direction = getNextDirection(guard.direction)
			isAboutToExit = false
			visited = -1
		}
	} else {
		isAboutToExit = true
		visited = guard.position.y * width + guard.position.x
	}
	return visited, isAboutToExit
}