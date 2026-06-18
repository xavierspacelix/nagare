//go:build windows || darwin
// +build windows darwin

package camera

import (
	"fmt"

	"gocv.io/x/gocv"
)

type openCVCam struct {
	cam *gocv.VideoCapture
}

func newDevice() CamDevice {
	return &openCVCam{}
}

func discoverCameras() []Info {
	var list []Info
	for i := 0; i < 10; i++ {
		cam, err := gocv.OpenVideoCapture(i)
		if err != nil || !cam.IsOpened() {
			continue
		}
		cam.Close()
		list = append(list, Info{
			ID:   i,
			Name: fmt.Sprintf("Camera %d", i),
		})
	}
	if len(list) == 0 {
		list = []Info{{ID: 0, Name: "Default Camera"}}
	}
	return list
}

func (d *openCVCam) Open(id int) error {
	cam, err := gocv.OpenVideoCapture(id)
	if err != nil {
		return fmt.Errorf("open video capture: %w", err)
	}
	if !cam.IsOpened() {
		return fmt.Errorf("camera %d could not be opened", id)
	}
	d.cam = cam
	return nil
}

func (d *openCVCam) Read() (*Frame, error) {
	if d.cam == nil {
		return nil, fmt.Errorf("camera not opened")
	}
	mat := gocv.NewMat()
	defer mat.Close()

	if !d.cam.Read(&mat) {
		return nil, fmt.Errorf("failed to read frame")
	}

	return &Frame{
		Data:   mat.ToBytes(),
		Width:  mat.Cols(),
		Height: mat.Rows(),
	}, nil
}

func (d *openCVCam) Close() error {
	if d.cam != nil {
		d.cam.Close()
		d.cam = nil
	}
	return nil
}
