package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/voutasaurus/p/gen"
)

var (
	flagChar = flag.String("c", "a1", "character sets to include in generated password")
	flagLen  = flag.Int("l", 40, "length of password to generate")
	flagWord = flag.Bool("w", false, "generate list of random words instead of characters")
)

func main() {
	flag.Parse()

	sets := map[string]bool{"word": true}
	if !*flagWord {
		sets = gen.CharSets(*flagChar)
	}

	c := gen.Config{
		CharSets: sets,
		Length:   *flagLen,
	}
	s, err := c.Gen()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(s)
}
