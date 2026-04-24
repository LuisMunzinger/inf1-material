package bintree_int

import "testing"

func TestElement_IsEmpty_empty(t *testing.T) {
	e := Empty()
	if !e.IsEmpty() {
		t.Errorf("Expected e to be empty.")
	}
}

func TestElement_IsEmpty_notEmpty(t *testing.T) {
	e := New(5)
	if e.IsEmpty() {
		t.Errorf("Expected e to not be empty.")
	}
}

func TestElement_IsLeaf(t *testing.T) {
	root := New(10)
	root.left = New(5)
	root.right = New(15)
	root.left.left = New(3)

	if root.IsLeaf() {
		t.Errorf("Expected root to not be a leaf.")
	}
	if root.left.IsLeaf() {
		t.Errorf("Expected root.left not to be a leaf.")
	}
	if !root.right.IsLeaf() {
		t.Errorf("Expected root.right to be a leaf.")
	}
	if !root.left.left.IsLeaf() {
		t.Errorf("Expected root.left.left to be a leaf.")
	}
}

func TestElement_Count(t *testing.T) {
	root := New(1)
	root.left = New(2)
	root.right = New(3)
	root.left.left = New(4)

	if root.Count() != 4 {
		t.Errorf("Expected root.Count() to be %d, got %d", 4, root.Count())
	}
	if root.left.Count() != 2 {
		t.Errorf("Expected root.left.Count() to be %d, got %d", 2, root.left.Count())
	}
	if root.right.Count() != 1 {
		t.Errorf("Expected root.right.Count() to be %d, got %d", 1, root.right.Count())
	}
	if root.left.left.Count() != 1 {
		t.Errorf("Expected root.left.left.Count() to be %d, got %d", 1, root.left.left.Count())
	}
	if root.left.right.Count() != 0 {
		t.Errorf("Expected root.left.right.Count() to be %d, got %d", 0, root.left.right.Count())
	}
}

func TestElement_Height(t *testing.T) {
	root := New(1)
	root.left = New(2)
	root.right = New(3)
	root.left.left = New(4)

	if root.Height() != 3 {
		t.Errorf("Expected root.Height() to be %d, got %d", 3, root.Height())
	}
	if root.left.Height() != 2 {
		t.Errorf("Expected root.left.Height() to be %d, got %d", 2, root.left.Height())
	}
	if root.right.Height() != 1 {
		t.Errorf("Expected root.right.Height() to be %d, got %d", 1, root.right.Height())
	}
	if root.left.left.Height() != 1 {
		t.Errorf("Expected root.left.left.Height() to be %d, got %d", 1, root.left.left.Height())
	}
	if root.left.right.Height() != 0 {
		t.Errorf("Expected root.left.right.Height() to be %d, got %d", 0, root.left.right.Height())
	}
}
