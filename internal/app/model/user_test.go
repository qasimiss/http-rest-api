package model_test

import (
	"testing"

	"github.com/qasimiss/http-rest-api/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}

func TestUser_Validate(t *testing.T) {

	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name:    "valid",
			u:       func() *model.User { return model.TestUser(t) },
			isValid: true,
		},
		{
			name: "empty_email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "empty_email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "Invalid"
				return u
			},
			isValid: false,
		},
		{
			name: "empty_password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "short_password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "short"
				return u
			},
			isValid: false,
		},
		{
			name: "with_encrypted_password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "ecnregergerg"
				return u
			},
			isValid: true,
		},
	}
	for _, tc := range testCases {
		t.Run(
			tc.name, func(t *testing.T) {
				if tc.isValid {
					assert.NoError(t, tc.u().Validate())
				} else {
					assert.Error(t, tc.u().Validate())
				}
			},
		)
	}

}
