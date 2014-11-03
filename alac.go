package alac

import (
	// #include "alac.h"
	"C"
	"unsafe"
)

// Wrapper type around C type.
type AlacFile struct {
	file *C.alac_file
}

// Allocates new AlacFile.
func New(sampleSize, numberOfChannels int) AlacFile {
	alac := AlacFile{}
	alac.file = C.alac_create(C.int(sampleSize), C.int(numberOfChannels))
	return alac
}

// Decodes a frame from inputBuffer and puts it in the outputBuffer.
// Might make sense to change this API to return the output buffer instead.
func (f *AlacFile) DecodeFrame(inputBuffer, outputBuffer []byte) {
	size := C.int(len(outputBuffer))
	C.alac_decode_frame(f.file, (*C.uchar)(unsafe.Pointer(&inputBuffer)), (unsafe.Pointer(&outputBuffer)), &size)
}

// Set's the "info" for our AlacFile.
func (f *AlacFile) SetInfo(inputBuffer []byte) {
	C.alac_set_info(f.file, (*C.char)(unsafe.Pointer(&inputBuffer)))
}

// Allocates the C buffers for our AlacFile.
func (f *AlacFile) AllocateBuffers() {
	C.alac_allocate_buffers(f.file)
}

// Free's the C buffers we wrap in our AlacFile type.
func (f *AlacFile) Free() {
	C.alac_free(f.file)
}
