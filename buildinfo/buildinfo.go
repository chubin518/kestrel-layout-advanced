package buildinfo

import (
	"os"
	"path"
)

type Env string

const (
	DEV   Env = "dev"
	TEST  Env = "test"
	STAGE Env = "stage"
	PROD  Env = "prod"
)

var (
	// build version
	Version string = "0.0.1-beta1"
	// build name
	Name string = path.Base(os.Args[0])
	// build env: dev/test/stage/prod
	Environment string = "dev"
	// build time
	BuildTime string
	// build go version
	GoVersion string
)

// Active
func Active() Env {
	return Env(Environment)
}

// IsDev
func IsDev() bool {
	return Active() == DEV
}

// IsTest
func IsTest() bool {
	return Active() == TEST
}

// IsStage
func IsStage() bool {
	return Active() == STAGE
}

// IsProd
func IsProd() bool {
	return Active() == PROD
}
