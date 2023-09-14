package xtypes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	s := NewSet()
	s.Add("hello")
	s.Add("world")
	s.Add("hello")
	assert.Equal(t, 2, len(s))
}

func TestNewSetWithSlice(t *testing.T) {
	data := []interface{}{1, 3, 2, 4, 2, 3}
	s := NewSetWithSlice(data)
	assert.Equal(t, 4, len(s))
	fmt.Println("set data", s.All())
}
