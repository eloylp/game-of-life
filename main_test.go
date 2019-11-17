package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type sample struct {
	Matrix         Board
	ExpectedMatrix Board
}

func samples() []sample {
	return []sample{
		{
			Matrix: Board{
				{false, false, false, false},
				{false, true, false, false},
				{false, true, true, false},
				{false, true, false, false},
			},
			ExpectedMatrix: Board{
				{false, false, false, false},
				{false, true, true, false},
				{true, true, true, false},
				{false, true, true, false},
			},
		},
	}
}

func TestBoard_NextGen(t *testing.T) {
	for _, s := range samples() {

		nextGen := s.Matrix.NextGen()
		assert.Equal(t, s.ExpectedMatrix, nextGen)
	}
}

func TestIsReachableNeighbour(t *testing.T) {
	b := Board{
		{true, true, true, true},
		{true, true, true, true},
		{true, true, true, true},
		{true, true, true, true},
	}
	type sample struct {
		X, Y, OX, OY int
		Case         string
		Expected     bool
	}

	samples := func() []sample {
		return []sample{
			{0, 0, 0, 0, "top left cell without offsets is in range", true},
			{0, 0, 3, 3, "top right cell without offsets is in range", true},
			{0, 0, -1, 0, "top left cell with offset x-1 is out of range", false},
			{0, 0, 0, -1, "top left cell with offset y-1 is out of range", false},
			{0, 3, 0, 1, "bottom left cell with offset y+1 is out of range", false},
			{0, 3, -1, 0, "bottom left cell with offset x-1 is out of range", false},
			{3, 0, 1, 0, "top right cell with offset x+1 is out of range", false},
			{3, 0, 0, -1, "top right cell with offset y-1 is out of range", false},
			{3, 3, 1, 0, "bottom right cell with offset x+1 is out of range", false},
			{3, 3, 1, 1, "bottom right cell with offset y+1 is out of range", false},
		}
	}

	for _, s := range samples() {
		result := IsReachableNeighbour(b, s.X, s.Y, s.OX, s.OY)
		assert.Equal(t, s.Expected, result, fmt.Sprintf("Failed case: %s", s.Case))
	}
}
