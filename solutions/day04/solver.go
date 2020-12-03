package day04

import (
	"crypto/md5"
	"fmt"
)

type Solver struct {
	input string
}

func New(input string) *Solver {
	return &Solver{input: input}
}

func findHash(input, prefix string) int64 {
	var i int64 = 1
	for {
		key := []byte(fmt.Sprintf("%s%d", input, i))
		hash := fmt.Sprintf("%x", md5.Sum(key))
		if hash[:len(prefix)] == prefix {
			return i
		}
		i++
	}
}

func (s *Solver) SolveA() string {
	i := findHash(s.input, "00000")
	return fmt.Sprint(i)
}

func (s *Solver) SolveB() string {
	i := findHash(s.input, "000000")
	return fmt.Sprint(i)
}
