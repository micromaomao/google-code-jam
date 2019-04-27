package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func start() {
	var t int
	mustReadLineOfInts(&t)
	for i := 0; i < t; i++ {
		fmt.Fprintf(stdout, "Case #%d: ", i+1)
		debug(fmt.Sprintf("Processing case %d", i+1))
		best := test()
		if best >= 0 {
			fmt.Fprintf(stdout, "%d\n", best)
		} else {
			stdout.WriteString("IMPOSSIBLE\n")
		}
	}
}

func test() int {
	var Hd, Ad, Hk, Ak, B, D int
	mustReadLineOfInts(&Hd, &Ad, &Hk, &Ak, &B, &D)
	best := -1
	maxDebuff := 0
	if D != 0 {
		maxDebuff = Ak/D + 1
	}
	currentT := 0
	currentDebuff := 0
	currentHealth := Hd
	lastActionIsCure := false

	// This part calculates the minimum amount of attack and cures needed to defeat the knight, not counting cures.
	// It is independent of D or Hd.
	tBA, buff := optimizeBA(Hk, Ad, B)
	debug(fmt.Sprintf("tBA = %v", tBA))
	assert(tBA > buff)

	// Try every possible amount of debuff.
	for currentDebuff <= maxDebuff {
		// For each amount of debuff, simulate what amount of attacks, buffs and cures is needed, maintaining a global minimum `best`.
		debug(fmt.Sprintf("currentDebuff = %v, currentHealth = %v", currentDebuff, currentHealth))

		// This is the cures needed for the attack stage.
		// Cures needed in the debuff stage are added directly onto currentT.
		cureNeeded := 0
		impossible := false
		lastActionIsCure = false
		// Once we finish simulating this choice of debuff, we need to restore the health value.
		healthBeforeBA := currentHealth

		// Simulate attacks (adding cures when necessary), except for the last attack, which don't need to be followed by a cure.
		for i := 0; i < tBA-1; {
			// If we attacked/buffed, than...
			healthAfterThis := currentHealth - attack(Ak, currentDebuff, D)
			if healthAfterThis <= 0 {
				// Maybe we shouldn't had attacked/buffed, but cured instead.
				if lastActionIsCure {
					impossible = true
					break
				} else {
					lastActionIsCure = true
					cureNeeded++
					currentHealth = Hd - attack(Ak, currentDebuff, D) // Knight will attack us after we cured.
					if currentHealth < 0 {
						impossible = true
						break // This debuff amount isn't going to work.
					}
				}
			} else {
				lastActionIsCure = false
				currentHealth = healthAfterThis
				i++
			}
		}

		// One last attack here kills the knight, which is counted in tBA

		if impossible {
			currentHealth = healthBeforeBA
			if currentHealth-attack(Ak, currentDebuff+1, D) >= 0 {
				// We can add one more debuff and survive.
				// Not curing here would be safe, since we can cure later if needed
				currentT++
				currentDebuff++
				currentHealth = currentHealth - attack(Ak, currentDebuff, D)
				continue
			} else if Hd-attack(Ak, currentDebuff, D)-attack(Ak, currentDebuff+1, D) <= 0 { // if cure (and be attacked) and debuff (and be attacked) still result in death, than all further debuff trials will result in IMPOSSIBLE result.
				break
			} else {
				// cure
				currentT++
				currentHealth = Hd - attack(Ak, currentDebuff, D)
				currentDebuff++
				currentT++
				currentHealth -= attack(Ak, currentDebuff, D)
				continue
			}
		}
		thisTurnNumber := currentT + tBA + cureNeeded
		debug(fmt.Sprintf("thisTurnNumber = %d", thisTurnNumber))
		if best > thisTurnNumber || best == -1 {
			best = thisTurnNumber
		}

		// Try adding more debuff
		currentHealth = healthBeforeBA
		currentDebuff++
		currentT++
		if currentHealth-attack(Ak, currentDebuff, D) <= 0 {
			currentT--
			currentDebuff--
			if Hd-attack(Ak, currentDebuff, D) <= 0 {
				break // impossible from now on
			} else {
				// Cure
				currentT++
				currentHealth = Hd - attack(Ak, currentDebuff, D)
				// Debuff
				currentT++
				currentDebuff++
				currentHealth -= attack(Ak, currentDebuff, D)
				if currentHealth <= 0 {
					break
				}
			}
		} else {
			currentHealth -= attack(Ak, currentDebuff, D)
			// I initially forgot to write this else branch, which caused a lot of hard-to-debug incorrect answers and much frustration.
		}
	}
	return best
}

func attack(Ak, currentDebuff, D int) int {
	if currentDebuff*D >= Ak {
		return 0
	} else {
		return Ak - currentDebuff*D
	}
}

func optimizeBA(Hk, Ad, B int) (tBA, buff int) {
	if B == 0 {
		return int(math.Ceil(float64(Hk) / float64(Ad))), 0
	}
	b := sort.Search(Hk/B+1, func(b int) bool {
		thisB := b + int(math.Ceil(float64(Hk)/float64(Ad+B*b)))
		nextB := b + 1 + int(math.Ceil(float64(Hk)/float64(Ad+B*(b+1))))
		return nextB >= thisB
	})
	debug(fmt.Sprintf("b = %d", b))
	return b + int(math.Ceil(float64(Hk)/float64(Ad+B*b))), b
}

/*********Start boilerplate***********/

var stdin *bufio.Reader
var stdout *bufio.Writer

func main() {
	readFrom := os.Stdin
	if len(os.Args) == 2 {
		var err error
		readFrom, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
	}
	if len(os.Args) > 2 {
		panic("Too much arguments.")
	}
	stdin = bufio.NewReader(readFrom)
	stdout = bufio.NewWriter(os.Stdout)
	defer stdout.Flush()
	start()
}

func mustReadLine() string {
	str, err := stdin.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	str = strings.TrimRight(str, "\n")
	return str
}

func mustAtoi(in string) int {
	i, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return i
}

func debug(str string) {
	stdout.Flush()
	os.Stderr.WriteString("[debug] " + str + "\n")
}

func mustReadLineOfInts(a ...*int) {
	line := mustReadLine()
	strs := strings.Split(line, " ")
	if len(strs) != len(a) {
		panic("Expected " + strconv.Itoa(len(a)) + " numbers, got " + strconv.Itoa(len(strs)) + ".")
	}
	for i := 0; i < len(a); i++ {
		(*a[i]) = mustAtoi(strs[i])
	}
}

func mustReadLineOfIntsIntoArray() []int {
	line := mustReadLine()
	strs := strings.Split(line, " ")
	res := make([]int, len(strs))
	for i := 0; i < len(strs); i++ {
		res[i] = mustAtoi(strs[i])
	}
	return res
}

func assert(t bool) {
	if !t {
		panic("Assertion failed.")
	}
}
