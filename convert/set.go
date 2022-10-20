package convert

import "github.com/code-tbd/gox/set"

func Slice2Set[T comparable](slice []T) set.Set[T] {
	return set.Slice2Set(slice)
}
