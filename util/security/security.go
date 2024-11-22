package security

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

// Hash functions

func HashMD5(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

func HashSHA1(data []byte) string {
	hash := sha1.Sum(data)
	return hex.EncodeToString(hash[:])
}

func HashSHA256(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

func HashSHA512(data []byte) string {
	hash := sha512.Sum512(data)
	return hex.EncodeToString(hash[:])
}

// Keccak algorithm
func HashSHA3(data []byte) string {
	hash := sha3.Sum256(data)
	return hex.EncodeToString(hash[:])
}

// HMAC functions

func HmacSHA256_hex(data, key []byte) string {
	return hex.EncodeToString(HmacSHA256(data, key))
}

func HmacSHA256(data, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	hmac := h.Sum(nil)
	return hmac
}

func HmacSHA256_verify_hex(data, key []byte, hex_hmac string) bool {
	expectedHMAC, _ := hex.DecodeString(hex_hmac)
	return HmacSHA256_verify(data, key, expectedHMAC)
}

func HmacSHA256_verify(data, key, expectedHMAC []byte) bool {
	computedHMAC := HmacSHA256(data, key)
	return hmac.Equal(computedHMAC, expectedHMAC)
}

func HmacSHA512_hex(data, key []byte) string {
	return hex.EncodeToString(HmacSHA512(data, key))
}

func HmacSHA512(data, key []byte) []byte {
	h := hmac.New(sha512.New, key)
	h.Write(data)
	hmac := h.Sum(nil)
	return hmac
}

func HmacSHA512_verify_hex(data, key []byte, hex_hmac string) bool {
	expectedHMAC, _ := hex.DecodeString(hex_hmac)
	return HmacSHA512_verify(data, key, expectedHMAC)
}

func HmacSHA512_verify(data, key, expectedHMAC []byte) bool {
	computedHMAC := HmacSHA512(data, key)
	return hmac.Equal(computedHMAC, expectedHMAC)
}

func HmacSHA3_hex(data, key []byte) string {
	return hex.EncodeToString(HmacSHA3(data, key))
}

func HmacSHA3(data, key []byte) []byte {
	h := hmac.New(sha3.New256, key)
	h.Write(data)
	hmac := h.Sum(nil)
	return hmac
}

func HmacSHA3_verify_hex(data, key []byte, hex_hmac string) bool {
	expectedHMAC, _ := hex.DecodeString(hex_hmac)
	return HmacSHA3_verify(data, key, expectedHMAC)
}

func HmacSHA3_verify(data, key, expectedHMAC []byte) bool {
	computedHMAC := HmacSHA3(data, key)
	return hmac.Equal(computedHMAC, expectedHMAC)
}
