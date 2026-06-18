package models

type HandLandmark struct {
	X float64
	Y float64
	Z float64
}

type HandOrientation int

const (
	HandOrientationUnknown HandOrientation = 0
	HandOrientationLeft    HandOrientation = 1
	HandOrientationRight   HandOrientation = 2
)

type FingerName int

const (
	FingerThumb   FingerName = 0
	FingerIndex   FingerName = 1
	FingerMiddle  FingerName = 2
	FingerRing    FingerName = 3
	FingerPinky   FingerName = 4
)

type FingerState int

const (
	FingerUnknown  FingerState = 0
	FingerExtended FingerState = 1
	FingerFolded   FingerState = 2
)

type FingerStates [5]FingerState

func (f FingerStates) Thumb() FingerState  { return f[FingerThumb] }
func (f FingerStates) Index() FingerState  { return f[FingerIndex] }
func (f FingerStates) Middle() FingerState { return f[FingerMiddle] }
func (f FingerStates) Ring() FingerState   { return f[FingerRing] }
func (f FingerStates) Pinky() FingerState  { return f[FingerPinky] }

var LandmarkNames = [21]string{
	"wrist",
	"thumb_cmc",
	"thumb_mcp",
	"thumb_ip",
	"thumb_tip",
	"index_mcp",
	"index_pip",
	"index_dip",
	"index_tip",
	"middle_mcp",
	"middle_pip",
	"middle_dip",
	"middle_tip",
	"ring_mcp",
	"ring_pip",
	"ring_dip",
	"ring_tip",
	"pinky_mcp",
	"pinky_pip",
	"pinky_dip",
	"pinky_tip",
}

type HandData struct {
	Landmarks   [21]HandLandmark
	Confidence  float64
	Orientation HandOrientation
	Fingers     FingerStates
}
