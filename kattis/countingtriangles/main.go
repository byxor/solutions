package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const maxSegments = 50

type identifier byte
type real float32
type ot byte

type lineSegment struct {
	id identifier
}

type triangle [3]identifier

var foo int64
var numSegments identifier
var segments [maxSegments]lineSegment
var intersections [maxSegments]map[identifier]byte
var axs [maxSegments]real
var ays [maxSegments]real
var bxs [maxSegments]real
var bys [maxSegments]real

var aTriangle triangle
var triangles map[triangle]byte
var numTriangles uint16

var lowId, midId, highId identifier
var theOrientation ot
var o1, o2, o3, o4 ot
var val real

var px, py, qx, qy, rx, ry real

var id0, id1, id2 identifier

var itIntersects bool
var itOnSegment bool

var scanner *bufio.Scanner

var found bool

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	for {
		// Set initial variables
		numTriangles = 0

		// Get initial input
		scanner.Scan()
		if scanner.Text() == "0" {
			return
		}

		foo, _ = strconv.ParseInt(scanner.Text(), 10, 8)
		numSegments = identifier(foo)

		// Read line segments
		for id0 = 0; id0 < numSegments; id0++ {
			segments[id0] = lineSegment{id: id0}
			intersections[id0] = make(map[identifier]byte)
			scanner.Scan()
			fmt.Sscanf(scanner.Text(), "%f %f %f %f", &axs[id0], &ays[id0], &bxs[id0], &bys[id0])
		}

		// Gather intersection information
		for id0 = 0; id0 < numSegments; id0++ {
			for id1 = 0; id1 < numSegments; id1++ {
				if id0 == id1 {
					continue
				}

				if _, found = intersections[id0][id1]; found {
					continue
				}

				// Store intersection info
				intersects()
				if itIntersects {
					intersections[id0][id1] = 0
					intersections[id1][id0] = 0
				}
			}
		}

		// Search for triangles
		triangles = make(map[triangle]byte)
		for id0 = 0; id0 < numSegments; id0++ {
			for id1, _ = range intersections[id0] {
				if id0 == id1 {
					continue
				}
				for id2, _ = range intersections[id1] {
					if id0 == id2 {
						continue
					}

					// doesn't form a triangle
					if _, found = intersections[id0][id2]; !found {
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

					triangles[aTriangle] = 0
					numTriangles++
				}
			}
		}

		fmt.Println(numTriangles)
	}
}

// INTERSECTION JUNK

func intersects() {
	px = axs[id0]
	py = ays[id0]
	qx = bxs[id0]
	qy = bys[id0]
	rx = axs[id1]
	ry = ays[id1]
	orientation()
	o1 = theOrientation

	px = axs[id0]
	py = ays[id0]
	qx = bxs[id0]
	qy = bys[id0]
	rx = bxs[id1]
	ry = bys[id1]
	orientation()
	o2 = theOrientation

	px = axs[id1]
	py = ays[id1]
	qx = bxs[id1]
	qy = bys[id1]
	rx = axs[id0]
	ry = ays[id0]
	orientation()
	o3 = theOrientation

	px = axs[id1]
	py = ays[id1]
	qx = bxs[id1]
	qy = bys[id1]
	rx = bxs[id0]
	ry = bys[id0]
	orientation()
	o4 = theOrientation

	if (o1 != o2) && (o3 != o4) {
		itIntersects = true
		return
	}

	px = axs[id0]
	py = ays[id0]
	qx = axs[id1]
	qy = ays[id1]
	rx = bxs[id0]
	ry = bys[id0]
	if o1 == colinear {
		onSegment()
		if itOnSegment {
			itIntersects = true
			return
		}
	}

	px = axs[id0]
	py = ays[id0]
	qx = bxs[id1]
	qy = bys[id1]
	rx = bxs[id0]
	ry = bys[id0]
	if o2 == colinear {
		onSegment()
		if itOnSegment {
			itIntersects = true
			return
		}
	}

	px = axs[id1]
	py = ays[id1]
	qx = axs[id0]
	qy = ays[id0]
	rx = bxs[id1]
	ry = bys[id1]
	if o3 == colinear {
		onSegment()
		if itOnSegment {
			itIntersects = true
			return
		}
	}

	px = axs[id1]
	py = ays[id1]
	qx = bxs[id0]
	qy = bys[id0]
	rx = bxs[id1]
	ry = bys[id1]
	if o4 == colinear {
		onSegment()
		if itOnSegment {
			itIntersects = true
			return
		}
	}

	itIntersects = false
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

func onSegment() {
	if py > ry {
		if qy > py {
			itOnSegment = false
		} else if qy < ry {
			itOnSegment = false
		}
	} else {
		if qy > ry {
			itOnSegment = false
		} else if qy < py {
			itOnSegment = false
		}
	}
	if px <= rx {
		if qx > rx {
			itOnSegment = false
		} else if qx < px {
			itOnSegment = false
		}
	} else {
		if qx > px {
			itOnSegment = false
		} else if qx < rx {
			itOnSegment = false
		}
	}
	itOnSegment = true
}
