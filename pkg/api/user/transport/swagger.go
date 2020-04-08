package transport

import (
	"github.com/wednesday-solution/go-boiler"
)

// User model response
// swagger:response userResp
type swaggUserResponse struct {
	// in:body
	Body struct {
		*goboiler.User
	}
}

// Users model response
// swagger:response userListResp
type swaggUserListResponse struct {
	// in:body
	Body struct {
		Users []goboiler.User `json:"users"`
		Page  int          `json:"page"`
	}
}
