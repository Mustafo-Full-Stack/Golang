package models

// id ролей совпадают с roles.id в migrations/001_roles_users.sql:
// там роли вставляются в этом же порядке, и база выдаёт им id 1, 2, 3...
// В access-токене хранится именно id, права проверяются по этим константам.
const (
	RoleUser      int64 = 1
	RoleAdmin     int64 = 2
	RoleAuditor   int64 = 3
	RoleModerator int64 = 4
	RoleSupport   int64 = 5
	RoleEmployer  int64 = 6
	RoleApplicant int64 = 7
)

// writeRoles — роли, которым можно создавать контент.
// Неизвестный id не попадает в map и получает отказ.
var writeRoles = map[int64]bool{
	RoleUser:      true,
	RoleAdmin:     true,
	RoleModerator: true,
	RoleEmployer:  true,
	RoleApplicant: true,
}
