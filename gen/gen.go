package gen

import (
	"bytes"
	"crypto/rand"
	"errors"
	"math/big"
)

var (
	alphabet = []string{"a", "A", "b", "B", "c", "C", "d", "D", "e", "E", "f", "F", "g", "G", "h", "H", "i", "I", "j", "J", "k", "K", "l", "L", "m", "M", "n", "N", "o", "O", "p", "P", "q", "Q", "r", "R", "s", "S", "t", "T", "u", "U", "v", "V", "w", "W", "x", "X", "y", "Y", "z", "Z"}
	digit    = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	special  = []string{"~", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "+", "=", "{", "}", "[", "]", "\\", "|", ":", ";", "\"", "'", ",", "<", ".", ">", "?", "/", "`"}
)

var set = map[string][]string{
	"alphabet": alphabet,
	"digit":    digit,
	"special":  special,
	"word":     word,
	"space":    []string{" "},
}

var setName = func() map[string]string {
	m := map[string]string{}
	for _, s := range alphabet {
		m[s] = "alphabet"
	}
	for _, s := range digit {
		m[s] = "digit"
	}
	for _, s := range special {
		m[s] = "special"
	}
	m["word"] = "word"
	m[" "] = "space"
	return m
}()

// Config is used to set the config for generating a random password
type Config struct {
	// CharSets example: map[string]bool{"alphabet":true,"digit":true}
	// CharSets specifies which sets of characters to use when generating a
	// random password
	// Available CharSets
	//   "alphabet"
	//   "digit"
	//   "special"
	//   "space"
	//   "word"
	CharSets map[string]bool

	// Length specifies how long the generated password(s) will be
	Length int
}

// CharSets returns a map of the character sets that appear in the input string
func CharSets(s string) map[string]bool {
	m := map[string]bool{}
	for _, c := range s {
		m[setName[string(c)]] = true
	}
	return m
}

// Gen generates a random string based on the Config
func (c *Config) Gen() (string, error) {
	var valid []string
	for name := range c.CharSets {
		valid = append(valid, set[name]...)
	}
	return randString(valid, c.Length)
}

func randString(list []string, length int) (string, error) {
	b := bytes.NewBufferString("")
	for i := 0; i < length; i++ {
		s, err := randElem(list)
		if err != nil {
			return "", err
		}
		b.WriteString(s)
	}
	return b.String(), nil
}

func randElem(list []string) (string, error) {
	if len(list) == 0 {
		return "", errors.New("element list is empty")
	}
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(list))))
	if err != nil {
		return "", err
	}
	return list[int(n.Int64())], nil
}
