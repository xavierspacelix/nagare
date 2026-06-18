//go:build !windows

package controller

func New() OSController {
	return NewStubController()
}
