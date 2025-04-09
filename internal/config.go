package internal

var (
	GlobalConfig = Config{}
)

type Config struct {
	MusicPath  string `json:"music_path" mapstructure:"music_path" env:"music_path"`
	DBPath     string `json:"db_path" mapstructure:"db_path" env:"db_path"`
	OutputPath string `json:"output_path" mapstructure:"output_path" env:"output_path"`
}
