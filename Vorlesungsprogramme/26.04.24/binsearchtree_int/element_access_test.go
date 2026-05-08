package binsearchtree_int

import "testing"

func TestElement_GetValue(t *testing.T) {
	root := New(50)
	root.InsertValueAt("L", 25)
	root.InsertValueAt("R", 75)
	root.InsertValueAt("LL", 12)

	if v := root.Left().Value(); v != 25 {
		t.Errorf("Expected value at path 'L' to be %d, got %d", 25, v)
	}
	if v := root.Right().Value(); v != 75 {
		t.Errorf("Expected value at path 'R' to be %d, got %d", 75, v)
	}
	if v := root.Left().Left().Value(); v != 12 {
		t.Errorf("Expected value at path 'LL' to be %d, got %d", 12, v)
	}
}

func TestElement_InsertValue(t *testing.T) {
	root := Empty()
	root.InsertValue(100)
	root.InsertValue(50)
	root.InsertValue(150)
	root.InsertValue(25)

	if v := root.Value(); v != 100 {
		t.Errorf("Expected root value to be %d, got %d", 100, v)
	}
	if v := root.Left().Value(); v != 50 {
		t.Errorf("Expected value at path 'L' to be %d, got %d", 50, v)
	}
	if v := root.Right().Value(); v != 150 {
		t.Errorf("Expected value at path 'R' to be %d, got %d", 150, v)
	}
	if v := root.Left().Left().Value(); v != 25 {
		t.Errorf("Expected value at path 'LL' to be %d, got %d", 25, v)
	}
}
