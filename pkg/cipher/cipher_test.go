package cipher

import (
	"testing"
)

func TestEcnrypt(t *testing.T) {
	key := []byte("B3FFCA612CD0C3D9050A4DE3")
	byteText := []byte("test-text")

	encodedText := Ecnrypt(key, byteText)
	decodedText := Decrypt(key, encodedText)

	if string(byteText) != decodedText {
		t.Errorf("Encrypt was incorrect, got:%s, expected:%s", encodedText, decodedText)
	}
}

func TestDecrypt(t *testing.T) {
	key := []byte("B3FFCA612CD0C3D9050A4DE3")
	byteText := []byte("test-text")
	encodedText := Ecnrypt(key, byteText)
	decodedText := Decrypt(key, encodedText)
	if string(decodedText) != string(byteText) {
		t.Errorf("Encrypt was incorrect, got:%s, expected:%s", decodedText, string(byteText))
	}
}
