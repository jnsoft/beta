package aesutil

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"golang.org/x/crypto/scrypt"
)

const SALT_LENGTH = 16
const SCRYPT_N = 32768

func AesEncrypt(plain []byte, password string) ([]byte, error) {
	salt := make([]byte, SALT_LENGTH)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	key, err := scrypt.Key([]byte(password), salt, SCRYPT_N, 8, 1, 32)
	if err != nil {
		return nil, err
	}

	encrypted, err := GcmEncrypt(plain, key)
	if err != nil {
		return nil, err
	}

	return append(salt, encrypted...), nil
}

func AesDecrypt(encrypted []byte, password string) ([]byte, error) {
	salt := encrypted[:SALT_LENGTH]
	key, err := scrypt.Key([]byte(password), salt, SCRYPT_N, 8, 1, 32)
	if err != nil {
		return nil, err
	}

	return GcmDecrypt(encrypted[SALT_LENGTH:], key)
}

func GcmEncrypt(plain, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	return gcm.Seal(nonce, nonce, plain, nil), nil
}

func GcmDecrypt(encrypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := encrypted[:gcm.NonceSize()]
	ciphertext := encrypted[gcm.NonceSize():]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
