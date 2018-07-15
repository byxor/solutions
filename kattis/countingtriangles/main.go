package main

import (
	"bufio"
	"fmt"
	"os"
)

type identifier byte
type real float32
type ot byte

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
var theOrientation ot
var o1, o2, o3, o4 ot
var val real

var px, py, qx, qy, rx, ry real

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
				if intersects() {
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

func intersects() bool {
	px = segments[id0].ax
	py = segments[id0].ay
	qx = segments[id0].bx
	qy = segments[id0].by
	rx = segments[id1].ax
	ry = segments[id1].ay
	orientation()
	o1 = theOrientation

	px = segments[id0].ax
	py = segments[id0].ay
	qx = segments[id0].bx
	qy = segments[id0].by
	rx = segments[id1].bx
	ry = segments[id1].by
	orientation()
	o2 = theOrientation

	px = segments[id1].ax
	py = segments[id1].ay
	qx = segments[id1].bx
	qy = segments[id1].by
	rx = segments[id0].ax
	ry = segments[id0].ay
	orientation()
	o3 = theOrientation

	px = segments[id1].ax
	py = segments[id1].ay
	qx = segments[id1].bx
	qy = segments[id1].by
	rx = segments[id0].bx
	ry = segments[id0].by
	orientation()
	o4 = theOrientation

	if (o1 != o2) && (o3 != o4) {
		return true
	}

	px = segments[id0].ax
	py = segments[id0].ay
	qx = segments[id1].ax
	qy = segments[id1].ay
	rx = segments[id0].bx
	ry = segments[id0].by
	if (o1 == colinear) && onSegment() {
		return true
	}

	px = segments[id0].ax
	py = segments[id0].ay
	qx = segments[id1].bx
	qy = segments[id1].by
	rx = segments[id0].bx
	ry = segments[id0].by
	if (o2 == colinear) && onSegment() {
		return true
	}

	px = segments[id1].ax
	py = segments[id1].ay
	qx = segments[id0].ax
	qy = segments[id0].ay
	rx = segments[id1].bx
	ry = segments[id1].by
	if (o3 == colinear) && onSegment() {
		return true
	}

	px = segments[id1].ax
	py = segments[id1].ay
	qx = segments[id0].bx
	qy = segments[id0].by
	rx = segments[id1].bx
	ry = segments[id1].by
	if (o4 == colinear) && onSegment() {
		return true
	}
	return false
}

const (
	colinear = iota
	clockwise
	counterclockwise
)

func orientation() {
	val = ((qy - py) * (rx - qx)) - ((qx - px) * (ry - qy))
	if val == 0 {
		theOrientation = colinear
		return
	} else if val > 0 {
		theOrientation = clockwise
		return
	} else {
		theOrientation = counterclockwise
		return
	}
}

func onSegment() bool {
	if px <= rx {
		if qx > rx {
			return false
		} else if qx < px {
			return false
		}
	} else {
		if qx > px {
			return false
		} else if qx < rx {
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
