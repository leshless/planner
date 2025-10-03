package config

type Holder interface {
	Config() Config
}

type holder struct {
	config *Config
}

var _ Holder = (*holder)(nil)

func NewHolder(config Config) *holder {
	return &holder{
		config: &config,
	}
}

func (r *holder) Config() Config {
	return *r.config
}
