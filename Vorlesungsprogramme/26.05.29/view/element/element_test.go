package element

import (
	"inf1-material/Vorlesungsprogramme/26.05.29/core/user"
	"testing"
)

func TestElement_IsEmpty(t *testing.T) {
	e1 := Empty()

	if !e1.IsEmpty() {
		t.Errorf("Expected empty element")
	}

	if e1.Left != nil {
		t.Errorf("Expected left child to be nil for empty element")
	}

	if e1.Right != nil {
		t.Errorf("Expected right child to be nil for empty element")
	}
}

func TestElement_SetUser(t *testing.T) {
	e1 := Empty()
	e1.SetUser(user.New("1", "Al", "Alan", "Turing"))

	if e1.IsEmpty() {
		t.Errorf("Expected element to be non-empty after setting user")
	}

	if e1.User.Id != "1" {
		t.Errorf("Expected user ID to be '1', got '%s'", e1.User.Id)
	}

	if e1.User.Nickname != "Al" {
		t.Errorf("Expected user nickname to be 'Al', got '%s'", e1.User.Nickname)
	}

	if e1.User.Name != "Alan" {
		t.Errorf("Expected user name to be 'Alan', got '%s'", e1.User.Name)
	}

	if e1.User.Surname != "Turing" {
		t.Errorf("Expected user surname to be 'Turing', got '%s'", e1.User.Surname)
	}

	if !e1.Left.IsEmpty() {
		t.Errorf("Expected left child to be empty after setting user")
	}

	if !e1.Right.IsEmpty() {
		t.Errorf("Expected right child to be empty after setting user")
	}
}

func TestElement_List(t *testing.T) {
	root := Empty()

	root.SetUser(user.New("2", "Bob", "Bob", "Builder"))
	root.Left.SetUser(user.New("1", "Al", "Alan", "Turing"))
	root.Right.SetUser(user.New("3", "Charlie", "Charlie", "Chocolate"))

	users := root.List()

	if len(users) != 3 {
		t.Errorf("Expected 3 users in list, got %d", len(users))
	}

	if users[0].Id != "1" {
		t.Errorf("Expected first user ID to be '1', got '%s'", users[0].Id)
	}

	if users[1].Id != "2" {
		t.Errorf("Expected second user ID to be '2', got '%s'", users[1].Id)
	}

	if users[2].Id != "3" {
		t.Errorf("Expected third user ID to be '3', got '%s'", users[2].Id)
	}
}
