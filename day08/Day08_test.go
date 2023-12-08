package day08

import (
	"testing"

	"github.com/kptlr/aoc2023/utils"
	"github.com/stretchr/testify/assert"
)

func TestPartOne_1(t *testing.T) {
	input := utils.ReadLines("test", "day08_p1")
	result := PartOne(input)
	assert.Equal(t, 2, result)
}

func TestPartOne_2(t *testing.T) {
	input := utils.ReadLines("test", "day08_p2")
	result := PartOne(input)
	assert.Equal(t, 6, result)
}

func TestPartTwo(t *testing.T) {
	input := utils.ReadLines("test", "day08_p1")
	result := PartTwo(input)
	assert.Equal(t, 1, result)
}
