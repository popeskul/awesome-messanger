package config

import "time"

type Config struct {
	ConnectionString  string
	Host              string
	Port              int
	User              string
	Password          string
	DBName            string
	SSLMode           string
	MaxConnections    int32
	MinConnections    int32
	MaxConnLifetime   time.Duration
	MaxConnIdleTime   time.Duration
	HealthCheckPeriod time.Duration
}

func New() *Config {
	return &Config{
		MaxConnections:    10,
		MinConnections:    2,
		MaxConnLifetime:   time.Hour,
		MaxConnIdleTime:   time.Minute * 30,
		HealthCheckPeriod: time.Minute,
		SSLMode:           "disable",
	}
}

func (c *Config) WithConnectionString(cs string) *Config {
	c.ConnectionString = cs
	return c
}

func (c *Config) WithHost(host string) *Config {
	c.Host = host
	return c
}

func (c *Config) WithPort(port int) *Config {
	c.Port = port
	return c
}

func (c *Config) WithUser(user string) *Config {
	c.User = user
	return c
}

func (c *Config) WithPassword(password string) *Config {
	c.Password = password
	return c
}

func (c *Config) WithDBName(dbName string) *Config {
	c.DBName = dbName
	return c
}

func (c *Config) WithSSLMode(sslMode string) *Config {
	c.SSLMode = sslMode
	return c
}

func (c *Config) WithMaxConnections(mc int32) *Config {
	c.MaxConnections = mc
	return c
}

func (c *Config) WithMinConnections(mc int32) *Config {
	c.MinConnections = mc
	return c
}

func (c *Config) WithMaxConnLifetime(mcl time.Duration) *Config {
	c.MaxConnLifetime = mcl
	return c
}

func (c *Config) WithMaxConnIdleTime(mcit time.Duration) *Config {
	c.MaxConnIdleTime = mcit
	return c
}

func (c *Config) WithHealthCheckPeriod(hcp time.Duration) *Config {
	c.HealthCheckPeriod = hcp
	return c
}

func (c *Config) Build() *Config {
	return c
}
