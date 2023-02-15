package engine

import (
	"testing"
)

func TestNewEngine(t *testing.T) {
	app := NewEngine(true)
	app.Run(":8881")
}
