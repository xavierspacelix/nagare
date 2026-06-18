//go:build darwin

package permissions

/*
#cgo LDFLAGS: -framework AVFoundation -framework ApplicationServices -framework Foundation
#include <objc/runtime.h>
#include <objc/message.h>
#include <ApplicationServices/ApplicationServices.h>

int cameraPermissionStatus(void) {
    Class avClass = objc_getClass("AVCaptureDevice");
    if (avClass == NULL) return -1;

    SEL authSel = sel_registerName("authorizationStatusForMediaType:");
    SEL stringSel = sel_registerName("stringWithUTF8String:");
    Class nsStr = objc_getClass("NSString");

    id mediaType = ((id(*)(id, SEL, const char*))objc_msgSend)((id)nsStr, stringSel, "AVMediaTypeVideo");
    int status = ((int(*)(id, SEL, id))objc_msgSend)((id)avClass, authSel, mediaType);
    return status;
}

int accessibilityPermissionStatus(void) {
    return AXIsProcessTrusted() ? 1 : 0;
}
*/
import "C"

func Check(kind Kind) Status {
	switch kind {
	case KindCamera:
		return convertCameraStatus(C.cameraPermissionStatus())
	case KindAccessibility, KindInputMonitoring:
		return convertAccessibilityStatus(C.accessibilityPermissionStatus())
	default:
		return StatusUnsupported
	}
}

func convertCameraStatus(s C.int) Status {
	switch s {
	case -1:
		return StatusUnsupported
	case 0:
		return StatusNotDetermined
	case 1:
		return StatusRestricted
	case 2:
		return StatusDenied
	case 3:
		return StatusGranted
	default:
		return StatusUnsupported
	}
}

func convertAccessibilityStatus(s C.int) Status {
	if s == 1 {
		return StatusGranted
	}
	return StatusDenied
}

func AllGranted() bool {
	return Check(KindCamera) == StatusGranted &&
		Check(KindAccessibility) == StatusGranted
}

func StatusText(kind Kind) string {
	switch Check(kind) {
	case StatusGranted:
		return "Granted"
	case StatusDenied:
		return "Denied"
	case StatusRestricted:
		return "Restricted"
	case StatusNotDetermined:
		return "Not Determined"
	default:
		return "Unsupported"
	}
}

func Request(kind Kind) bool {
	return Check(kind) == StatusGranted
}
