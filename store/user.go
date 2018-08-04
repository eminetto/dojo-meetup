package store

import "github.com/meetup/model"

type Storage struct {
	users []model.User
}

func (s *Storage) StoreUser(u *model.User) error {
	s.users = append(s.users, *u)
	return nil
}

func (s *Storage) GetLastUser() *model.User {
	if len(s.users) > 0 {
		return &s.users[len(s.users)-1]
	}
	return nil
}
