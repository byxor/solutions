package main

import (
	"fmt"
)

var numSegments int
var segments [50]lineSegment

var triangles map[triangle]bool
var numTriangles int

func main() {
	for {
		// Set initial variables
		numTriangles = 0

		// Get initial input
		fmt.Scan(&numSegments)
		if numSegments == 0 {
			return
		}

		// Read line segments
		for i := 0; i < numSegments; i++ {
			segments[i] = lineSegment{id: i, intersections: make(map[int]bool)}
			fmt.Scanf("%f %f %f %f\n", &segments[i].ax, &segments[i].ay, &segments[i].bx, &segments[i].by)
		}

		// Gather intersection information
		for id0 := 0; id0 < numSegments; id0++ {
			for id1 := 0; id1 < numSegments; id1++ {
				if id0 == id1 {
					continue
				}

				if _, found := segments[id0].intersections[id1]; found {
					continue
				}

				if intersects(segments[id0], segments[id1]) {
					segments[id0].intersections[id1] = true
					segments[id1].intersections[id0] = true
				}
			}
		}

		// Search for triangles
		triangles = make(map[triangle]bool)
		for id0 := 0; id0 < numSegments; id0++ {
			segment := segments[id0]
			for id1, _ := range segment.intersections {
				for id2, _ := range segments[id1].intersections {

					if _, found := segments[id2].intersections[segment.id]; !found {
						continue
					}

					if id0 == id2 || id1 == id2 || id0 == id1 {
						continue
					}

					if _, found := triangles[triangle{id0, id1, id2}]; found {
						continue
					}

					if _, found := triangles[triangle{id0, id2, id1}]; found {
						continue
					}

					if _, found := triangles[triangle{id1, id2, id0}]; found {
						continue
					}

					if _, found := triangles[triangle{id1, id0, id2}]; found {
						continue
					}

					if _, found := triangles[triangle{id2, id0, id1}]; found {
						continue
					}

					if _, found := triangles[triangle{id2, id1, id0}]; found {
						continue
					}

					triangles[triangle{id0, id1, id2}] = true
					numTriangles++
				}
			}
		}

		fmt.Println(numTriangles)
	}
}

type point struct {
	x, y float64
}

type lineSegment struct {
	id             int
	ax, ay, bx, by float64
	intersections  map[int]bool
}

type triangle struct {
	a, b, c int
}

// INTERSECTION JUNK

const (
	colinear = iota
	clockwise
	counterclockwise
)

func intersects(segmentA, segmentB lineSegment) bool {
	o1 := orientation(segmentA.ax, segmentA.ay, segmentA.bx, segmentA.by, segmentB.ax, segmentB.ay)
	o2 := orientation(segmentA.ax, segmentA.ay, segmentA.bx, segmentA.by, segmentB.bx, segmentB.by)
	o3 := orientation(segmentB.ax, segmentB.ay, segmentB.bx, segmentB.by, segmentA.ax, segmentA.ay)
	o4 := orientation(segmentB.ax, segmentB.ay, segmentB.bx, segmentB.by, segmentA.bx, segmentA.by)

	if (o1 != o2) && (o3 != o4) {
		return true
	}

	if (o1 == colinear) && onSegment(segmentA.ax, segmentA.ay, segmentB.ax, segmentB.ay, segmentA.bx, segmentA.by) {
		return true
	}

	if (o2 == colinear) && onSegment(segmentA.ax, segmentA.ay, segmentB.bx, segmentB.by, segmentA.bx, segmentA.by) {
		return true
	}

	if (o3 == colinear) && onSegment(segmentB.ax, segmentB.ay, segmentA.ax, segmentA.ay, segmentB.bx, segmentB.by) {
		return true
	}

	if (o4 == colinear) && onSegment(segmentB.ax, segmentB.ay, segmentA.bx, segmentA.by, segmentB.bx, segmentB.by) {
		return true
	}

	return false
}

func orientation(px, py, qx, qy, rx, ry float64) int {
	val := ((qy - py) * (rx - qx)) - ((qx - px) * (ry - qy))
	if val == 0 {
		return colinear
	}
	if val > 0 {
		return clockwise
	}
	return counterclockwise
}

func onSegment(px, py, qx, qy, rx, ry float64) bool {
	if (qx <= max(px, rx)) && (qx >= min(px, rx)) && (qy <= max(py, ry)) && (qy >= min(py, ry)) {
		return true
	}
	return false
}

func max(a, b float64) float64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b float64) float64 {
	if a < b {
		return a
	} else {
		return b
	}
}
