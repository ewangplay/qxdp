package adapter_test

import (
	"testing"

	"github.com/ewangplay/qxdp/adapter"
)

func TestAdapter(t *testing.T) {
	var target adapter.Target 

	target = adapter.NewAdapter("Tom")
	target.Greet()
}
