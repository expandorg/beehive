package matcher

import (
	"fmt"
)

const EXACT = "exact"

type Type struct {
	Match string
}

type Response struct {
	Value interface{}
}

type Solution struct {
	Type  Type
	Value interface{}
}

type Comparable struct {
	Response Response
	Solution Solution
}

func Match(comp Comparable) (bool, error) {
	switch comp.Solution.Type.Match {
	case EXACT:
		return exact(comp), nil
	default:
		return false, NoMatcherError{comp.Solution.Type}
	}

}

func exact(comp Comparable) bool {
	return comp.Response.Value == comp.Solution.Value
}

type NoMatcherError struct {
	Type Type
}

func (e NoMatcherError) Error() string {
	return fmt.Sprintf("No matcher found for Type: %s", e.Type.Match)
}
