package convexhull

// Based on the Stack implementation from https://gist.github.com/bemasher/1777766

import (
	"errors"
	"fmt"
)

type PointStack struct {
	top  *Element
	size int
}

type Element struct {
	value Point
	next  *Element
}

// Return the stack's length
func (s *PointStack) Len() int {
	return s.size
}

// Push a new element onto the stack
func (s *PointStack) Push(value Point) {
	s.top = &Element{value, s.top}
	s.size++
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *PointStack) Pop() (Point, error) {
	var value Point
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return value, nil
	}
	return value, errors.New("empty")
}

func PrintPointStack(s *PointStack) {
	v := s.top
	fmt.Printf("PointStack: ")
	for v != nil {
		fmt.Printf("%v ", v.value)
		v = v.next
	}
	fmt.Println("")
}
