package main

import (
	"fmt"
	"sync"
)

type User struct {
	ID    int
	Name  string
	Email string
}

type UserMap struct {
	mu   sync.RWMutex
	data map[int]*User
}

func NewUserMap() *UserMap {
	return &UserMap{
		data: make(map[int]*User),
	}
}

func (userMap *UserMap) AddUser(user *User) bool {
	userMap.mu.Lock()
	defer userMap.mu.Unlock()

	if _, exists := userMap.data[user.ID]; exists {
		return false
	}

	userMap.data[user.ID] = user
	return true
}

func (userMap *UserMap) GetUser(id int) (*User, bool) {
	userMap.mu.RLock()
	defer userMap.mu.RUnlock()

	user, exists := userMap.data[id]
	return user, exists
}

func (userMap *UserMap) GetAllUsers() []*User {
	userMap.mu.RLock()
	defer userMap.mu.RUnlock()

	users := make([]*User, 0, len(userMap.data))
	for _, user := range userMap.data {
		users = append(users, user)
	}
	return users
}

func main() {
	userMap := NewUserMap()
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			user := &User{
				ID:    id,
				Name:  fmt.Sprintf("User_%d", id),
				Email: fmt.Sprintf("user_%d@mail.com", id),
			}

			if userMap.AddUser(user) {
				fmt.Printf("Добавлен: %s\n", user.Name)
			} else {
				fmt.Printf("Уже существует: %s\n", user.Name)
			}
		}(i)
	}

	wg.Wait()

	duplicate := &User{ID: 1, Name: "Duplicate", Email: "duplicate@mail.com"}
	if userMap.AddUser(duplicate) {
		fmt.Println("\nДубликат добавлен")
	} else {
		fmt.Println("\nДубликат отклонен")
	}

	searchIDs := []int{1, 2, 7}

	fmt.Printf("\n")
	for _, id := range searchIDs {
		if user, exists := userMap.GetUser(id); exists {
			fmt.Printf("Найден пользователь ID %d: %s (%s)\n", id, user.Name, user.Email)
		} else {
			fmt.Printf("Пользователь ID %d не найден\n", id)
		}
	}

	fmt.Println("\nВсе пользователи:")
	for _, user := range userMap.GetAllUsers() {
		fmt.Printf("Name: %s (ID: %d) (Email: %s)\n", user.Name, user.ID, user.Email)
	}
}
