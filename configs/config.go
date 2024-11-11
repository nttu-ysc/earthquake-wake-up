package configs

type Config struct {
	AppTimeout   int      `mapstructure:"app_timeout"`
	NotifyGroups []string `mapstructure:"notify_groups"`
	Line         `yaml:"line"`
	File         `yaml:"file"`
	Discord      `yaml:"discord"`
}

type Line struct {
	Token string `mapstructure:"token"`
}

type File struct {
	Path string `mapstructure:"path"`
}

type Discord struct {
	Webhook string `mapstructure:"webhook"`
}
