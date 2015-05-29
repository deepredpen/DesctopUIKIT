package uikit

import (
	"io/ioutil"
	"log"

	"github.com/google/gxui"
	"github.com/google/gxui/samples/flags"
	"github.com/google/gxui/themes/dark"
)

const (
	windowName string = "test window"
)

func getFont(name string) []byte {
	//file, err := os.Open("../font/ + name")
	file, err := ioutil.ReadFile("../font/" + name)
	if err != nil {
		log.Fatalln(err)
	}
	return file
}

func AppMain(driver gxui.Driver) {

	theme := dark.CreateTheme(driver)
	font, err := driver.CreateFont(getFont("VL-Gothic-Regular.ttf"), 75)

	window := theme.CreateWindow(1920, 1080, windowName)
	//canvas := driver.CreateCanvas(math.Size{W: 1920, H: 1080})

	if err != nil {
		log.Fatalln(err)
	}

	label := theme.CreateLabel()

	label.SetFont(font)
	label.SetText("本日は晴天なり")

	//label.SetSize(math.Size{H: 10000, W: 10000})

	window.SetScale(flags.DefaultScaleFactor)

	window.OnClose(driver.Terminate)
}
