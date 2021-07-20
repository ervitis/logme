package logme

import "github.com/ervitis/logme/v2/config_loaders"

type (
	Loggerme struct {
		logWrap *Wrapper
		opts    *configLoader
	}

	configLoader struct {
		loader *config_loaders.LoggermeConfig
	}

	Options func(*configLoader)
)

func WithConfigLoader(loader config_loaders.Loader) Options {
	return func(l *configLoader) {
		l.loader = config_loaders.NewLoggermeConfig(loader)
	}
}

func defaultConfigLoader() *configLoader {
	return &configLoader{}
}

func NewLogme(options ...Options) (*Loggerme, error) {
	opts := defaultConfigLoader()

	for _, v := range options {
		v(opts)
	}

	cfg, err := opts.loader.LoggermeConfigLoader()
	if err != nil {
		return nil, err
	}

	return &Loggerme{logWrap: LoggerWrapper(cfg)}, nil
}

func (l *Loggerme) L() *Wrapper {
	return l.logWrap
}
