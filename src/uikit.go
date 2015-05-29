package uikit

import (
	"github.com/google/gxui"
	"github.com/google/gxui/themes/dark"
)

const (
	windowName string = "test window"
)

func AppMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)
	window := theme.CreateWindow(1920, 1080, windowName)

	window.OnClose(driver.Terminate)
}
