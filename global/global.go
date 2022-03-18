package global

import "time"

const (
	// Org is the organization
	Org = "MegaEase"
	// Prog is the program name
	Prog = "EaseProbe"
	// Ver is the program version
	Ver = "0.1"

	//OrgProg combine organization and program
	OrgProg = Org + " " + Prog
	//OrgProgVer combine organization and program and version
	OrgProgVer = Org + " " + Prog + "/" + Ver
)

const (
	// DefaultRetryTimes is 3 times
	DefaultRetryTimes = 3
	// DefaultRetryInterval is 5 seconds
	DefaultRetryInterval = time.Second * 5
	// DefaultTimeFormat is "2006-01-02 15:04:05 UTC"
	DefaultTimeFormat = "2006-01-02 15:04:05 UTC"
	// DefaultProbeInterval is 1 minutes
	DefaultProbeInterval = time.Second * 60
	// DefaultTimeOut is 30 seconds
	DefaultTimeOut = time.Second * 30
)

// Retry is the settings of retry
type Retry struct {
	Times    int           `yaml:"times"`
	Interval time.Duration `yaml:"interval"`
}

// ProbeSettings is the global probe setting
type ProbeSettings struct {
	TimeFormat string
	Interval   time.Duration
	Timeout    time.Duration
}

// NotifySettings is the global notification setting
type NotifySettings struct {
	TimeFormat string
	Retry      Retry
}