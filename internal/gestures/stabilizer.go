package gestures

import (
	"nagare/models"
)

type Stabilizer struct {
	history    map[models.Gesture]*gestureBuffer
	bufferSize int
}

type gestureBuffer struct {
	detections []bool
	count      int
}

func NewStabilizer(bufferSize int) *Stabilizer {
	if bufferSize < 1 {
		bufferSize = 1
	}
	return &Stabilizer{
		history:    make(map[models.Gesture]*gestureBuffer),
		bufferSize: bufferSize,
	}
}

func (s *Stabilizer) Record(gesture models.Gesture, detected bool) bool {
	buf, exists := s.history[gesture]
	if !exists {
		buf = &gestureBuffer{
			detections: make([]bool, 0, s.bufferSize),
		}
		s.history[gesture] = buf
	}

	if len(buf.detections) >= s.bufferSize {
		if buf.detections[0] {
			buf.count--
		}
		buf.detections = buf.detections[1:]
	}

	buf.detections = append(buf.detections, detected)
	if detected {
		buf.count++
	}

	return s.isStable(buf)
}

func (s *Stabilizer) isStable(buf *gestureBuffer) bool {
	if len(buf.detections) < s.bufferSize {
		return false
	}
	threshold := (s.bufferSize * 2) / 3
	if threshold < 1 {
		threshold = 1
	}
	return buf.count >= threshold
}

func (s *Stabilizer) Reset(gesture models.Gesture) {
	delete(s.history, gesture)
}

func (s *Stabilizer) ResetAll() {
	s.history = make(map[models.Gesture]*gestureBuffer)
}
