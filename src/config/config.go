package config

type Config struct {
	Port       string `toml:"port"`
	Db_url     string `toml:"db_url"`
	Db_type    string `toml:"db_type"`
	Log_path   string `toml:"log_path"`
	SessionKey string `toml:"session"`
}
