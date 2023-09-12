package test

import (
	"fmt"
	"testing"

	"github.com/chubin518/kestrel-layout-advanced/pkg/config"
)

func TestConfig(t *testing.T) {
	a := config.Default().Get("server.port", 8080)

	fmt.Printf("a: %v\n", a)
	// config.Default().Set("server.port", 8080)
}
