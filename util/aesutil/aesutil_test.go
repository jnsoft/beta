package aesutil

import (
	"bytes"
	"crypto/rand"
	"testing"
)

func TestAesEncryptDecrypt(t *testing.T) {
	plainText := []byte("This is a secret message")
	password := "strongpassword"

	encrypted, err := AesEncrypt(plainText, password)
	if err != nil {
		t.Fatalf("AesEncrypt failed: %v", err)
	}

	decrypted, err := AesDecrypt(encrypted, password)
	if err != nil {
		t.Fatalf("AesDecrypt failed: %v", err)
	}

	if !bytes.Equal(plainText, decrypted) {
		t.Fatalf("Decrypted text does not match original. Got %s, want %s", decrypted, plainText)
	}
}

func TestAesDecryptWithWrongPassword(t *testing.T) {
	plainText := []byte("This is a secret message")
	password := "strongpassword"
	wrongPassword := "wrongpassword"

	encrypted, err := AesEncrypt(plainText, password)
	if err != nil {
		t.Fatalf("AesEncrypt failed: %v", err)
	}

	_, err = AesDecrypt(encrypted, wrongPassword)
	if err == nil {
		t.Fatalf("AesDecrypt should have failed with wrong password")
	}
}

func TestGCMEncryptDecrypt(t *testing.T) {
	key := make([]byte, 32)
	rand.Read(key)

	plain := []byte("Hello, World!")
	cipher, _ := GcmEncrypt(plain, key)
	plain_again, _ := GcmDecrypt(cipher, key)

	if !bytes.Equal(plain, plain_again) {
		t.Errorf("GCM encrypt/decrypt failed, got %s, expected %s", plain_again, plain)
	}
}
