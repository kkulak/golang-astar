package astar

import (
	"github.com/deckarep/golang-set"
	"image/color"
	"golang.org/x/image/draw"
)

func ReadMapForMultiplePoints(path string) (Node, Node) {
	startingPointPosition := make([]AStarNodeState, 0)
	var endsPosition AStarNodeState
	obstacles := mapset.NewSet()

	bitmap := readBitmap(path + ".bmp")
	bitmapSize := bitmap.Bounds().Size()

	for x := 0; x < bitmapSize.X; x++ {
		for y := 0; y < bitmapSize.Y; y++ {
			pixel := bitmap.At(x, y)
			if isStart(pixel) {
				startingPointPosition = append(startingPointPosition, AStarNodeState{coordinates: Coordinates{x: x, y: y}, velocity: Velocity{x: 0, y: 0}})
			}
			if isEnd(pixel) {
				endsPosition = AStarNodeState{coordinates: Coordinates{x: x, y: y}, velocity: Velocity{x: 0, y: 0}}
			}
			if isObstacle(pixel) {
				obstacles.Add(Coordinates{x: x, y: y})
			}
		}
	}

	graph := MapOfSize(bitmapSize.X, bitmapSize.Y, obstacles)

	destination := make([]Coordinates, 0)
	for range startingPointPosition {
		destination = append(destination, endsPosition.coordinates)
	}

	startingPoints := make([]AStarNode, 0)
	endingPoints := make([]AStarNode, 0)
	for _, startingPointState := range startingPointPosition {
		startingPoints = append(startingPoints, graph.PointOf(startingPointState))
		endingPoints = append(endingPoints, graph.PointOf(endsPosition))
	}

	return MultiPointAstarNode{startingPoints, destination}, MultiPointAstarNode{endingPoints, destination}
}

func PersistGraphWithMultiplePointsToBitmap(path []Node, baseGraphPath string) {
	bitmap := readBitmap(baseGraphPath + ".bmp").(draw.Image)

	for _, aNode := range path {
		multiplePoints := aNode.(MultiPointAstarNode)
		for _, singlePoint := range multiplePoints.points {
			bitmap.Set(singlePoint.state.coordinates.x, singlePoint.state.coordinates.y, color.RGBA{R: 0, G: 0, B: 255, A: 255})
		}

	}

	writeBitmap(bitmap, baseGraphPath + "_out.bmp")
}
