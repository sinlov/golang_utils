package uuidv4

import (
	"fmt"
	"regexp"
	"time"
	"crypto/rand"
	mRand "math/rand"
	"encoding/hex"
	"strings"
	"errors"
)

// seeded indicates if math/rand has been seeded
var seeded bool = false

// uuidRegex matches the UUID string
var uuidRegex *regexp.Regexp = regexp.MustCompile(`^\{?([a-fA-F0-9]{8})-?([a-fA-F0-9]{4})-?([a-fA-F0-9]{4})-?([a-fA-F0-9]{4})-?([a-fA-F0-9]{12})\}?$`)

// UUID type.
type UUIDV4 [16]byte

// Hex string returns a hex string representation of the UUID v4 in xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx format.
func (uuid UUIDV4) HexV4() string {
	x := [16]byte(uuid)
	return fmt.Sprintf("%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		x[0], x[1], x[2], x[3], x[4],
		x[5], x[6],
		x[7], x[8],
		x[9], x[10], x[11], x[12], x[13], x[14], x[15])

}

func (uuid UUIDV4) RawV4() [16]byte {
	return [16]byte(uuid)
}

// Rand generates a new version 4 UUID.
func RandV4() UUIDV4 {
	var x [16]byte
	randBytes(x[:])
	x[6] = (x[6] & 0x0F) | 0x40
	x[8] = (x[8] & 0x3F) | 0x80
	return x
}

// FromStr returns a UUID based on a string.
// The string could be in the following format:
//
// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
//
// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
//
// {xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}
//
// If the string is not in one of these formats, it'll return an error.
func FromStrV4(s string) (id UUIDV4, err error) {
	if s == "" {
		err = errors.New("Empty string")
		return
	}

	parts := uuidRegex.FindStringSubmatch(s)
	if parts == nil {
		err = errors.New("Invalid string format")
		return
	}

	var array [16]byte
	slice, _ := hex.DecodeString(strings.Join(parts[1:], ""))
	copy(array[:], slice)
	id = array
	return
}

// MustFromStr behaves similarly to FromStr except that it'll panic instead of
// returning an error.
func MustFromStrV4(s string) UUIDV4 {
	id, err := FromStrV4(s)
	if err != nil {
		panic(err)
	}
	return id
}

// randBytes uses crypto random to get random numbers. If fails then it uses math random.
func randBytes(x []byte) {

	length := len(x)
	n, err := rand.Read(x)

	if n != length || err != nil {
		if !seeded {
			mRand.Seed(time.Now().UnixNano())
		}

		for length > 0 {
			length--
			x[length] = byte(mRand.Int31n(256))
		}
	}
}
