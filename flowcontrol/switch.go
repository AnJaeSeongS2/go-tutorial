package flowcontrol

import (
	"fmt"
	"runtime"
	"time"
)

func TourSwitch() {
	fmt.Printf("%s.\n", getCurrentOS())
	fmt.Println(GetGoodHourMessage())
}

func getCurrentOS() string {
	result := "NONE"
	switch os := runtime.GOOS; os {
	case "darwin":
		result = "OS X"
		// auto break
	case "linux":
		result = "Linux"
		// auto break
	default:
		result = os
	}
	return result
}

func GetGoodHourMessage() string {
	return getGoodHourMessage(time.Now())
}

func getGoodHourMessage(t time.Time) string {
	// switch without expression is true.
	switch {
	case t.Hour() < 7:
		return "Good night...zzzZZZ"
	case t.Hour() < 12:
		return "Good morning!"
	case t.Hour() < 17:
		return "Good afternoon."
	default:
		return "Good evening..."
	}
}
