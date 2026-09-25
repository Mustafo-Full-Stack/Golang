package users

// User хранит имя и возраст одного пользователя.
type User struct {
	Name string
	Age  int
}

// HaveBirthday увеличивает возраст пользователя на один год.
func (u *User) HaveBirthday() {
	u.Age++
}

// IsAdult проверяем пользователю он совершилетный или нет
func (u *User) IsAdult() bool {
	if u.Age >= 18 {
		return true
	}

	return false
}
