package permissions

type Status int

const (
	StatusNotDetermined Status = iota
	StatusGranted
	StatusDenied
	StatusRestricted
	StatusUnsupported
)

type Kind string

const (
	KindCamera         Kind = "camera"
	KindAccessibility  Kind = "accessibility"
	KindInputMonitoring Kind = "input_monitoring"
)
