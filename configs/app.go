package configs

type AppConfig struct {
	Env Env
}

func LoadAppConfig() AppConfig {
	return AppConfig{Env: LoadEnv()}
}
