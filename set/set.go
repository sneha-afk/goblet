// Package set defines simple set utilities, using an underlying map.
package set

// Inspired by https://bitfieldconsulting.com/posts/generic-set :)

// Empty struct as the "value" within this hash
type Set[E comparable] map[E]struct{}

// NewSet is used to construct a new set with vals
func NewSet[E comparable](vals ...E) Set[E] {
	s := make(Set[E], len(vals))
	for _, v := range vals {
		s[v] = struct{}{}
	}
	return s
}

// Clone this set, making a deep, new copy
func (s Set[E]) Clone() Set[E] {
	out := make(Set[E], len(s))
	for v := range s {
		out[v] = struct{}{}
	}
	return out
}

// Count of elements in this set
func (s Set[E]) Count() int {
	return len(s)
}

func (s Set[E]) IsEmpty() bool {
	return len(s) == 0
}

// Clear this set of all elements; note memory is kept allocated.
func (s Set[E]) Clear() {
	s.Clear()
}

func (s Set[E]) ToSlice() []E {
	vals := []E{}

	for v, _ := range s {
		vals = append(vals, v)
	}
	return vals
}

// Add v to set s
func (s Set[E]) Add(v E) {
	s[v] = struct{}{}
}

// Delete v from s; no-op if not in s
func (s Set[E]) Delete(v E) {
	delete(s, v)
}

// Contains checks if v is in the set
func (s Set[E]) Contains(v E) bool {
	_, ok := s[v]
	return ok
}

// Union this set with another, returning a new set
func (s Set[E]) Union(o Set[E]) Set[E] {
	union := s.Clone()

	// Harmless (I suppose) overwrite of shared elements
	for oe, _ := range o {
		union[oe] = struct{}{}
	}
	return union
}

// Intersection of this set with another, returning a new set
func (s Set[E]) Intersection(o Set[E]) Set[E] {
	intersection := []E{}

	for oe, _ := range o {
		if s.Contains(oe) {
			intersection = append(intersection, oe)
		}
	}
	return NewSet(intersection...)
}

// Difference of this set and another, i.e this - o, returning a new set
func (s Set[E]) Difference(o Set[E]) Set[E] {
	diff := s.Clone()

	for oe, _ := range o {
		delete(diff, oe)
	}
	return diff
}

// SymmetricDifference of this set and another, i.e this ^ o, returning a new set.
// This new set is of elements in exactly ONE set, but not both.
func (s Set[E]) SymmetricDifference(o Set[E]) Set[E] {
	symdiff := s.Clone()

	for oe, _ := range o {
		if symdiff.Contains(oe) {
			delete(symdiff, oe)
		}
	}
	return symdiff
}

// IsSubset returns true if this is a subset of o; i.e all elements in s
// is contained within o.
func (s Set[E]) IsSubset(o Set[E]) bool {
	if s.Count() > o.Count() {
		return false
	}

	for se, _ := range s {
		if !o.Contains(se) {
			return false
		}
	}
	return true
}

// IsSuperset returns true if this is a superset of o, i.e this contains all
// elements of o.
func (s Set[E]) IsSuperset(o Set[E]) bool {
	if s.Count() < o.Count() {
		return false
	}

	for oe, _ := range o {
		if !s.Contains(oe) {
			return false
		}
	}
	return true
}

// IsEqual returns true if this and the other set have the same count, and same
// values for all elements.
func (s Set[E]) IsEqual(o Set[E]) bool {
	if s.Count() != o.Count() {
		return false
	}

	for oe, _ := range o {
		if !s.Contains(oe) {
			return false
		}
	}
	return true
}
