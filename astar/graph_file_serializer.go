package astar

import (
	"os"
	"fmt"
)

func PersistGraphToFile(distance float64, nodes []Node, path string) {
	f, _ := os.Create(path + "_out.txt")
	fmt.Fprintf(f, "%d 1\n", int(distance))

	for _, entry := range nodes {
		astarNodeState := entry.(AStarNode).state
		coordinates := astarNodeState.coordinates
		velocity := astarNodeState.velocity
		fmt.Fprintf(f, "%d %d %d %d\n", coordinates.x, coordinates.y, velocity.x, velocity.y)
	}
}
