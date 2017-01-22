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

	var destination [100]Coordinates
	for idx:= range startingPointPosition {
		destination[idx] = endsPosition.coordinates
	}

	var startingPoints [100]AStarNode
	var endingPoints [100]AStarNode
	for idx, startingPointState := range startingPointPosition {
		startingPoints[idx] = graph.PointOf(startingPointState)
		endingPoints[idx] = graph.PointOf(endsPosition)
	}

	return MultiPointAstarNode{startingPoints, destination}, MultiPointAstarNode{endingPoints, destination}
}

func PersistGraphWithMultiplePointsToBitmap(path []Node, baseGraphPath string) {
	bitmap := readBitmap(baseGraphPath + ".bmp").(draw.Image)

	for _, aNode := range path {
		multiplePoints := aNode.(MultiPointAstarNode)
		for _, singlePoint := range multiplePoints.Points() {
			bitmap.Set(singlePoint.state.coordinates.x, singlePoint.state.coordinates.y, color.RGBA{R: 0, G: 0, B: 255, A: 255})
		}

	}

	writeBitmap(bitmap, baseGraphPath + "_out.bmp")
}
