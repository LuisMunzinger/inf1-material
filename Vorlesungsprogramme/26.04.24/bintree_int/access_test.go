package bintree_int

import "testing"

func TestElement_SetValue(t *testing.T) {
	e := Empty()
	e.SetValue(42)

	if e.value != 42 {
		t.Errorf("Expected e.value to be %d, got %d", 42, e.value)
	}
	if e.IsEmpty() {
		t.Errorf("Expected e to be non-empty after setting value.")
	}
}

func TestElement_GetValue(t *testing.T) {
	root := New(10)
	root.left = New(5)
	root.right = New(15)
	root.left.left = New(3)

	if value, err := root.GetValueAt(""); err != nil {
		t.Errorf("Expected GetValue to succeed for empty path.")
	} else if value != 10 {
		t.Errorf("Expected value at root to be %d, got %d", 10, value)
	}

	if value, err := root.GetValueAt("L"); err != nil {
		t.Errorf("Expected GetValue to succeed for path 'L'.")
	} else if value != 5 {
		t.Errorf("Expected value at path 'L' to be %d, got %d", 5, value)
	}

	if value, err := root.GetValueAt("R"); err != nil {
		t.Errorf("Expected GetValue to succeed for path 'R'.")
	} else if value != 15 {
		t.Errorf("Expected value at path 'R' to be %d, got %d", 15, value)
	}

	if value, err := root.GetValueAt("LL"); err != nil {
		t.Errorf("Expected GetValue to succeed for path 'LL'.")
	} else if value != 3 {
		t.Errorf("Expected value at path 'LL' to be %d, got %d", 3, value)
	}

	if _, err := root.GetValueAt("LR"); err == nil {
		t.Errorf("Expected GetValue to fail for path 'LR' (empty element).")
	}

	if _, err := root.GetValueAt("LLL"); err == nil {
		t.Errorf("Expected GetValue to fail for path 'LLL' (non-existent element).")
	}

	if _, err := root.GetValueAt("X"); err == nil {
		t.Errorf("Expected GetValue to fail for invalid path 'X'.")
	}
}

func TestElement_InsertValue(t *testing.T) {
	root := Empty()
	err := root.InsertValueAt("", 100)

	if err != nil {
		t.Errorf("Expected InsertValue to succeed for empty path.")
		if v, e := root.GetValueAt(""); e != nil {
			t.Errorf("Expected GetValue to succeed for empty path after insertion.")
		} else if v != 100 {
			t.Errorf("Expected root value to be %d, got %d", 100, v)
		}
	}

	if err = root.InsertValueAt("L", 50); err != nil {
		t.Errorf("Expected InsertValue to succeed for path 'L'.")
		if v, e := root.GetValueAt("L"); e != nil {
			t.Errorf("Expected GetValue to succeed for path 'L' after insertion.")
		} else if v != 50 {
			t.Errorf("Expected value at path 'L' to be %d, got %d", 50, v)
		}
	}

	if err = root.InsertValueAt("R", 150); err != nil {
		t.Errorf("Expected InsertValue to succeed for path 'R'.")
		if v, e := root.GetValueAt("R"); e != nil {
			t.Errorf("Expected GetValue to succeed for path 'R' after insertion.")
		} else if v != 150 {
			t.Errorf("Expected value at path 'R' to be %d, got %d", 150, v)
		}
	}

	if err = root.InsertValueAt("LL", 25); err != nil {
		t.Errorf("Expected InsertValue to succeed for path 'LL'.")
		if v, e := root.GetValueAt("LL"); e != nil {
			t.Errorf("Expected GetValue to succeed for path 'LL' after insertion.")
		} else if v != 25 {
			t.Errorf("Expected value at path 'LL' to be %d, got %d", 25, v)
		}
	}

	if err = root.InsertValueAt("LRL", 75); err == nil {
		t.Errorf("Expected InsertValue to fail for path 'LRL'.")
	}
}
