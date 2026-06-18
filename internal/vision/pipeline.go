package vision

type Config struct {
	TargetWidth  int
	TargetHeight int
	Mirror       bool
}

func DefaultConfig() Config {
	return Config{
		TargetWidth:  640,
		TargetHeight: 480,
		Mirror:       true,
	}
}

type ProcessedFrame struct {
	Data   []byte
	Width  int
	Height int
}

type Pipeline struct {
	config Config
}

func NewPipeline(config Config) *Pipeline {
	return &Pipeline{config: config}
}

func (p *Pipeline) Config() Config {
	return p.config
}
