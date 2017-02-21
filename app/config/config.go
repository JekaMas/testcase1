package config

//AppConfig application config
var AppConfig Config

//New - create new config
func New() *Config {
	return &AppConfig
}

//Config - config for app
type Config struct {
	Host            string
	Port            int
}

//GetPort - get service port
func (this Config) GetPort() int {
	if this.Port <= 0 || this.Port > 65535 {
		panic("invalid http port number")
	}
	return this.Port
}

