package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/binary"
	"fmt"
	"time"
)

func generateTOTP(secret string) string {
	// Get the current Unix time in seconds
	currentTime := time.Now().Unix()
	// Convert the time to a 30-second time step
	timeStep := currentTime / 30

	// Convert the time step to a byte array
	timeBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(timeBytes, uint64(timeStep))

	// Create an HMAC-SHA-512 hash using the secret key and the time step
	hmacHash := hmac.New(sha512.New, []byte(secret))
	hmacHash.Write(timeBytes)
	hash := hmacHash.Sum(nil)
	fmt.Println(hash)
	// Extract the dynamic binary code from the hash
	offset := hash[len(hash)-1] & 0x0F
	code := (int(hash[offset])&0x7F)<<24 |
		(int(hash[offset+1])&0xFF)<<16 |
		(int(hash[offset+2])&0xFF)<<8 |
		(int(hash[offset+3])&0xFF)

	fmt.Println(code)
	// Convert the code to a 10-digit string
	totp := fmt.Sprintf("%010d", code%10000000000)

	return totp
}

func main() {
	solJson := map[string]string {
		"github_url": "https://gist.github.com/FrankHitman/22f24a39d3ecc9171ce0617715722730",
		"contact_email": "dong.y.sun@outlook.com",
		"solution_language": "golang",
	  }
	secret := solJson["contact_email"]+"HENNGECHALLENGE003"
	totp := generateTOTP(secret)
	fmt.Println("TOTP:", totp)
}
