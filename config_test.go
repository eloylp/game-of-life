package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckConfig(t *testing.T) {

	type sample struct {
		Config Cfg
		Error  error
	}

	samples := []sample{
		{Cfg{-1, 10, 10, 1},
			errors.New("Cannot set X axis under 1")},
		{Cfg{1, -1, 10, 1},
			errors.New("Cannot set Y axis under 1")},
		{Cfg{1, 1, 0, 1},
			errors.New("Cannot set generations  under 1")},
		{Cfg{1, 1, 1, 0},
			errors.New("Cannot set intervals  under 1")},
		{Cfg{1, 1, 1, 11},
			nil},
	}

	for _, s := range samples {
		assert.Equal(t, s.Error, CheckConfig(s.Config))
	}
}
