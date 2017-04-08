package logging

type Level int

const (
	INFO Level = iota
	WARNING
	ERROR
)

func (l Level) String() string {
	switch l {
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	default:
		return "ERROR"
	}
}
