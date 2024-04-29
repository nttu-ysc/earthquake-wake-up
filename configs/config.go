package configs

type Config struct {
	AppTimeout   int      `mapstructure:"app_timeout"`
	NotifyGroups []string `mapstructure:"notify_groups"`
	Line         struct {
		Token string `mapstructure:"token"`
	} `yaml:"line"`
	File struct {
		Path string `mapstructure:"path"`
	} `yaml:"file"`
}
