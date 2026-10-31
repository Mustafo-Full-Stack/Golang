package models

import (
	"strconv"
	"testing"
)

func TestCurrentUserCanCreate(t *testing.T) {
	tests := []struct {
		roleID int64
		want   bool
	}{
		{roleID: RoleUser, want: true},
		{roleID: RoleAdmin, want: true},
		{roleID: RoleModerator, want: true},
		{roleID: RoleEmployer, want: true},
		{roleID: RoleApplicant, want: true},
		{roleID: RoleAuditor, want: false},
		{roleID: RoleSupport, want: false},
		{roleID: 999, want: false},
		{roleID: 0, want: false},
	}

	for _, tt := range tests {
		t.Run(strconv.FormatInt(tt.roleID, 10), func(t *testing.T) {
			user := CurrentUser{RoleID: tt.roleID}
			if got := user.CanCreate(); got != tt.want {
				t.Fatalf("CanCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCurrentUserCanEditPost(t *testing.T) {
	tests := []struct {
		name     string
		user     CurrentUser
		authorID int64
		want     bool
	}{
		{name: "author edits own post", user: CurrentUser{ID: 1, RoleID: RoleUser}, authorID: 1, want: true},
		{name: "user cannot edit another post", user: CurrentUser{ID: 1, RoleID: RoleUser}, authorID: 2, want: false},
		{name: "moderator edits any post", user: CurrentUser{ID: 10, RoleID: RoleModerator}, authorID: 2, want: true},
		{name: "admin edits any post", user: CurrentUser{ID: 10, RoleID: RoleAdmin}, authorID: 2, want: true},
		{name: "auditor cannot edit even own post", user: CurrentUser{ID: 1, RoleID: RoleAuditor}, authorID: 1, want: false},
		{name: "support cannot edit even own post", user: CurrentUser{ID: 1, RoleID: RoleSupport}, authorID: 1, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.user.CanEditPost(tt.authorID); got != tt.want {
				t.Fatalf("CanEditPost(%d) = %v, want %v", tt.authorID, got, tt.want)
			}
		})
	}
}

func TestCurrentUserCanManageFile(t *testing.T) {
	tests := []struct {
		name    string
		user    CurrentUser
		ownerID int64
		want    bool
	}{
		{name: "owner", user: CurrentUser{ID: 1, RoleID: RoleUser}, ownerID: 1, want: true},
		{name: "another user", user: CurrentUser{ID: 1, RoleID: RoleUser}, ownerID: 2, want: false},
		{name: "moderator cannot manage another file", user: CurrentUser{ID: 10, RoleID: RoleModerator}, ownerID: 2, want: false},
		{name: "admin manages any file", user: CurrentUser{ID: 10, RoleID: RoleAdmin}, ownerID: 2, want: true},
		{name: "auditor cannot manage even own file", user: CurrentUser{ID: 1, RoleID: RoleAuditor}, ownerID: 1, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.user.CanManageFile(tt.ownerID); got != tt.want {
				t.Fatalf("CanManageFile(%d) = %v, want %v", tt.ownerID, got, tt.want)
			}
		})
	}
}
