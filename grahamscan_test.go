package convexhull

import (
	"fmt"
)

func Example() {
	ps := Points{
		&Point{0, 0},
		&Point{1, 2},
		&Point{3, 4},
		&Point{-4, 5},
		&Point{20, 70},
	}
	newPoints, err := ps.Compute()
	if err != nil {
		panic(err)
	}
	fmt.Println(ps)
	fmt.Println(newPoints)
	// Output:
	// [{0 0} {3 4} {1 2} {20 70} {-4 5}]
	// [{-4 5} {20 70} {3 4} {0 0}]
}
