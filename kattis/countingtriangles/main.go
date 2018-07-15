package main

import (
	"fmt"
)

type point struct {
	x, y float64
}

type lineSegment struct {
	id            int
	pointA        point
	pointB        point
	intersections map[int]bool
}

type triangle struct {
	a, b, c int
}

var numSegments int
var segments []lineSegment

var segboi lineSegment

func main() {
	for {
		fmt.Scan(&numSegments)

		if numSegments == 0 {
			return
		}

		segments = make([]lineSegment, numSegments)

		for i := 0; i < numSegments; i++ {
			segboi = lineSegment{}
			segboi.id = i
			segboi.intersections = make(map[int]bool)
			fmt.Scanf("%f %f %f %f\n", &segboi.pointA.x, &segboi.pointA.y, &segboi.pointB.x, &segboi.pointB.y)
			segments[i] = segboi
		}

		for _, segmentA := range segments {
			for _, segmentB := range segments {
				if segmentA.id == segmentB.id {
					continue
				}

				if segmentA.intersects(&segmentB) {
					continue
				}

				if intersects(&segmentA, &segmentB) {
					segmentA.intersections[segmentB.id] = true
					segmentB.intersections[segmentA.id] = true
				}
			}
		}

		triangles := make(map[triangle]bool)
		numTriangles := 0
		for id0 := 0; id0 < numSegments; id0++ {
			segment := segments[id0]
			for id1, _ := range segment.intersections {
				for id2, _ := range segments[id1].intersections {
					if !segments[id2].intersects(&segment) {
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

func (ls *lineSegment) intersects(other *lineSegment) bool {
	_, found := ls.intersections[other.id]
	return found
}

// INTERSECTION JUNK

const (
	colinear = iota
	clockwise
	counterclockwise
)

func intersects(segmentA, segmentB *lineSegment) bool {
	o1 := orientation(segmentA.pointA, segmentA.pointB, segmentB.pointA)
	o2 := orientation(segmentA.pointA, segmentA.pointB, segmentB.pointB)
	o3 := orientation(segmentB.pointA, segmentB.pointB, segmentA.pointA)
	o4 := orientation(segmentB.pointA, segmentB.pointB, segmentA.pointB)

	if (o1 != o2) && (o3 != o4) {
		return true
	}

	if (o1 == colinear) && onSegment(segmentA.pointA, segmentB.pointA, segmentA.pointB) {
		return true
	}

	if (o2 == colinear) && onSegment(segmentA.pointA, segmentB.pointB, segmentA.pointB) {
		return true
	}

	if (o3 == colinear) && onSegment(segmentB.pointA, segmentA.pointA, segmentB.pointB) {
		return true
	}

	if (o4 == colinear) && onSegment(segmentB.pointA, segmentA.pointB, segmentB.pointB) {
		return true
	}

	return false
}

func orientation(p, q, r point) int {
	val := ((q.y - p.y) * (r.x - q.x)) - ((q.x - p.x) * (r.y - q.y))
	if val == 0 {
		return colinear
	}
	if val > 0 {
		return clockwise
	}
	return counterclockwise
}

func onSegment(p, q, r point) bool {
	if (q.x <= max(p.x, r.x)) && (q.x >= min(p.x, r.x)) && (q.y <= max(p.y, r.y)) && (q.y >= min(p.y, r.y)) {
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
