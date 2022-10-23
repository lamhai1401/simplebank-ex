package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func genPassWord() (string, error) {
	password := RandomString(6)
	return HashPassword(password)
}
func TestPassword(t *testing.T) {
	t.Run(("Test Check Password"), func(t *testing.T) {
		password := RandomString(6)
		hashedPassword1, err := HashPassword(password)
		require.NoError(t, err)
		require.NotEmpty(t, hashedPassword1)
		err = CheckPassword(password, hashedPassword1)
		require.NoError(t, err)
	})

	t.Run("Test Password Compare", func(t *testing.T) {
		hashedPassword1, err := genPassWord()
		require.NoError(t, err)
		require.NotEmpty(t, hashedPassword1)

		hashedPassword2, err := genPassWord()
		require.NoError(t, err)
		require.NotEmpty(t, hashedPassword2)

		// compate two password
		require.NotEqual(t, hashedPassword1, hashedPassword2)
	})

	t.Run("Test Wrong Password ", func(t *testing.T) {
		hashedPassword1, err := genPassWord()
		require.NoError(t, err)
		wrongPassword := RandomString(6)
		err = CheckPassword(wrongPassword, hashedPassword1)
		require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
	})
}
