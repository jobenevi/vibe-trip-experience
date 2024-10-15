package database_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jobenevi/vibe-trip-experience/database"
	"github.com/stretchr/testify/assert"
)

func createTempEnvFile(t *testing.T, content string) string {
	dir, err := os.MkdirTemp("", "configtest")
	assert.NoError(t, err)

	envFilePath := filepath.Join(dir, ".env")
	err = os.WriteFile(envFilePath, []byte(content), 0644)
	assert.NoError(t, err)

	return dir
}

func TestConnectionToDatabaseFailLoadConfig(t *testing.T) {
	envContent := `
DATABASE_USER=testuser
DATABASE_PASSWORD=testpassword
DATABASE_NAME=testdb
DATABASE_SSL_MODE=disable
WEB_SERVER_PORT=8080
JWT_SECRET=testsecret
JWT_EXPIRES_IN=3600
`
	dir := createTempEnvFile(t, envContent)

	os.Setenv("CONFIG_PATH", dir)
	defer os.Unsetenv("CONFIG_PATH")

	assert.Panics(t, func() { database.ConnectionToDatabase() })
}

func TestConnectionToDatabaseFailConnection(t *testing.T) {

	dir := createTempEnvFile(t, "")
	os.Setenv("CONFIG_PATH", dir)
	defer os.Unsetenv("CONFIG_PATH")

	assert.Panics(t, func() { database.ConnectionToDatabase() })
}
