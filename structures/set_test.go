package structures

import (
	"reflect"
	"testing"
)

func TestNewSet(t *testing.T) {
	t.Run("NewSet", func(t *testing.T) {
		s := NewSet()
		if s == nil {
			t.Error("NewSet() = nil, want non-nil")
		}
	})
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		s    Set
		v    string
	}{
		{
			name: "Add foo",
			s:    NewSet(),
			v:    "foo",
		},
		{
			name: "Add empty string",
			s:    NewSet(),
			v:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Add(tt.v)
			if !tt.s.Contains(tt.v) {
				t.Errorf("Set.Add(%q) failed", tt.v)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name string
		s    Set
		v    string
	}{
		{
			name: "Remove foo",
			s:    NewSet(),
			v:    "foo",
		},
		{
			name: "Remove empty string",
			s:    NewSet(),
			v:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Add(tt.v)
			tt.s.Remove(tt.v)
			if tt.s.Contains(tt.v) {
				t.Errorf("Set.Remove(%q) failed", tt.v)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name string
		s    Set
		v    string
		add  bool
		want bool
	}{
		{
			name: "Contains foo",
			s:    NewSet(),
			v:    "foo",
			add:  true,
			want: true,
		},
		{
			name: "Contains empty string",
			s:    NewSet(),
			v:    "",
			add:  true,
			want: true,
		},
		{
			name: "Does not contain bar",
			s:    NewSet(),
			v:    "bar",
			add:  false,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.add {
				tt.s.Add(tt.v)
			}
			if got := tt.s.Contains(tt.v); got != tt.want {
				t.Errorf("Set.Contains(%q) = %v, want %v", tt.v, got, tt.want)
			}
		})
	}
}

func TestSize(t *testing.T) {
	tests := []struct {
		name   string
		s      Set
		values []string
		want   int
	}{
		{
			name:   "Size of empty set",
			s:      NewSet(),
			values: []string{},
			want:   0,
		},
		{
			name:   "Size of set with one element",
			s:      NewSet(),
			values: []string{"foo"},
			want:   1,
		},
		{
			name:   "Size of set with two elements",
			s:      NewSet(),
			values: []string{"foo", "bar"},
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, v := range tt.values {
				tt.s.Add(v)
			}
			if got := tt.s.Size(); got != tt.want {
				t.Errorf("Set.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		name  string
		s1    Set
		v1    []string
		s2    Set
		v2    []string
		want  Set
		vwant []string
	}{
		{
			name:  "Intersection of empty sets",
			s1:    NewSet(),
			v1:    []string{},
			s2:    NewSet(),
			v2:    []string{},
			want:  NewSet(),
			vwant: []string{},
		},
		{
			name:  "Intersection of empty set and non-empty set",
			s1:    NewSet(),
			v1:    []string{},
			s2:    NewSet(),
			v2:    []string{"foo"},
			want:  NewSet(),
			vwant: []string{},
		},
		{
			name:  "Intersection of non-empty set and empty set",
			s1:    NewSet(),
			v1:    []string{"foo"},
			s2:    NewSet(),
			v2:    []string{},
			want:  NewSet(),
			vwant: []string{},
		},
		{
			name:  "Intersection of non-empty sets with no common elements",
			s1:    NewSet(),
			v1:    []string{"foo"},
			s2:    NewSet(),
			v2:    []string{"bar"},
			want:  NewSet(),
			vwant: []string{},
		},
		{
			name:  "Intersection of non-empty sets with one common element",
			s1:    NewSet(),
			v1:    []string{"foo"},
			s2:    NewSet(),
			v2:    []string{"foo"},
			want:  NewSet(),
			vwant: []string{"foo"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, v := range tt.v1 {
				tt.s1.Add(v)
			}
			for _, v := range tt.v2 {
				tt.s2.Add(v)
			}
			for _, v := range tt.vwant {
				tt.want.Add(v)
			}
			if got := tt.s1.Intersection(tt.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set.Intersection(%v) = %v, want %v", tt.s2, got, tt.want)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		name  string
		s1    Set
		v1    []string
		s2    Set
		v2    []string
		want  Set
		vwant []string
	}{
		{
			name:  "Union of empty sets",
			s1:    NewSet(),
			v1:    []string{},
			s2:    NewSet(),
			v2:    []string{},
			want:  NewSet(),
			vwant: []string{},
		},
		{
			name:  "Union of empty set and non-empty set",
			s1:    NewSet(),
			v1:    []string{},
			s2:    NewSet(),
			v2:    []string{"foo"},
			want:  NewSet(),
			vwant: []string{"foo"},
		},
		{
			name:  "Union of non-empty set and empty set",
			s1:    NewSet(),
			v1:    []string{"foo"},
			s2:    NewSet(),
			v2:    []string{},
			want:  NewSet(),
			vwant: []string{"foo"},
		},
		{
			name:  "Intersection of non-empty sets with no common elements",
			s1:    NewSet(),
			v1:    []string{"foo"},
			s2:    NewSet(),
			v2:    []string{"bar"},
			want:  NewSet(),
			vwant: []string{"foo", "bar"},
		},
		{
			name:  "Intersection of non-empty sets with one common element",
			s1:    NewSet(),
			v1:    []string{"foo"},
			s2:    NewSet(),
			v2:    []string{"foo"},
			want:  NewSet(),
			vwant: []string{"foo"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, v := range tt.v1 {
				tt.s1.Add(v)
			}
			for _, v := range tt.v2 {
				tt.s2.Add(v)
			}
			for _, v := range tt.vwant {
				tt.want.Add(v)
			}
			if got := tt.s1.Union(tt.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set.Union(%v) = %v, want %v", tt.s2, got, tt.want)
			}
		})
	}
}
