package search

import (
	"reflect"
	"testing"

	"github.com/calvn/al-go/tree"
)

func TestBFS(t *testing.T) {
	want := []int{1, 2, 3, 4}
	root := tree.New(want)

	got := BFS(root)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("BFS(%q): got, want%#[1]q", want)
	}
}
