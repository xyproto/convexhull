package convexhull

import (
	"fmt"
)

func ExampleCompute() {
	ps := Points{
		Point{0, 0},
		Point{1, 2},
		Point{3, 4},
	}
	newPoints, err := ps.Compute()
	if err != nil {
		panic(err)
	}
	fmt.Println(ps)
	fmt.Println(newPoints)
	// Output:
	// [{0 0} {3 4} {1 2}]
	// [{1 2} {3 4} {0 0}]
}
