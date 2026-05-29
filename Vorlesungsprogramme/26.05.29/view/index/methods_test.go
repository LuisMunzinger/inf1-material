package index

import (
	"inf1-material/Vorlesungsprogramme/26.05.29/core/user"
	"slices"
	"testing"
)

func TestIndex_Insert(t *testing.T) {
	idx := ById()

	u1 := &user.User{Id: "1", Nickname: "alice", Name: "Alice", Surname: "Smith"}
	u2 := &user.User{Id: "2", Nickname: "bob", Name: "Bob", Surname: "Johnson"}
	u3 := &user.User{Id: "3", Nickname: "charlie", Name: "Charlie", Surname: "Brown"}

	idx.Insert(u2)
	idx.Insert(u1)
	idx.Insert(u3)

	if idx.root.User != u2 {
		t.Errorf("Expected root to be u2")
	}
	if idx.root.Left.User != u1 {
		t.Errorf("Expected left child to be u1")
	}
	if idx.root.Right.User != u3 {
		t.Errorf("Expected right child to be u3")
	}
}

func TestIndex_Find(t *testing.T) {
	idx := ById()

	u1 := &user.User{Id: "1", Nickname: "alice", Name: "Alice", Surname: "Smith"}
	u2 := &user.User{Id: "2", Nickname: "bob", Name: "Bob", Surname: "Johnson"}
	u3 := &user.User{Id: "3", Nickname: "charlie", Name: "Charlie", Surname: "Brown"}

	idx.Insert(u2)
	idx.Insert(u1)
	idx.Insert(u3)

	if idx.Find("1").User != u1 {
		t.Errorf("Expected to find u1 with key '1'")
	}
	if idx.Find("2").User != u2 {
		t.Errorf("Expected to find u2 with key '2'")
	}
	if idx.Find("3").User != u3 {
		t.Errorf("Expected to find u3 with key '3'")
	}
	if idx.Find("4") != nil {
		t.Errorf("Expected to not find any user with key '4'")
	}
}

func TestIndex_List(t *testing.T) {
	idx := ById()

	u1 := &user.User{Id: "1", Nickname: "alice", Name: "Alice", Surname: "Smith"}
	u2 := &user.User{Id: "2", Nickname: "bob", Name: "Bob", Surname: "Johnson"}
	u3 := &user.User{Id: "3", Nickname: "charlie", Name: "Charlie", Surname: "Brown"}

	idx.Insert(u2)
	idx.Insert(u1)
	idx.Insert(u3)

	users := idx.List()
	if len(users) != 3 {
		t.Errorf("Expected list to contain 3 users, got %d", len(users))
	}
	if users[0] != u1 {
		t.Errorf("Expected first user to be u1")
	}
	if users[1] != u2 {
		t.Errorf("Expected second user to be u2")
	}
	if users[2] != u3 {
		t.Errorf("Expected third user to be u3")
	}
}

func TestIndex_Keys(t *testing.T) {
	idx := ById()

	u1 := &user.User{Id: "1", Nickname: "alice", Name: "Alice", Surname: "Smith"}
	u2 := &user.User{Id: "2", Nickname: "bob", Name: "Bob", Surname: "Johnson"}
	u3 := &user.User{Id: "3", Nickname: "charlie", Name: "Charlie", Surname: "Brown"}

	idx.Insert(u2)
	idx.Insert(u1)
	idx.Insert(u3)

	keys := idx.Keys()
	expectedKeys := []string{"1", "2", "3"}

	if !slices.Equal(keys, expectedKeys) {
		t.Errorf("Expected keys to be %v, got %v", expectedKeys, keys)
	}
}

func TestIndex_String(t *testing.T) {
	idx := ById()

	u1 := &user.User{Id: "1", Nickname: "alice", Name: "Alice", Surname: "Smith"}
	u2 := &user.User{Id: "2", Nickname: "bob", Name: "Bob", Surname: "Johnson"}
	u3 := &user.User{Id: "3", Nickname: "charlie", Name: "Charlie", Surname: "Brown"}

	idx.Insert(u2)
	idx.Insert(u1)
	idx.Insert(u3)

	expectedString := "[1 2 3]"
	if idx.String() != expectedString {
		t.Errorf("Expected string representation to be %s, got %s", expectedString, idx.String())
	}
}
