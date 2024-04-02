package geometric

import (
	"math"
	"sort"
)

type Point2D struct {
	x float64
	y float64
}

type Line struct {
	p1 Point2D
	p2 Point2D
}

type PointSet = []Point2D

// Intersecting Lines

func onOppositeSides(a Point2D, b Point2D, l Line) bool {
	g := (l.p2.x-l.p1.x)*(a.y-l.p1.y) - (l.p2.y-l.p1.y)*(a.x-l.p1.x)

	h := (l.p2.x-l.p1.x)*(b.y-l.p1.y) - (l.p2.y-l.p1.y)*(b.x-l.p1.x)

	return g*h <= 0
}

func boundingBox(l1 Line, l2 Line) bool {
	l1Left := min(l1.p1.x, l1.p2.x)
	l1Right := max(l1.p1.x, l1.p2.x)
	l2Left := min(l2.p1.x, l2.p2.x)
	l2Right := min(l2.p1.x, l2.p2.x)

	l1Bottom := min(l1.p1.y, l1.p2.y)
	l1Top := max(l1.p1.y, l1.p2.y)
	l2Bottom := min(l2.p1.y, l2.p2.y)
	l2Top := max(l2.p1.y, l2.p2.y)

	xOverlap := (l1Top > l2Bottom && l2Top > l1Bottom)
	yOverlap := (l1Right > l2Left && l2Right > l1Left)

	return xOverlap && yOverlap
}

func intersection(l1 Line, l2 Line) bool {
	pqOpposite := onOppositeSides(l1.p1, l1.p2, l2)
	rsOpposite := onOppositeSides(l2.p1, l2.p2, l1)
	bounding := boundingBox(l1, l2)

	return pqOpposite && rsOpposite && bounding
}

type ByX PointSet

func (a ByX) Len() int      { return len(a) }
func (a ByX) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByX) Less(i, j int) bool {
	if a[i].x == a[j].x {
		return a[i].y < a[j].y
	}
	return a[i].x < a[j].x
}

func simplePolygon(ps PointSet) {
	angle := make([]float64, len(ps))
	sort.Sort(ByX(ps))

	for i := 1; i < len(ps); i++ {
		num := ps[0].x - ps[i].x
		denom := ps[i].y - ps[0].y

		if denom > 0 {
			angle[i] = math.Atan(num / denom)
		} else if denom == 0 {
			angle[i] = math.Pi / 2
		} else {
			angle[i] = math.Atan(num/denom) + math.Pi
		}
	}

	//Sort ps[1..n-1] by angle[1..n-1]
	sort.Slice(ps[1:], func(i, j int) bool {
		return angle[i+1] < angle[j+1]
	})
}
