// Code generated by goderive DO NOT EDIT.

package intern

import (
	ast "github.com/katydid/katydid/relapse/ast"
)

// deriveTraverse returns a list where each element of the input list has been morphed by the input function or an error.
func deriveTraverse(f func(*ast.Pattern) (*Pattern, error), list []*ast.Pattern) ([]*Pattern, error) {
	out := make([]*Pattern, len(list))
	var err error
	for i, elem := range list {
		out[i], err = f(elem)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

// deriveTraverseAst returns a list where each element of the input list has been morphed by the input function or an error.
func deriveTraverseAst(f func(*Pattern) (*ast.Pattern, error), list []*Pattern) ([]*ast.Pattern, error) {
	out := make([]*ast.Pattern, len(list))
	var err error
	for i, elem := range list {
		out[i], err = f(elem)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

// deriveFilter returns a list of all items in the list that matches the predicate.
func deriveFilter(predicate func(*Pattern) bool, list []*Pattern) []*Pattern {
	j := 0
	for i, elem := range list {
		if predicate(elem) {
			if i != j {
				list[j] = list[i]
			}
			j++
		}
	}
	return list[:j]
}

// deriveEquals returns whether this and that are equal.
func deriveEquals(this, that []*Pattern) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i].Equal(that[i])) {
			return false
		}
	}
	return true
}

// deriveHashString returns the hash of the object.
func deriveHashString(object string) uint64 {
	var h uint64
	for _, c := range object {
		h = 31*h + uint64(c)
	}
	return h
}

// deriveAny reports whether the predicate returns true for any of the elements in the given slice.
func deriveAny(pred func(*Pattern) bool, list []*Pattern) bool {
	for _, elem := range list {
		if pred(elem) {
			return true
		}
	}
	return false
}

// deriveAll reports whether the predicate returns true for all of the elements in the given slice.
func deriveAll(predicate func(*Pattern) bool, slice []*Pattern) bool {
	for _, elem := range slice {
		if !predicate(elem) {
			return false
		}
	}
	return true
}
