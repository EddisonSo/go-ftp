package types 

type ServerConfig struct {
    Host Host
    Port int
}

type Host struct {
    Hostname string
}
