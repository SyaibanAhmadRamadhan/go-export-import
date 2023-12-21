package conf

import (
	"fmt"

	"github.com/SyaibanAhmadRamadhan/gocatch/genv"
)

// PostgresConf is a struct for storing the configuration of postgresql connection
type PostgresConf struct {
	User     string
	Password string
	Host     string
	Port     string
	DB       string
	SSL      string
}

// ConnString is a method PostgresConf for generate connetion string postgresql
func (p PostgresConf) ConnString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s dbname=%s",
		p.Host, p.Port, p.User, p.Password, p.SSL, p.DB)
}

// EnvPostgresConf is a function to set PostgresConf from the environment,
// if there is no variable in os.Env then there will be a default value
func EnvPostgresConf() PostgresConf {
	return PostgresConf{
		User:     genv.GetEnv("POSTGRES_USER", "ROOT"),
		Password: genv.GetEnv("POSTGRES_PASSWORD", "root"),
		Host:     genv.GetEnv("POSTGRES_HOST", "127.0.0.1"),
		Port:     genv.GetEnv("POSTGRES_PORT", "5432"),
		DB:       genv.GetEnv("POSTGRES_DB", "dbroot"),
		SSL:      genv.GetEnv("POSTGRES_SSL", "disable"),
	}
}
