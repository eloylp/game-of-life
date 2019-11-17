package main

import (
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

func TestBoard_CheckNeighbour(t *testing.T) {
	b := Board{
		{true, true, true, true},
		{true, true, true, true},
		{true, true, true, true},
		{true, true, true, true},
	}
	type sample struct {
		X, Y, OX, OY int
		Expected     Cell
	}

	samples := func() []sample {
		return []sample{
			{0, 0, 0, 0, true},
			{0, 0, 3, 3, true},
			{0, 0, -1, 0, false},
			{0, 0, 0, -1, false},
		}
	}

	for _, s := range samples() {
		result := CheckNeighbour(b, s.X, s.Y, s.OX, s.OY)
		assert.Equal(t, s.Expected, result)
	}
}
