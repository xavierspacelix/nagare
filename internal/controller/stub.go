package controller

type StubController struct {
	LastAction string
	LastX      int
	LastY      int
	LastTicks  int
}

func NewStubController() *StubController {
	return &StubController{}
}

func (s *StubController) MoveMouse(x, y int) error {
	s.LastAction = "move_mouse"
	s.LastX = x
	s.LastY = y
	return nil
}

func (s *StubController) LeftClick() error {
	s.LastAction = "left_click"
	return nil
}

func (s *StubController) RightClick() error {
	s.LastAction = "right_click"
	return nil
}

func (s *StubController) MouseDown() error {
	s.LastAction = "mouse_down"
	return nil
}

func (s *StubController) MouseUp() error {
	s.LastAction = "mouse_up"
	return nil
}

func (s *StubController) Scroll(ticks int) error {
	s.LastAction = "scroll"
	s.LastTicks = ticks
	return nil
}

func (s *StubController) VolumeUp() error {
	s.LastAction = "volume_up"
	return nil
}

func (s *StubController) VolumeDown() error {
	s.LastAction = "volume_down"
	return nil
}

func (s *StubController) MediaPlayPause() error {
	s.LastAction = "media_play_pause"
	return nil
}

func (s *StubController) MediaNext() error {
	s.LastAction = "media_next"
	return nil
}

func (s *StubController) MediaPrevious() error {
	s.LastAction = "media_prev"
	return nil
}
