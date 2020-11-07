package structual

import "fmt"

// User 结构
type User struct {
	ID int32
}

// UserList 栈
type UserList []User

// FindUser 接口实现
func (t *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*t); i++ {
		if (*t)[i].ID == id {
			return (*t)[i], nil
		}
	}
	return User{}, fmt.Errorf("User %d could not be cound", id)
}

func (t *UserList) addUser(newUser User) {
	*t = append(*t, newUser)
}

// UserFinder 代理接口
type UserFinder interface {
	FindUser(id int32) (User, error)
}

// UserListProxy 代理结构
type UserListProxy struct {
	MockedDatabase      *UserList
	StackCache          UserList
	StackSize           int
	LastSearchUsedCache bool
}

// FindUser 接口实现
func (u *UserListProxy) FindUser(id int32) (User, error) {
	// Search cache list first
	user, err := u.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("Returning user from cache")
		u.LastSearchUsedCache = true
		return user, nil
	}

	// Search in the database
	user, err = u.MockedDatabase.FindUser(id)
	if err != nil {
		return User{}, err
	}

	u.addUserToStack(user)
	u.LastSearchUsedCache = false
	return user, nil
}

func (u *UserListProxy) addUserToStack(user User) {
	if len(u.StackCache) >= u.StackSize {
		u.StackCache = append(u.StackCache[1:], user)
	} else {
		u.StackCache.addUser(user)
	}
}
