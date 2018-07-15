package main

import (
	"bufio"
	"fmt"
	"os"
)

type identifier byte
type real float32

type lineSegment struct {
	id             identifier
	ax, ay, bx, by real
	intersections  map[identifier]bool
}

type triangle [3]identifier

var numSegments identifier
var segments [50]lineSegment

var aTriangle triangle
var triangles map[triangle]bool
var numTriangles int

var lowId, midId, highId identifier
var o1, o2, o3, o4 int
var val real

var id0, id1, id2 identifier

var scanner *bufio.Scanner

var found bool

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	for {
		// Set initial variables
		numTriangles = 0

		// Get initial input
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &numSegments)
		if numSegments == 0 {
			return
		}

		// Read line segments
		for id0 = 0; id0 < numSegments; id0++ {
			segments[id0] = lineSegment{id: id0, intersections: make(map[identifier]bool)}
			scanner.Scan()
			fmt.Sscanf(scanner.Text(), "%f %f %f %f", &segments[id0].ax, &segments[id0].ay, &segments[id0].bx, &segments[id0].by)
		}

		// Gather intersection information
		for id0 = 0; id0 < numSegments; id0++ {
			for id1 = 0; id1 < numSegments; id1++ {
				if id0 == id1 {
					continue
				}

				if _, found = segments[id0].intersections[id1]; found {
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
		triangles = make(map[triangle]bool)
		for id0 = 0; id0 < numSegments; id0++ {
			for id1, _ = range segments[id0].intersections {
				if id0 == id1 {
					continue
				}
				for id2, _ = range segments[id1].intersections {
					if id0 == id2 {
						continue
					}

					// doesn't form a triangle
					if _, found = segments[id0].intersections[id2]; !found {
						continue
					}

					// sort ids
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

					aTriangle = triangle{lowId, midId, highId}

					if _, found = triangles[aTriangle]; found {
						continue
					}

					triangles[aTriangle] = true
					numTriangles++
				}
			}
		}

		fmt.Println(numTriangles)
	}
}

// INTERSECTION JUNK

func intersects(id0, id1 identifier) bool {
	o1 = orientation(segments[id0].ax, segments[id0].ay, segments[id0].bx, segments[id0].by, segments[id1].ax, segments[id1].ay)
	o2 = orientation(segments[id0].ax, segments[id0].ay, segments[id0].bx, segments[id0].by, segments[id1].bx, segments[id1].by)
	o3 = orientation(segments[id1].ax, segments[id1].ay, segments[id1].bx, segments[id1].by, segments[id0].ax, segments[id0].ay)
	o4 = orientation(segments[id1].ax, segments[id1].ay, segments[id1].bx, segments[id1].by, segments[id0].bx, segments[id0].by)
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

const (
	colinear = iota
	clockwise
	counterclockwise
)

func orientation(px, py, qx, qy, rx, ry real) int {
	val = ((qy - py) * (rx - qx)) - ((qx - px) * (ry - qy))
	if val == 0 {
		return colinear
	} else if val > 0 {
		return clockwise
	} else {
		return counterclockwise
	}
}

func onSegment(px, py, qx, qy, rx, ry real) bool {
	if px > rx {
		if qx > px {
			return false
		} else if qx < rx {
			return false
		}
	} else {
		if qx > rx {
			return false
		} else if qx < px {
			return false
		}
	}
	if py > ry {
		if qy > py {
			return false
		} else if qy < ry {
			return false
		}
	} else {
		if qy > ry {
			return false
		} else if qy < py {
			return false
		}
	}
	return true
}
