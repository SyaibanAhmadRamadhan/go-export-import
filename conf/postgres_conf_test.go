package conf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvPostgresConf(t *testing.T) {
	t.Run("test default value", func(t *testing.T) {
		expected := PostgresConf{
			User:     "root",
			Password: "root",
			Host:     "127.0.0.1",
			Port:     "5432",
			DB:       "dbroot",
			SSL:      "disable",
		}

		actual := EnvPostgresConf()
		assert.Equal(t, expected, actual)
	})
}
