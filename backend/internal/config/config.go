package config

type Logger struct {
	FileOutputPath string `yaml:"output_file_path"`
	OutputToStdout bool   `yaml:"output_to_stdout"`
	Development    bool   `yaml:"development"`
}

type Config struct {
	Logger Logger `yaml:"logger"`
}
