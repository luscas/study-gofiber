package database_test

import (
	"log"
	"testing"

	"api.droppy.com.br/pkg/database"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env:", err)
	}
}

func TestConnectPostgres(t *testing.T) {
	db, err := database.ConnectPostgres()
	require.NoError(t, err)
	assert.NotNil(t, db)
}
