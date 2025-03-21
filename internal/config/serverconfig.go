package config

type ServerConfig struct {
    Host Host
    Port int
    Homedir string
}

type Host struct {
    Hostname string
}
