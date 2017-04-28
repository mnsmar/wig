// Package rand provides types and functions to randomize wigs.
package rand

import (
	"math/rand"
	"regexp"
)

// Perm assigns each unit of value at a wig position to a random wig position.
func Perm(wig []int) {
	wigL := len(wig)
	rwig := make([]int, wigL)
	for _, v := range wig {
		for i := 0; i < v; i++ {
			j := rand.Intn(wigL)
			rwig[j]++
		}
	}
	copy(wig, rwig)
}

// PermPos assigns all units at a wig position to a random wig position.
func PermPos(wig []int) {
	for i := 0; i < len(wig); i++ {
		j := rand.Intn(i + 1)
		wig[i], wig[j] = wig[j], wig[i]
	}
}

// PermKeepByte assigns each unit of value at a wig position to a random wig
// position maintaining the byte content (extracted from s) in the region
// [from, to] around the wig position. Positions for which f returns false are
// neither randomized nor modified.
func PermKeepByte(wig []int, s []byte, from, to int, f func(i int) bool) {
	wigL := len(wig)
	rwig := make([]int, wigL)
	availUnits := 0

	content := make(map[int][]byte)
	for pos, units := range wig {
		if units == 0 {
			continue
		}
		if f(pos) == false {
			continue
		}
		for j := from; j <= to; j++ {
			if pos+j < 0 || pos+j > wigL {
				continue
			}
			b := s[pos+j]
			for k := 0; k < units; k++ {
				content[j] = append(content[j], b)
			}
		}
		availUnits += units
	}

	var tries int
	for availUnits > 0 && tries < 1000 {
		pattern := make([]byte, to-from+1)
		for j := from; j <= to; j++ {
			contentAtJ := content[j]
			pattern[j-from] = contentAtJ[rand.Intn(len(contentAtJ))]
		}
		re := regexp.MustCompile(string(pattern))
		matches := re.FindAllIndex(s, -1)
		if matches == nil {
			tries++
			continue
		}
		randMatch := matches[rand.Intn(len(matches))]
		pos := randMatch[0] - from
		if f(pos) == false {
			tries++
			continue
		}
		rwig[pos]++
		availUnits--
		tries = 0
	}
	for i := range wig {
		if f(i) == false {
			continue
		}
		wig[i] = rwig[i]
	}
}
