package config

// Conf ... Data structure
type Conf struct {
	User        string
	Group       string
	Port        int
	EnableSSL   bool
	BaseFolder  string
	SslKeyFile  string
	SslCertFile string
	PidFile     string
}

// New ... Creates and populates conf data structure with sensible defaults
func New() Conf {
	var c = Conf{
		User:        "www-data",
		Group:       "www-data",
		Port:        8000,
		EnableSSL:   false,
		SslKeyFile:  "",
		SslCertFile: "",
		BaseFolder:  "/var/www",
		PidFile:     "/var/run/steelgrave.pid",
	}
	return c
}
