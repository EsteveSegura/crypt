package cipher

import "testing"

func TestEcnrypt(t *testing.T) {
	plainText := []byte("test-text")
	expectedResult := "dGVzdC10ZXh0"

	encodedText := Ecnrypt(plainText)

	if encodedText != expectedResult {
		t.Errorf("Encrypt was incorrect, got:%s, expected:%s", encodedText, expectedResult)
	}
}

func TestDecrypt(t *testing.T) {
	encodedText := []byte("dGVzdC10ZXh0")
	expectedResult := "test-text"

	decodedText := Decrypt(string(encodedText))

	if string(decodedText) != string(expectedResult) {
		t.Errorf("Encrypt was incorrect, got:%s, expected:%s", decodedText, expectedResult)
	}
}
