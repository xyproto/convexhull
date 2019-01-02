package convexhull

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type Point struct {
	X, Y float64
}

type Points []*Point

func New(x, y float64) *Point {
	return &Point{X: x, Y: y}
}

func isLeft(p0, p1, p2 *Point) bool {
	return Area2(p0, p1, p2) > 0
}

func Area2(a, b, c *Point) float64 {
	return (b.X-a.X)*(c.Y-a.Y) - (c.X-a.X)*(b.Y-a.Y)
}

//Implement sort interface
func (ps Points) Len() int {
	return len(ps)
}

func (ps Points) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps Points) Less(i, j int) bool {
	area := Area2(ps[0], ps[i], ps[j])

	if area == 0 {
		x := math.Abs(ps[i].X-ps[0].X) - math.Abs(ps[j].X-ps[0].X)
		y := math.Abs(ps[i].Y-ps[0].Y) - math.Abs(ps[j].Y-ps[0].Y)

		if x < 0 || y < 0 {
			return true
		}

		//} else if x > 0 || y > 0 {
		//	return false
		//} else {
		//	return false
		//}

		return false
	}

	return area > 0
}

func (ps Points) Lowest() {
	m := 0
	for i := 1; i < len(ps); i++ {
		//If lowest points are on the same line, take the rightmost point
		if (ps[i].Y < ps[m].Y) || ((ps[i].Y == ps[m].Y) && ps[i].X > ps[m].X) {
			m = i
		}
	}
	ps[0], ps[m] = ps[m], ps[0]
}

func (ps Points) Compute() (Points, error) {
	if len(ps) < 3 {
		return nil, nil
	}

	stack := &PointStack{}

	ps.Lowest()
	sort.Sort(&ps)

	stack.Push(ps[0])
	stack.Push(ps[1])

	//fmt.Printf("Sorted Points: %v\n", ps)

	i := 2
	for i < len(ps) {
		pi := ps[i]

		//PrintStack(stack)

		p1 := stack.top.next.value
		p2 := stack.top.value

		if isLeft(p1, p2, pi) {
			stack.Push(pi)
			i++
		} else {
			stack.Pop()
		}
	}

	// Copy the hull
	ret := make(Points, stack.Len())
	top := stack.top

	var count int
	for top != nil {
		ret[count] = top.value
		top = top.next
		count++
	}

	return ret, nil
}

func (ps Points) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for i, p := range ps {
		if i > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(fmt.Sprintf("{%v %v}", p.X, p.Y))
	}
	sb.WriteString("]")
	return sb.String()
}
