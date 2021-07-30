package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

/**
Generates a random Cursor Park, so that Cursors in your
document can feel safe and secure, while enjoying nature.

Cursor Park & Nature Preserve
🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳
🌳              🐇                🐢  🌳
🌳                                  🦃🌳
🌳  🦍              🌲          🐍    🌳
🌳                  🐘                🌳
🌳                                    🌳
🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳
**/

const (
	tree  = "🌳"
	blank = "⬜"
)

func main() {
	if err := inner(); err != nil {
		fmt.Printf("cursorpark: %s", err)
		os.Exit(1)
	}
}

var (
	width, height    int
	critterFrequency float64
)

func inner() error {
	// Set the size of the Cursor Park
	flag.IntVar(&width, "width", 20, "the width of the cursor park")
	flag.IntVar(&width, "w", 20, "the width of the cursor park (shorthand)")
	flag.IntVar(&height, "height", 7, "the height of the cursor park")
	flag.IntVar(&height, "h", 7, "the height of the cursor park (shorthand)")

	// Prepare to randomly generate critters
	rand.Seed(time.Now().UTC().UnixNano())
	flag.Float64Var(&critterFrequency, "freq", 0.07, "how often a critter appears")
	flag.Float64Var(&critterFrequency, "f", 0.07, "how often a critter appears (shorthand)")

	flag.Parse()

	// Generate the top and bottom tree row
	top := strings.Repeat(tree, width)
	bottom := top

	// Based on the desired height, generate a row starting and
	// ending with a tree, with either blank or a critter in each
	// space.
	rows := []string{}
	for i := 0; i < (height - 2); i++ {
		rows = append(rows, genRow(width, critterFrequency))
	}

	// Print our Cursor Park to the Terminal
	fmt.Println("Cursor Park & Nature Preserve")
	fmt.Println(top)
	for _, r := range rows {
		fmt.Println(r)
	}
	fmt.Println(bottom)

	return nil
}

// genRow generates a row of width, bookended by 🌳 with a
// critter in each space with a probablity of critterFrequency.
// For example:
//   `🌳        🐖                🐁        🌳`
func genRow(width int, critterFrequency float64) string {
	critters := critters()
	var str strings.Builder

	str.WriteString(tree)

	for i := 0; i < (width - 2); i++ {
		r := rand.Float64()

		space := blank
		if critterFrequency > r {
			space = critters[rand.Intn(len(critters))]
		}

		str.WriteString(space)
	}

	str.WriteString(tree)
	return str.String()
}

// critters returns a slice of animals and plants. Note that some emoji
// are not the same width as others, but these all are the same width as
// 🌳 anything added to this function must be as well.
func critters() []string {
	return []string{
		"🐍",
		"🐖",
		"🌴",
		"🐝",
		"🐢",
		"🐇",
		"🐐",
		"🐛",
		"🦍",
		"🐘",
		"🌵",
		"🌲",
		"🦀",
		"🐌",
		"🦃",
		"🐈",
		"🐕",
		"🐁",
		"🐑",
	}
}
