package jwt_test

import (
	"os"
	"testing"

	"api.droppy.com.br/pkg/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	os.Setenv("JWT_SECRET", "secret")
	jwt.InitSecretKey()
}

func TestValidTokenVerification(t *testing.T) {
	var payload = jwt.Payload{
		Sub:  "1234567890",
		Name: "Lucas",
	}

	var token, err = jwt.GenerateToken(payload)
	var verifyToken = jwt.VerifyToken(token)

	assert.NoError(t, err)
	require.NoError(t, err)

	assert.NoError(t, verifyToken)
	require.NoError(t, verifyToken)
}

func TestInvalidTokenVerification(t *testing.T) {
	var token = "invalid-token"

	var verifyToken = jwt.VerifyToken(token)

	assert.Error(t, verifyToken)
	require.Error(t, verifyToken)
}
