package config

type Config struct {
	Database DatabaseConfigurations
	Access   AccessConfigurations
}

type AccessConfigurations struct {
	ExpirationTime string
	SigningKey     string
}

type DatabaseConfigurations struct {
	User           string
	Password       string
	Host           string
	Port           string
	Name           string
	SSLMode        string
	MigrationsPath string
}
