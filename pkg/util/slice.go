package util

import (
	"fmt"
	"math/rand"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

func SliceDiff[T constraints.Ordered](old, new []T) (added, removed, kept []T) {
	return SliceDiffFunc(old, new, func(a, b T) int {
		if a == b {
			return 0
		} else if a < b {
			return -1
		} else {
			return 1
		}
	})
}

func SliceDiffFunc[T any](old, new []T, cmp func(a, b T) int) (added, removed, kept []T) {
	added = []T{}
	removed = []T{}
	kept = []T{}

	less := func(a, b T) bool {
		return cmp(a, b) < 0
	}

	slices.SortFunc(new, less)
	slices.SortFunc(old, less)

	i, j := 0, 0
	for i < len(old) && j < len(new) {
		c := cmp(old[i], new[j])
		switch {
		case c < 0: // removed
			removed = append(removed, old[i])
			i++

		case c > 0: // added
			added = append(added, new[j])
			j++

		default: // kept
			kept = append(kept, new[j])
			i++
			j++
		}
	}

	// Add rest

	for ; i < len(old); i++ {
		removed = append(removed, old[i])
	}

	for ; j < len(new); j++ {
		added = append(added, new[j])
	}

	return
}

func SliceShuffle[T any](s []T) {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

func SliceFilter[T any](s []T, cmp func(T) bool) []T {
	t := []T{}

	for _, i := range s {
		if cmp(i) {
			t = append(t, i)
		}
	}

	return t
}

func SliceContains[T any](s []T, cmp func(T) bool) bool {
	for _, i := range s {
		if cmp(i) {
			return true
		}
	}

	return false
}

func SliceMap[T any](s []T, cb func(T) T) []T {
	n := []T{}

	for _, t := range s {
		n = append(n, cb(t))
	}

	return n
}

func SliceString[T any](s []T) []string {
	n := []string{}

	for _, t := range s {
		n = append(n, fmt.Sprintf("%v", t))
	}

	return n
}
