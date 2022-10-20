package set

func Slice2Set[T comparable](slice []T) Set[T] {
	set := NewSet[T]()
	for _, value := range slice {
		set.Put(value)
	}

	return set
}
