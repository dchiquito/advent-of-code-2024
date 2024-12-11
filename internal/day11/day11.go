package day11

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/dchiquito/advent-of-code-2024/internal/util"
)

func parse(in io.Reader) []int {
	bytes, _ := io.ReadAll(in)
	line := string(bytes[:len(bytes)-1])
	seedStrings := strings.Split(line, " ")
	seeds := make([]int, len(seedStrings))
	for i, seed := range seedStrings {
		seeds[i], _ = strconv.Atoi(seed)
	}
	return seeds
}

func blinkOnce(rock int) (int, int) {
	if rock == 0 {
		return 1, -1
	}
	pow := 10
	half := 10
	for true {
		if rock < pow {
			return rock * 2024, -1
		}
		if rock < 10*pow {
			return rock / half, rock % half
		}
		pow *= 100
		half *= 10
	}
	util.Panic("unreachable")
	return -1, -1
}

func blinkRepeatedly(cache *[]int, rock int, blinks int) int {
	if blinks == 0 {
		return 1
	}
	if rock < CACHE_SIZE && (*cache)[rock*MAX_BLINKS+blinks] != 0 {
		return (*cache)[rock*MAX_BLINKS+blinks]
	}
	a, b := blinkOnce(rock)
	total := blinkRepeatedly(cache, a, blinks-1)
	if b != -1 {
		total += blinkRepeatedly(cache, b, blinks-1)
	}
	if rock < CACHE_SIZE {
		(*cache)[rock*MAX_BLINKS+blinks] = total
	}
	return total
}

func printCache(cache *[]int, maxBlinks int) {
	for i, c := range *cache {
		if c > 0 {
			rock := i / maxBlinks
			blinks := i % maxBlinks
			fmt.Printf("%d{%d}=%d ", rock, blinks, c)
		}
	}
}

const CACHE_SIZE = 100
const MAX_BLINKS = 25

func Level1(in io.Reader) string {
	seeds := parse(in)
	cache := make([]int, CACHE_SIZE*MAX_BLINKS)
	total := 0
	for _, seed := range seeds {
		// TODO this throws away a substantial amount of work
		// blinking thrice at 0 gives (1), (2024), (20, 24)
		// This approach will cache blinking twice at 1 and blinking once at 2024,
		// but will not cache blinking once or twice at 0.
		// Building up a stack might make caching up to twice as efficient?
		total += blinkRepeatedly(&cache, seed, 25)
	}
	return fmt.Sprint(total)
}

const CACHE_SIZE2 = 1200
const MAX_BLINKS2 = 75

func blinkRepeatedly2(cache *[]int, rock int, blinks int) int {
	if blinks == 0 {
		return 1
	}
	if rock < CACHE_SIZE2 && (*cache)[rock*MAX_BLINKS2+blinks] != 0 {
		return (*cache)[rock*MAX_BLINKS2+blinks]
	}
	a, b := blinkOnce(rock)
	total := blinkRepeatedly2(cache, a, blinks-1)
	if b != -1 {
		total += blinkRepeatedly2(cache, b, blinks-1)
	}
	if rock < CACHE_SIZE2 {
		(*cache)[rock*MAX_BLINKS2+blinks] = total
	}
	return total
}

func Level2(in io.Reader) string {
	seeds := parse(in)
	cache := make([]int, CACHE_SIZE2*MAX_BLINKS2)
	total := 0
	for _, seed := range seeds {
		total += blinkRepeatedly2(&cache, seed, 75)
	}
	return fmt.Sprint(total)
}
