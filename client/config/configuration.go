package config

type Config struct {
	IngestBaseUrl string
	DisableUpload bool
	SaveLocally   bool
	OfflinePath   string
	Offline       bool
	Debug         bool
	LogLevel      string
}

var GlobalConfiguration = &Config{
	LogLevel: "INFO",
}
