package main

import "fmt"
import "github.com/kkulak/golang-astar/astar"

func main() {

	start, end := astar.ReadMapForMultiplePoints("resources/test08")

	// when
	distance, actualPath := astar.AStar(start, end)

	// then
	fmt.Print(distance)
	astar.PersistGraphWithMultiplePointsToBitmap(actualPath, "resources/test08")
}
