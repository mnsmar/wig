// Package rand provides functions to randomize wigs.
package rand

import (
	"math/rand"
	"regexp"
)

// Perm assigns each unit of value at a wig position to a random wig position,
// independently. Positions for which f returns false are neither randomized
// nor modified.
func Perm(wig []int, f func(i int) bool) {
	wigL := len(wig)
	rwig := make([]int, wigL)
	availUnits := 0
	var valids []int

	for pos, v := range wig {
		if f != nil && f(pos) == false {
			continue
		}
		valids = append(valids, pos)
		availUnits += v
	}

	for i := 0; i < availUnits; i++ {
		j := rand.Intn(len(valids))
		rwig[valids[j]]++
	}

	for i := range wig {
		if f != nil && f(i) == false {
			continue
		}
		wig[i] = rwig[i]
	}
}

// PermPos assigns all units of value at a wig position to a random wig
// position, jointly.  Positions for which f returns false are neither
// randomized nor modified.
func PermPos(wig []int, f func(i int) bool) {
	var valids []int
	for i := 0; i < len(wig); i++ {
		if f != nil && f(i) == false {
			continue
		}
		valids = append(valids, i)
		j := rand.Intn(len(valids))
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
		if f != nil && f(pos) == false {
			continue
		}
		for j := from; j <= to; j++ {
			if pos+j < 0 || pos+j >= wigL {
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
			if len(contentAtJ) == 0 {
				pattern[j-from] = 46 // dot
				continue
			}
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
		if f != nil && f(pos) == false {
			tries++
			continue
		}
		rwig[pos]++
		availUnits--
		tries = 0
	}
	for i := range wig {
		if f != nil && f(i) == false {
			continue
		}
		wig[i] = rwig[i]
	}
}
