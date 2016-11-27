package astar

import (
	"os"
	"golang.org/x/image/bmp"
	"image"
	"image/color"
	"image/draw"
)

func GraphFromBitmap(path string) (Graph, Node, Node) {
	var start, end AStarNodeState
	var obstacles []AStarNodeState = make([]AStarNodeState, 0)

	bitmap := readBitmap(path + ".bmp")
	bitmapSize := bitmap.Bounds().Size()

	for x := 0; x < bitmapSize.X; x++ {
		for y := 0; y < bitmapSize.Y; y++ {
			pixel := bitmap.At(x, y)
			if isStart(pixel) {
				start = AStarNodeState{coordinates: Coordinates{x: x, y: y}, velocity: Velocity{x: 0, y: 0}}
			}
			if isEnd(pixel) {
				end = AStarNodeState{coordinates: Coordinates{x: x, y: y}, velocity: Velocity{x: 0, y: 0}}
			}
			if isObstacle(pixel) {
				obstacles = append(obstacles, AStarNodeState{coordinates: Coordinates{x: x, y: y}})
			}
		}
	}

	graph := MapOfSize(bitmapSize.X, bitmapSize.Y, obstacles)
	return graph, graph.PointOf(start), graph.PointOf(end)
}

func PersistGraphToBitmap(path []Node, baseGraphPath string) {
	bitmap := readBitmap(baseGraphPath + ".bmp").(draw.Image)

	for _, aNode := range path {
		astarNode := aNode.(AStarNode)
		bitmap.Set(astarNode.state.coordinates.x, astarNode.state.coordinates.y, color.RGBA{R: 0, G: 0, B: 255, A: 255})
	}

	writeBitmap(bitmap, baseGraphPath + "_out.bmp")
}

func isStart(pixel color.Color) bool {
	return pixel == color.RGBA{ R: 0, G: 255, B: 0, A: 255 }
}

func isEnd(pixel color.Color) bool {
	return pixel == color.RGBA{ R: 255, G: 0, B: 0, A: 255 }
}

func isObstacle(pixel color.Color) bool {
	return pixel == color.RGBA{ R: 0, G: 0, B: 0, A: 255 }
}

func readBitmap(path string) image.Image {
	file, _ := os.Open(path)
	bitmap, _ := bmp.Decode(file)
	return bitmap
}

func writeBitmap(image image.Image, path string) {
	res, _ := os.Create(path)
	bmp.Encode(res, image)
}