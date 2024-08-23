package biscuit

import "runtime"

type OS string

const (
	windows OS = "Windows"
	mac OS = "Mac"
	linux OS = "Linux"
	unknown OS = "Unknown"
)

func detectOS() OS {
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