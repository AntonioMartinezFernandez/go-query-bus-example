package user

import "time"

type UserRepository interface {
	GetUser(id string) (*User, error)
	SaveUser(user User) error
}

type InMemoryUserRepository struct {
	users map[string]User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	users := map[string]User{
		"1": {
			id:        "1",
			name:      "John",
			birthdate: time.Now(),
		},
		"2": {
			id:        "2",
			name:      "Matt",
			birthdate: time.Now(),
		},
	}

	return &InMemoryUserRepository{
		users: users,
	}
}

func (ur *InMemoryUserRepository) GetUser(id string) (*User, error) {
	user := ur.users[id]
	return &user, nil
}

func (ur *InMemoryUserRepository) SaveUser(user User) error {
	ur.users[user.Id()] = user
	return nil
}
