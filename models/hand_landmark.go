package models

type HandLandmark struct {
	X float64
	Y float64
	Z float64
}

type HandData struct {
	Landmarks [21]HandLandmark
	Confidence float64
}
