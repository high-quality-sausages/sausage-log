package conf

type SausageLogConf struct {
	SausageLogConfHandler
	Level       map[string]int64
	OutputPaths string
}

type SausageLogConfHandler interface {
	SetLevel(level string, score int64) bool
	GetLevel(level string) int64
	SetOutputPaths(path string) bool
	GetOutputPaths() string
}

// NewSausageLogConf returns a new SausageLogConf object
func NewSausageLogConf(level map[string]int64, outputPaths string) *SausageLogConf {
	return &SausageLogConf{
		Level:       level,
		OutputPaths: outputPaths,
	}
}

// SetLevel returns the result of set operation
func (s *SausageLogConf) SetLevel(level string, score int64) bool {

	// score == -1 is not available
	if score == -1 {
		return false
	}
	s.Level[level] = score
	return true
}

// GetLevel returns the score of the level
func (s *SausageLogConf) GetLevel(level string) int64 {
	if _, ok := s.Level[level]; ok {
		return s.Level[level]
	}
	return -1
}

// SetOutputPaths creates a directory if the path does not exist
func (s *SausageLogConf) SetOutputPaths(path string) bool {
	// TODO
	return true
}

// GetOutputPaths returns the path of log file
func (s *SausageLogConf) GetOutputPaths() string {
	return s.OutputPaths
}
