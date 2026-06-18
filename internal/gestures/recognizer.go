package gestures

import (
	"math"

	"nagare/models"
)

type Recognizer struct {
	palmCenter models.HandLandmark
}

func NewRecognizer() *Recognizer {
	return &Recognizer{}
}

func (r *Recognizer) IsOpenPalm(data *models.HandData) bool {
	if data == nil || data.Confidence < 0.5 {
		return false
	}
	return data.Fingers.Thumb() == models.FingerExtended &&
		data.Fingers.Index() == models.FingerExtended &&
		data.Fingers.Middle() == models.FingerExtended &&
		data.Fingers.Ring() == models.FingerExtended &&
		data.Fingers.Pinky() == models.FingerExtended
}

func (r *Recognizer) IsClosedFist(data *models.HandData) bool {
	if data == nil || data.Confidence < 0.5 {
		return false
	}
	return data.Fingers.Thumb() == models.FingerFolded &&
		data.Fingers.Index() == models.FingerFolded &&
		data.Fingers.Middle() == models.FingerFolded &&
		data.Fingers.Ring() == models.FingerFolded &&
		data.Fingers.Pinky() == models.FingerFolded
}

func (r *Recognizer) IsPinch(data *models.HandData) bool {
	if data == nil || data.Confidence < 0.5 {
		return false
	}
	return fingerDistance(data, 4, 8) < 0.05
}

func (r *Recognizer) IsTwoFingerPinch(data *models.HandData) bool {
	if data == nil || data.Confidence < 0.5 {
		return false
	}
	return fingerDistance(data, 4, 12) < 0.05
}

func (r *Recognizer) IsTwoFingersUp(data *models.HandData) bool {
	if data == nil || data.Confidence < 0.5 {
		return false
	}
	return data.Fingers.Index() == models.FingerExtended &&
		data.Fingers.Middle() == models.FingerExtended &&
		data.Fingers.Ring() == models.FingerFolded &&
		data.Fingers.Pinky() == models.FingerFolded
}

func (r *Recognizer) IsTwoFingersDown(data *models.HandData) bool {
	return r.IsTwoFingersUp(data)
}

func (r *Recognizer) IsSwipeLeft(data *models.HandData) bool {
	if data == nil || data.Confidence < 0.5 {
		return false
	}
	palm := handCenter(&data.Landmarks)
	if r.palmCenter.X == 0 && r.palmCenter.Y == 0 {
		r.palmCenter = palm
		return false
	}
	dx := palm.X - r.palmCenter.X
	r.palmCenter = palm
	return dx < -20
}

func (r *Recognizer) IsSwipeRight(data *models.HandData) bool {
	if data == nil || data.Confidence < 0.5 {
		return false
	}
	palm := handCenter(&data.Landmarks)
	if r.palmCenter.X == 0 && r.palmCenter.Y == 0 {
		r.palmCenter = palm
		return false
	}
	dx := palm.X - r.palmCenter.X
	r.palmCenter = palm
	return dx > 20
}

func (r *Recognizer) ResetSwipe() {
	r.palmCenter = models.HandLandmark{}
}

func fingerDistance(data *models.HandData, i, j int) float64 {
	a := data.Landmarks[i]
	b := data.Landmarks[j]
	dx := a.X - b.X
	dy := a.Y - b.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func handCenter(landmarks *[21]models.HandLandmark) models.HandLandmark {
	var cx, cy float64
	for _, lm := range landmarks {
		cx += lm.X
		cy += lm.Y
	}
	n := float64(len(landmarks))
	return models.HandLandmark{X: cx / n, Y: cy / n}
}
