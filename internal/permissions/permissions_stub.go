//go:build !darwin

package permissions

func Check(kind Kind) Status {
	return StatusUnsupported
}

func Request(kind Kind) bool {
	return false
}

func AllGranted() bool {
	return false
}

func StatusText(kind Kind) string {
	return "Unsupported"
}
