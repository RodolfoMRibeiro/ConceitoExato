package env

var (
	Database DatabaseConfig
	Server   ServerConfig
	Secret   SecretKey
)

type DatabaseConfig struct {
	DATABASE_URL string `env:"DATABASE_URL"`
}

type ServerConfig struct {
	PORT string `env:"PORT"`
	HOST string `env:"SERVER_HOST"`
}

type SecretKey struct {
	SECRET_KEY string `env:"SECRET_KEY"`
}

func Load() {
	loadStructWithEnvVars(&Database)
	loadStructWithEnvVars(&Server)
	loadStructWithEnvVars(&Secret)
}
