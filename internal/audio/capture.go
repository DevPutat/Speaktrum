package audio

import (
	"github.com/gen2brain/malgo"
)

type AudioCapturer struct {
	ctx    *malgo.AllocatedContext
	device *malgo.Device
	data   chan []byte
}
