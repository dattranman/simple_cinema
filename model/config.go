package model

type Configuration struct {
	ServiceSettings ServiceSettings `yaml:"service_settings"`
	SQLSettings     SQLSettings     `yaml:"sql_settings"`
	CacheSettings   CacheSetting    `yaml:"cache_settings"`
}

type ServiceSettings struct {
	Port string `yaml:"port"`
}

type SQLSettings struct {
	DriverName string `yaml:"driver_name"`
	URI        string `yaml:"uri"`
	Timeout    int    `yaml:"timeout"`
	Debug      bool   `yaml:"debug"`
}

type CacheSetting struct {
	URI      string `yaml:"uri"`
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
	Timeout  string `yaml:"timeout"`
}
