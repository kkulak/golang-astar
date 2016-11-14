package astar

import (
	"os"
	"golang.org/x/image/bmp"
	"image"
	"image/color"
)

func FindShortestPathInGraph(path string) (float64, []Node) {
	var start, end CartesianCoordinates
	var obstacles []CartesianCoordinates = make([]CartesianCoordinates, 0)

	bitmap := readBitmap(path)
	bitmapSize := bitmap.Bounds().Size()

	for x := 0; x < bitmapSize.X; x++ {
		for y := 0; y < bitmapSize.Y; y++ {
			pixel := bitmap.At(x, y)
			if isStart(pixel) {
				start = CartesianCoordinates{x: x, y: y}
			}
			if isEnd(pixel) {
				end = CartesianCoordinates{x: x, y: y}
			}
			if isObstacle(pixel) {
				obstacles = append(obstacles, CartesianCoordinates{x: x, y: y})
			}
		}
	}

	graph := MapOfSize(bitmapSize.X, bitmapSize.Y, obstacles)
	return AStar(graph.PointOf(start), graph.PointOf(end))
}

func isStart(pixel color.Color) bool {
	return false
}

func isEnd(pixel color.Color) bool {
	return false
}

func isObstacle(pixel color.Color) bool {
	return false
}

func readBitmap(path string) image.Image {
	file, _ := os.Open(path)
	bitmap, _ := bmp.Decode(file)
	return bitmap
}
