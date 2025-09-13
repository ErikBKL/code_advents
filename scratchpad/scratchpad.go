package scratchpad

import (
)

func CreateMap() map[int][]int {
	m := map[int][]int{}

	m[12] = append(m[12], 1)
	m[12] = append(m[12], 2)
	m[12] = append(m[12], 3)
	m[11] = append(m[11], 4)
	m[10] = append(m[10], 5)
	m[17] = append(m[17], 6)

	return m
}