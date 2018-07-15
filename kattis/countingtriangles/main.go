package main

import (
	"fmt"
)

var numSegments int
var segments [50]lineSegment

var triangles map[[3]int]bool
var numTriangles int

var lowId, midId, highId int

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

				// Store intersection info
				if intersects(id0, id1) {
					segments[id0].intersections[id1] = true
					segments[id1].intersections[id0] = true
				}
			}
		}

		// Search for triangles
		triangles = make(map[[3]int]bool)
		for id0 := 0; id0 < numSegments; id0++ {
			for id1, _ := range segments[id0].intersections {
				if id0 == id1 {
					continue
				}
				for id2, _ := range segments[id1].intersections {

					if _, found := segments[id2].intersections[id0]; !found {
						continue
					}

					if id0 == id2 || id1 == id2 {
						continue
					}

					if id0 < id1 {
						if id1 < id2 {
							lowId = id0
							midId = id1
							highId = id2
						} else if id0 < id2 {
							lowId = id0
							midId = id2
							highId = id1
						} else {
							lowId = id2
							midId = id0
							highId = id1
						}
					} else {
						if id0 < id2 {
							lowId = id1
							midId = id0
							highId = id2
						} else if id2 > id1 {
							lowId = id1
							midId = id2
							highId = id0
						} else {
							lowId = id2
							midId = id1
							highId = id0
						}
					}

					if _, found := triangles[[3]int{lowId, midId, highId}]; found {
						continue
					}

					triangles[[3]int{lowId, midId, highId}] = true
					numTriangles++
				}
			}
		}

		fmt.Println(numTriangles)
	}
}

type lineSegment struct {
	id             int
	ax, ay, bx, by float64
	intersections  map[int]bool
}

// INTERSECTION JUNK

const (
	colinear = iota
	clockwise
	counterclockwise
)

func intersects(id0, id1 int) bool {
	o1 := orientation(segments[id0].ax, segments[id0].ay, segments[id0].bx, segments[id0].by, segments[id1].ax, segments[id1].ay)
	o2 := orientation(segments[id0].ax, segments[id0].ay, segments[id0].bx, segments[id0].by, segments[id1].bx, segments[id1].by)
	o3 := orientation(segments[id1].ax, segments[id1].ay, segments[id1].bx, segments[id1].by, segments[id0].ax, segments[id0].ay)
	o4 := orientation(segments[id1].ax, segments[id1].ay, segments[id1].bx, segments[id1].by, segments[id0].bx, segments[id0].by)

	if (o1 != o2) && (o3 != o4) {
		return true
	}

	if (o1 == colinear) && onSegment(segments[id0].ax, segments[id0].ay, segments[id1].ax, segments[id1].ay, segments[id0].bx, segments[id0].by) {
		return true
	}

	if (o2 == colinear) && onSegment(segments[id0].ax, segments[id0].ay, segments[id1].bx, segments[id1].by, segments[id0].bx, segments[id0].by) {
		return true
	}

	if (o3 == colinear) && onSegment(segments[id1].ax, segments[id1].ay, segments[id0].ax, segments[id0].ay, segments[id1].bx, segments[id1].by) {
		return true
	}

	if (o4 == colinear) && onSegment(segments[id1].ax, segments[id1].ay, segments[id0].bx, segments[id0].by, segments[id1].bx, segments[id1].by) {
		return true
	}

	return false
}

func orientation(px, py, qx, qy, rx, ry float64) int {
	val := ((qy - py) * (rx - qx)) - ((qx - px) * (ry - qy))
	if val == 0 {
		return colinear
	} else if val > 0 {
		return clockwise
	} else {
		return counterclockwise
	}
}

func onSegment(px, py, qx, qy, rx, ry float64) bool {

	var highpxrx, lowpxrx float64
	if px > rx {
		highpxrx = px
		lowpxrx = rx
	} else {
		highpxrx = rx
		lowpxrx = px
	}

	if qx > highpxrx {
		return false
	} else if qx < lowpxrx {
		return false
	}

	var highpyry, lowpyry float64
	if py > ry {
		highpyry = py
		lowpyry = ry
	} else {
		highpyry = ry
		lowpyry = py
	}

	if qy > highpyry {
		return false
	} else if qy < lowpyry {
		return false
	}

	return true

}
