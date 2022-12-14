package structures

import (
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	tests := []struct {
		name string
		want Stack
	}{
		{
			name: "New stack",
			want: Stack{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPush(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		s    *Stack
		args args
		want Stack
	}{
		{
			name: "Push to empty stack",
			s:    &Stack{},
			args: args{v: 1},
			want: Stack{1},
		},
		{
			name: "Push to non-empty stack",
			s:    &Stack{1},
			args: args{v: 2},
			want: Stack{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Push(tt.args.v)
			if !reflect.DeepEqual(*tt.s, tt.want) {
				t.Errorf("Push() = %v, want %v", *tt.s, tt.want)
			}
		})
	}
}

func TestPushBottom(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		s    *Stack
		args args
		want Stack
	}{
		{
			name: "Push to empty stack",
			s:    &Stack{},
			args: args{v: 1},
			want: Stack{1},
		},
		{
			name: "Push to non-empty stack",
			s:    &Stack{1},
			args: args{v: 2},
			want: Stack{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.PushBottom(tt.args.v)
			if !reflect.DeepEqual(*tt.s, tt.want) {
				t.Errorf("PushBottom() = %v, want %v", *tt.s, tt.want)
			}
		})
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		name string
		s    *Stack
		want interface{}
	}{
		{
			name: "Pop from empty stack",
			s:    &Stack{},
			want: nil,
		},
		{
			name: "Pop from non-empty stack",
			s:    &Stack{1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeek(t *testing.T) {
	tests := []struct {
		name string
		s    *Stack
		want interface{}
	}{
		{
			name: "Peek from empty stack",
			s:    &Stack{},
			want: nil,
		},
		{
			name: "Peek from non-empty stack",
			s:    &Stack{1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Peek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}
