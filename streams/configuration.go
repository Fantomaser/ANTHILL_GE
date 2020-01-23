package streams

// MonitorRect is monitor real size
type MonitorRect struct {
	H int
	W int
}

// MonitorInfo struct of monitor info
type MonitorInfo struct {
	MonitorParam []string
	MonitorSize  []MonitorRect
}

type Config struct {
	Monitor MonitorInfo
}

var conf Config
