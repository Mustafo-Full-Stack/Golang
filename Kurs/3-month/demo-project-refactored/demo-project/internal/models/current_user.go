package models

// CurrentUser — пользователь, который выполняет запрос. Берётся из access-токена.
// Все правила прав — методы CurrentUser.
type CurrentUser struct {
	ID        int64
	RoleID    int64
	SessionID string
}

// CanCreate — может ли пользователь создавать посты и загружать файлы.
// auditor и support только читают.
func (u CurrentUser) CanCreate() bool {
	return writeRoles[u.RoleID]
}

// CanEditPost — автор правит свой пост, модератор и админ — любой.
func (u CurrentUser) CanEditPost(authorID int64) bool {
	if !u.CanCreate() {
		return false
	}
	return u.ID == authorID || u.RoleID == RoleAdmin || u.RoleID == RoleModerator
}

// CanManageFile — неприкреплённый файл видит и удаляет владелец, чужой — только админ.
func (u CurrentUser) CanManageFile(ownerID int64) bool {
	if !u.CanCreate() {
		return false
	}
	return u.ID == ownerID || u.RoleID == RoleAdmin
}
