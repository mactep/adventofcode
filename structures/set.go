package structures

// Set is a set of strings.
type Set map[string]struct{}

// NewSet creates a new set.
func NewSet() Set {
	return make(Set)
}

// Add adds a string to the set.
func (s Set) Add(str string) {
	s[str] = struct{}{}
}

// Remove removes a string from the set.
func (s Set) Remove(str string) {
	delete(s, str)
}

// Contains returns true if the set contains the string.
func (s Set) Contains(str string) bool {
	_, ok := s[str]
	return ok
}

// Size returns the number of strings in the set.
func (s Set) Size() int {
	return len(s)
}

// Intersection returns the intersection of two sets.
func (s Set) Intersection(t Set) Set {
	u := NewSet()
	for str := range s {
		if t.Contains(str) {
			u.Add(str)
		}
	}
	return u
}

// Union returns the union of two sets.
func (s Set) Union(t Set) Set {
	u := NewSet()
	for str := range s {
		u.Add(str)
	}
	for str := range t {
		u.Add(str)
	}
	return u
}
