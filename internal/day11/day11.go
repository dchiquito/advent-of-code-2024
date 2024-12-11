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

func printCache(cache *[]int) {
	for i, c := range *cache {
		if c > 0 {
			rock := i / MAX_BLINKS
			blinks := i % MAX_BLINKS
			fmt.Printf("%d{%d}=%d ", rock, blinks, c)
		}
	}
}

const CACHE_SIZE = 100
const MAX_BLINKS = 75 // TODO split this between levels

func Level1(in io.Reader) string {
	seeds := parse(in)
	cache := make([]int, CACHE_SIZE*MAX_BLINKS)
	total := 0
	for _, seed := range seeds {
		total += blinkRepeatedly(&cache, seed, 25)
	}
	return fmt.Sprint(total)
}

func Level2(in io.Reader) string {
	seeds := parse(in)
	cache := make([]int, CACHE_SIZE*MAX_BLINKS)
	total := 0
	for _, seed := range seeds {
		total += blinkRepeatedly(&cache, seed, 75)
	}
	return fmt.Sprint(total)
}
