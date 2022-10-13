package env

var (
	Mysql  MysqlConfig
	Server ServerConfig
	Secret SecretKey
)

type MysqlConfig struct {
	DB                 string `env:"MYSQL_DATABASE"`
	PORT               string `env:"MYSQL_PORT" envDefault:"3306"`
	HOST               string `env:"MYSQL_HOST"`
	USER               string `env:"MYSQL_USER"`
	PASSWORD           string `env:"MYSQL_PASSWORD"`
	ADDITIONAL_CONFIGS string `env:"ADDITIONAL_CONFIGS"`
}

type ServerConfig struct {
	PORT string `env:"SERVER_PORT"`
	HOST string `env:"SERVER_HOST"`
}

type SecretKey struct {
	SECRET_KEY string `env:"SECRET_KEY"`
}

func Load() {
	loadStructWithEnvVars(&Mysql)
	loadStructWithEnvVars(&Server)
	loadStructWithEnvVars(&Secret)
}
