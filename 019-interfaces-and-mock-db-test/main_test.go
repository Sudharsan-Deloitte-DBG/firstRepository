package main

import "testing"

func TestSaveUser(t *testing.T) {
	usr := User{
		ID:   1,
		name: "James",
	}

	db := MockDatastore{
		usr: make(map[int]User),
	}

	svc := Service{
		ds: db,
	}

	got := svc.SaveUser(usr)

	if got != nil {
		t.Errorf("SaveUser to DB failed: %v", got)
	}
}

func TestGetUser(t *testing.T) {
	usr := User{
		ID:   1,
		name: "James",
	}

	db := MockDatastore{
		usr: make(map[int]User),
	}

	svc := Service{
		ds: db,
	}

	svc.SaveUser(usr)

	user, err := svc.GetUser(1)

	if err != nil {
		t.Errorf("GetUser from MockDB failed:%v", err)
	}
	if user.name != "James" {
		t.Errorf("GetUser from MockDB failed:%v", err)
	}
}
