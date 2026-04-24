package bintree_int

import "testing"

func TestElement_IsComplete(t *testing.T) {
	root := New(1)
	root.left = New(2)
	root.right = New(3)
	root.left.left = New(4)

	if !root.IsComplete() {
		t.Errorf("Expected root to be complete.")
	}

	root.left.right = New(5)
	if !root.IsComplete() {
		t.Errorf("Expected root to still be complete after adding root.left.right.")
	}

	root.right.right = New(7)
	if root.IsComplete() {
		t.Errorf("Expected root not to be complete after adding root.right.right.")
	}
}

func TestElement_IsPerfect(t *testing.T) {
	root := New(1)
	root.left = New(2)
	root.right = New(3)

	if !root.IsPerfect() {
		t.Errorf("Expected root to be perfect.")
	}

	root.left.left = New(4)
	if root.IsPerfect() {
		t.Errorf("Expected root not to be perfect after adding root.left.left.")
	}

	root.left.right = New(5)
	root.right.left = New(6)
	root.right.right = New(7)
	if !root.IsPerfect() {
		t.Errorf("Expected root to be perfect after adding all children.")
	}
}

func TestElement_IsBalanced(t *testing.T) {
	root := New(1)
	root.left = New(2)
	root.right = New(3)

	if !root.IsBalanced() {
		t.Errorf("Expected root to be balanced.")
	}

	root.left.left = New(4)
	if !root.IsBalanced() {
		t.Errorf("Expected root to still be balanced after adding root.left.left.")
	}

	root.left.left.left = New(8)
	if root.IsBalanced() {
		t.Errorf("Expected root not to be balanced after adding root.left.left.left.")
	}
}
