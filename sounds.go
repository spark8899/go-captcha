package captcha

// This file has been generated from .wav files using generate.go.

var waveHeader = []byte{
	0x52, 0x49, 0x46, 0x46, 0x00, 0x00, 0x00, 0x00, 0x57, 0x41, 0x56, 0x45,
	0x66, 0x6d, 0x74, 0x20, 0x10, 0x00, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00,
	0x40, 0x1f, 0x00, 0x00, 0x40, 0x1f, 0x00, 0x00, 0x01, 0x00, 0x08, 0x00,
	0x64, 0x61, 0x74, 0x61,
}

// Byte slices contain raw 8 kHz unsigned 8-bit PCM data (without wav header).

var digitSounds = map[string][][]byte{
	"en": [][]byte{
		{ // 0
			