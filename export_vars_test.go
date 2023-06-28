//go:build linux && !android
// +build linux,!android

package openssl

// Export version variables for tests
var (
	TestMajor = &vMajor
	TestMinor = &vMinor
	TestPatch = &vPatch
)
