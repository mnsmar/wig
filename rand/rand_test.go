package rand

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

var wigs = []struct {
	name string
	wig  []int
	sum  int
}{
	{
		name: "simple wig",
		wig:  []int{10, 5, 0, 0, 0, 0},
		sum:  15,
	},
}

func TestPerm(t *testing.T) {
	for _, tt := range wigs {
		w := make([]int, len(tt.wig))
		copy(w, tt.wig)
		Perm(w)
		sum := 0
		for _, v := range w {
			sum += v
		}
		if sum != tt.sum {
			t.Errorf("%s: wrong sum: expected %d, actual %d", tt.name, tt.sum, sum)
		}
	}
}

func TestPermPos(t *testing.T) {
	for _, tt := range wigs {
		w := make([]int, len(tt.wig))
		copy(w, tt.wig)
		PermPos(w)
		sum := 0
		for _, v := range w {
			sum += v
		}
		if sum != tt.sum {
			t.Errorf("%s: wrong sum: expected %d, actual %d", tt.name, tt.sum, sum)
		}
	}
}

var keepByteTests = []struct {
	name      string
	wig, rwig []int
	byt       []byte
	sum       int
	fun       func(int) bool
	from, to  int
}{
	{
		name: "simple",
		wig:  []int{10, 5, 0, 0, 0, 0},
		byt:  []byte("ACCCGT"),
		sum:  15,
		fun:  func(i int) bool { return true },
		rwig: []int{9, 1, 2, 3, 0, 0},
	},
	{
		name: "restricted (one A)",
		wig:  []int{10, 5, 0, 0, 0, 0},
		byt:  []byte("ACCCGT"),
		sum:  15,
		fun: func(i int) bool {
			if i%3 == 0 {
				return true
			}
			return false
		},
		rwig: []int{10, 5, 0, 0, 0, 0},
	},
	{
		name: "restricted (two As)",
		wig:  []int{10, 5, 0, 0, 0, 0},
		byt:  []byte("ACCAGT"),
		sum:  15,
		fun: func(i int) bool {
			if i%3 == 0 {
				return true
			}
			return false
		},
		rwig: []int{5, 5, 0, 5, 0, 0},
	},
	{
		name: "with range",
		wig:  []int{10, 5, 0, 0, 0, 0},
		byt:  []byte("ACCCGT"),
		sum:  15,
		fun:  func(i int) bool { return true },
		rwig: []int{0, 15, 0, 0, 0, 0},
		from: -1,
		to:   1,
	},
}

func TestPermKeepByte(t *testing.T) {
	for _, tt := range keepByteTests {
		rand.Seed(1)
		w := make([]int, len(tt.wig))
		copy(w, tt.wig)
		PermKeepByte(w, tt.byt, tt.from, tt.to, tt.fun)
		sum := 0
		for _, v := range w {
			sum += v
		}
		if sum != tt.sum {
			t.Errorf("%s: wrong sum: expected %d, actual %d", tt.name, tt.sum, sum)
		}
		if !reflect.DeepEqual(w, tt.rwig) {
			t.Errorf("%s: wrong random: expected %v, actual %v", tt.name, tt.rwig, w)
		}
		// log.Println(content(w, tt.byt))
		// log.Println(content(tt.wig, tt.byt))
	}
}

func content(wig []int, s []byte) map[string]int {
	content := make(map[string]int)
	for i, v := range wig {
		content[string(s[i])] += v
	}
	return content
}

func ExamplePerm() {
	m := []int{10, 1, 2, 3}
	Perm(m)
	fmt.Println(m)
}

func ExamplePermPos() {
	m := []int{10, 1, 2, 3}
	PermPos(m)
	for _, v := range m {
		fmt.Println(v)
	}
	// Unordered output:
	// 2
	// 1
	// 3
	// 10
}

func ExamplePermKeepByte() {
	m := []int{10, 5, 0, 0, 0, 0}
	b := []byte("ACCCGT")
	PermKeepByte(m, b, 0, 0, func(i int) bool { return true })
	fmt.Println(m)
}
