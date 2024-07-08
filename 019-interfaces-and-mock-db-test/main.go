package main

import (
	"fmt"
	"log"
)

type User struct {
	ID   int
	name string
}

type MockDatastore struct {
	usr map[int]User
}

func (m MockDatastore) GetUser(id int) (User, error) {
	user, ok := m.usr[id]
	if !ok {
		return user, fmt.Errorf("User with ID %d not found in DB", id)
	}
	return m.usr[id], nil
}

func (m MockDatastore) SaveUser(u User) error {
	_, ok := m.usr[u.ID]
	if ok {
		return fmt.Errorf("User with ID %d already exists", u.ID)
	}
	m.usr[u.ID] = u
	return nil
}

type Datastore interface {
	GetUser(int) (User, error)
	SaveUser(User) error
}

type Service struct {
	ds Datastore
}

func (s Service) GetUser(i int) (User, error) {
	return s.ds.GetUser(i)
}

func (s Service) SaveUser(u User) error {
	return s.ds.SaveUser(u)
}

func main() {
	u1 := User{
		ID:   1,
		name: "James",
	}

	db := MockDatastore{
		usr: make(map[int]User),
	}

	svc := Service{
		ds: db,
	}

	if err := svc.SaveUser(u1); err != nil {
		log.Fatalf("SaveUser to MockDatastore failed: %v", err)
	}

	if usr, err := svc.GetUser(2); err != nil {
		log.Fatalf("GetUser from MockDatastore failed: %v", err)
	} else {
		fmt.Println(usr)
	}

}
