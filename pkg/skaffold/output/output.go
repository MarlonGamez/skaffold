package output

import (
	"fmt"
	"io"

	colors "github.com/heroku/color"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/constants"
	eventV2 "github.com/GoogleContainerTools/skaffold/pkg/skaffold/event/v2"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/util"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/color"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

func SetupOutput(out io.Writer, defaultColor int, forceColors bool) io.Writer {
	_, isTerm := util.IsTerminal(out)
	supportsColor, err := util.SupportsColor()
	if err != nil {
		logrus.Debugf("error checking for color support: %v", err)
	}

	useColors := (isTerm && supportsColor) || forceColors
	if useColors {
		// Use EnableColorsStdout to enable use of color on Windows
		useColors = false // value is updated if color-enablement is successful
		colorable.EnableColorsStdout(&useColors)
	}
	colors.Disable(!useColors)

	// Maintain compatibility with the old color coding.
	color.Default = map[int]color.Color{
		91: color.LightRed,
		92: color.LightGreen,
		93: color.LightYellow,
		94: color.LightBlue,
		95: color.LightPurple,
		31: color.Red,
		32: color.Green,
		33: color.Yellow,
		34: color.Blue,
		35: color.Purple,
		36: color.Cyan,
		37: color.White,
		0:  color.None,
	}[defaultColor]

	eventOut := eventV2.NewLogger(constants.DevLoop, "0")
	multiOut := io.MultiWriter(eventOut, out)

	if useColors {
		fmt.Println("using colors")
		return color.NewWriter(multiOut)
	}

	fmt.Println("not using colors")
	return multiOut
}
