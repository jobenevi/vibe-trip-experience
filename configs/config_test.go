package config

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTempEnvFile(t *testing.T, content string) string {
	t.Helper()

	dir := t.TempDir()

	envFilePath := filepath.Join(dir, ".env")
	file, err := os.Create(envFilePath)
	if err != nil {
		t.Fatalf("failed to create temp env file: %v", err)
	}
	defer file.Close()

	_, err = io.WriteString(file, content)
	if err != nil {
		t.Fatalf("failed to write to temp env file: %v", err)
	}

	return dir
}

func TestLoadConfigSuccess(t *testing.T) {
	envContent := `
DATABASE_DRIVER=postgres
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USER=testuser
DATABASE_PASSWORD=testpassword
DATABASE_NAME=testdb
DATABASE_SSL_MODE=disable
WEB_SERVER_PORT=8080
JWT_SECRET=testsecret
JWT_EXPIRES_IN=3600
`
	dir := createTempEnvFile(t, envContent)

	config, err := LoadConfig(dir)
	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, "postgres", config.DatabaseDriver)
}

func TestLoadConfigFail(t *testing.T) {
	dir := t.TempDir()

	config, err := LoadConfig(dir)
	assert.Error(t, err)
	assert.Nil(t, config)
}
