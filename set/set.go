package set

// Inspired by https://bitfieldconsulting.com/posts/generic-set :)

// Empty struct as the "value" within this hash
type Set[E comparable] map[E]struct{}

// NewSet is used to construct a new set with vals
func NewSet[E comparable](vals ...E) Set[E] {
	s := Set[E]{}
	for _, v := range vals {
		s[v] = struct{}{}
	}
	return s
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
