package config

type Config struct {
	Database DatabaseConfigurations
	Access   AccessConfigurations
	Server ServerConfigurations
	Logger LoggerConfigurations
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

type ServerConfigurations struct {
	ListenPort string
	Hostname string
}

type LoggerConfigurations struct {
	Environment string
}
