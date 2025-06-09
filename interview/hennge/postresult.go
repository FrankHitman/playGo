// Then, make an HTTP POST request to the following URL with the JSON string as the body part.
// https://api.challenge.hennge.com/challenges/003

// Content type
// The Content-Type: of the request must be application/json.

// Authorization
// The URL is protected by HTTP Basic Authentication, which is explained on Chapter 2 of RFC2617, so you have to provide an Authorization: header field in your POST request

// For the userid of HTTP Basic Authentication, use the same email address you put in the JSON string.
// For the password, provide a 10-digit time-based one time password conforming to RFC6238 TOTP.
// Authorization password

// For generating the TOTP password, you will need to use the following setup:

// You have to read RFC6238 (and the errata too!) and get a correct one time password by yourself.
// TOTP's Time Step X is 30 seconds. T0 is 0.
// Use HMAC-SHA-512 for the hash function, instead of the default HMAC-SHA-1.
// Token shared secret is the userid followed by ASCII string value "HENNGECHALLENGE003" (not including double quotations).

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/binary"
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

	// Extract the dynamic binary code from the hash
	offset := hash[len(hash)-1] & 0x0F
	code := (int(hash[offset])&0x7F)<<24 |
		(int(hash[offset+1])&0xFF)<<16 |
		(int(hash[offset+2])&0xFF)<<8 |
		(int(hash[offset+3])&0xFF)

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

	// Convert the JSON data to a byte array
	jsonData, err := json.Marshal(solJson)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "https://api.challenge.hennge.com/challenges/003", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the headers
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(solJson["contact_email"], totp)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Print the response status
	fmt.Println("Response Status:", resp.Status)
}
