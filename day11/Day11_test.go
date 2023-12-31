package day11

import (
	"testing"

	"github.com/kptlr/aoc2023/utils"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input := utils.ReadLines("test", "day11")
	result := PartOne(input)
	assert.Equal(t, 0, result)
}

func TestPartTwo(t *testing.T) {
	input := utils.ReadLines("test", "day11")
	result := PartTwo(input)
	assert.Equal(t, 1, result)
}
