package biscuit

import "runtime"

type platform string

const (
	windows platform = "Windows"
	mac platform = "Mac"
	linux platform = "Linux"
	unknown platform = "Unknown"
)

func detectOS() platform {
	switch os := runtime.GOOS; os {
	case "darwin":
		return mac
	case "linux":
		return linux
	case "windows":
		return windows
	default:
		return unknown
	}
}