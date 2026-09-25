package pass

func CheckPassword(password string) bool {
	return password == "go123"
}

func IsStrong(password string) bool {
	if len(password) >= 8 {
		return true
	}
	return false
}
