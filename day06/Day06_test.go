package day06

import (
	"testing"

	"github.com/kptlr/aoc2023/utils"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input := utils.ReadLines("test", "day06")
	result := PartOne(input)
	assert.Equal(t, 288, result)
}

func TestPartTwo(t *testing.T) {
	input := utils.ReadLines("test", "day06")
	result := PartTwo(input)
	assert.Equal(t, 71503, result)
}
