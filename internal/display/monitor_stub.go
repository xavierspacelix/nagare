//go:build !windows

package display

type stubDiscoverer struct{}

func NewDiscoverer() Discoverer {
	return &stubDiscoverer{}
}

func (d *stubDiscoverer) Discover() ([]Info, error) {
	return []Info{
		{Index: 0, X: 0, Y: 0, Width: 1920, Height: 1080, Primary: true},
	}, nil
}
