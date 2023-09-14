package xtypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEnum(t *testing.T) {
	color := NewEnum()
	color.NewMember("RED", 1, "红色")
	color.NewMember("YELLOW", 2, "黄色")
	assert.Equal(t, 1, color.Number("RED"))
	assert.Equal(t, "黄色", color.Desc("YELLOW"))
}
