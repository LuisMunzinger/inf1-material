package bintree_int

import (
	"slices"
	"testing"
)

func TestElement_InOrder(t *testing.T) {
	root := New(1)
	root.left = New(2)
	root.right = New(3)
	root.left.left = New(4)

	expected := []int{4, 2, 1, 3}
	result := root.InOrder()

	if !slices.Equal(expected, result) {
		t.Errorf("Expected InOrder to return %v, got %v", expected, result)
	}
}

func TestElement_PreOrder(t *testing.T) {
	root := New(1)
	root.left = New(2)
	root.right = New(3)
	root.left.left = New(4)

	expected := []int{1, 2, 4, 3}
	result := root.PreOrder()

	if !slices.Equal(expected, result) {
		t.Errorf("Expected PreOrder to return %v, got %v", expected, result)
	}
}

func TestElement_PostOrder(t *testing.T) {
	root := New(1)
	root.left = New(2)
	root.right = New(3)
	root.left.left = New(4)

	expected := []int{4, 2, 3, 1}
	result := root.PostOrder()

	if !slices.Equal(expected, result) {
		t.Errorf("Expected PostOrder to return %v, got %v", expected, result)
	}
}
