package main

import (
	"fmt"
	"math"
	"strings"

	base64 "encoding/base64"
)

func main() {
	// Test
	fmt.Println(hashString("This is a test string."))
}

func hashString(payload string) string {
	payload = strings.ReplaceAll(payload, " ", "") // Remove spaces
	payloadBytes := []byte(payload)
	payloadLen := len(payloadBytes)

	var hashOffsetBuffer int64

	for i, v := range payloadBytes {
		var leftshiftFactor int

		endVal := payloadBytes[payloadLen-i-1]
		startVal := v

		// Circumvent reverse-engineering.
		if int(endVal) == int(startVal) {
			startVal = byte(1 + int(startVal)*i)
		}

		if endVal > startVal {
			leftshiftFactor = (int(endVal) * i) % int(startVal)
		} else {
			leftshiftFactor = (int(startVal) * i) % int(endVal)
		}
		leftshiftFactor = int(math.Abs(float64(leftshiftFactor)))

		hashOffset := math.Abs(float64((int(v) * int(endVal)) << leftshiftFactor))
		if hashOffset == 0 { // Prevent multiplying the buffer by 0 and creating a hash collision.
			hashOffset = 1
		}
		hashOffsetBuffer = int64(hashOffset) * int64(hashOffset)
	}

	encodedHash := base64.StdEncoding.EncodeToString([]byte(fmt.Sprint(hashOffsetBuffer)))
	return encodedHash
}
