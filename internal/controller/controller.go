//go:build !windows && !darwin

package controller

func New() OSController {
	return NewStubController()
}
