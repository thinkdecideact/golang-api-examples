package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Username string
	Password string
	Host     string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Username: "learn_by_examples",
			Password: "123456",
			Host:     "127.0.0.1:3306",
			Name:     "learn_by_examples",
			Charset:  "utf8",
		},
	}
}
