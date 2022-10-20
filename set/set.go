package set

import "errors"

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return Set[T]{}
}

func (s Set[T]) Put(value T) {
	s[value] = struct{}{}
}

// PutNE put if not exist
func (s Set[T]) PutNE(value T) error {
	if s.Exist(value) {
		return errors.New("value already existed")
	}

	s[value] = struct{}{}
	return nil
}

func (s Set[T]) Del(value T) {
	delete(s, value)
}

// DelE delete if exist
func (s Set[T]) DelE(value T) error {
	if !s.Exist(value) {
		return errors.New("value not existed")
	}

	delete(s, value)
	return nil
}

func (s Set[T]) Exist(value T) bool {
	if _, ok := s[value]; ok {
		return true
	}

	return false
}

func (s Set[T]) Slice() []T {
	slice := make([]T, len(s))
	for value, _ := range s {
		slice = append(slice, value)
	}
	return slice
}

func (s Set[T]) GetAll() []T {
	return s.Slice()
}
