package gen

import (
	"crypto/rand"
	"errors"
	"math/big"
)

var alphabet = map[string]bool{
	"a": true, "A": true,
	"b": true, "B": true,
	"c": true, "C": true,
	"d": true, "D": true,
	"e": true, "E": true,
	"f": true, "F": true,
	"g": true, "G": true,
	"h": true, "H": true,
	"i": true, "I": true,
	"j": true, "J": true,
	"k": true, "K": true,
	"l": true, "L": true,
	"m": true, "M": true,
	"n": true, "N": true,
	"o": true, "O": true,
	"p": true, "P": true,
	"q": true, "Q": true,
	"r": true, "R": true,
	"s": true, "S": true,
	"t": true, "T": true,
	"u": true, "U": true,
	"v": true, "V": true,
	"w": true, "W": true,
	"x": true, "X": true,
	"y": true, "Y": true,
	"z": true, "Z": true,
}

var digit = map[string]bool{
	"0": true, "1": true, "2": true, "3": true, "4": true,
	"5": true, "6": true, "7": true, "8": true, "9": true,
}

var special = map[string]bool{
	"~": true, "!": true, "@": true, "#": true, "$": true, "%": true, "^": true, "&": true, "*": true, "(": true, ")": true,
	"-": true, "_": true, "+": true, "=": true, "{": true, "}": true, "[": true, "]": true, "\\": true, "|": true, ":": true,
	";": true, "\"": true, "'": true, ",": true, "<": true, ".": true, ">": true, "?": true, "/": true, "`": true,
}

var set = map[string]map[string]bool{
	"alphabet": alphabet,
	"digit":    digit,
	"special":  special,
	"word":     word,
	"space":    map[string]bool{" ": true},
}

var setName = func() map[string]string {
	m := map[string]string{}
	for i := 'A'; i < 'Z'; i++ {
		m[string(i)] = "alphabet"
	}
	for i := 'a'; i < 'z'; i++ {
		m[string(i)] = "alphabet"
	}
	for i := '0'; i < '9'; i++ {
		m[string(i)] = "digit"
	}
	for i := '!'; i < '/'; i++ {
		m[string(i)] = "special"
	}
	for i := ':'; i < '@'; i++ {
		m[string(i)] = "special"
	}
	for i := '['; i < '`'; i++ {
		m[string(i)] = "special"
	}
	for i := '{'; i < '~'; i++ {
		m[string(i)] = "special"
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
	valid := merge(c.CharSets)
	return randString(valid, c.Length)
}

func merge(ss map[string]bool) []string {
	sets := make([]map[string]bool, len(ss))
	var length int
	for name := range ss {
		s := set[name]
		sets = append(sets, s)
		length += len(s)
	}
	list := make([]string, 0, length)
	for _, m := range sets {
		for s := range m {
			list = append(list, s)
		}
	}
	return list
}

func randString(list []string, length int) (string, error) {
	var ss string
	for i := 0; i < length; i++ {
		s, err := randElem(list)
		if err != nil {
			return "", err
		}
		ss += s
	}
	return ss, nil
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
