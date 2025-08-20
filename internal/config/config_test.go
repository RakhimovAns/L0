package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/RakhimovAns/L0/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig_Success(t *testing.T) {
	tmpDir := t.TempDir()
	envPath := filepath.Join(tmpDir, "local.env")

	content := `
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=test
POSTGRES_DB=testdb
POSTGRES_PASSWORD=secret
DATABASE_POSTGRES_MIGRATIONS_PATH=./migrations
HTTP_PORT=8080
HTTP_HOST=127.0.0.1
KAFKA_BROKER=localhost:9092
KAFKA_TOPIC=orders
KAFKA_GROUP_ID=group1
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0
`
	err := os.WriteFile(envPath, []byte(content), 0644)
	assert.NoError(t, err)

	oldPath := "./configs/local.env"
	_ = os.Rename(oldPath, oldPath+".bak")
	defer os.Rename(oldPath+".bak", oldPath)

	_ = os.MkdirAll("./configs", 0755)
	_ = os.WriteFile(oldPath, []byte(content), 0644)

	cfg := config.NewConfig()
	assert.NotNil(t, cfg)
	assert.Equal(t, "localhost", cfg.PostgresHost)
	assert.Equal(t, "8080", cfg.HTTPPort)
}

func TestNewConfig_Fail(t *testing.T) {
	oldPath := "./configs/local.env"
	_ = os.Rename(oldPath, oldPath+".bak")
	defer os.Rename(oldPath+".bak", oldPath)

	cfg := config.NewConfig()
	assert.Nil(t, cfg)
}
