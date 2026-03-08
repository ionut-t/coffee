package styles

import (
	"image/color"

	"charm.land/fang/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/exp/charmtone"
)

func FangColorScheme(c lipgloss.LightDarkFunc) fang.ColorScheme {
	styles := New(IsDark())

	return fang.ColorScheme{
		Base:           styles.Primary.GetForeground(),
		Title:          styles.Accent.GetForeground(),
		Codeblock:      styles.Surface0.GetBackground(),
		Program:        styles.Primary.GetForeground(),
		Command:        styles.Primary.GetForeground(),
		DimmedArgument: styles.Overlay1.GetForeground(),
		Comment:        styles.Subtext1.GetForeground(),
		Flag:           styles.Success.GetForeground(),
		Argument:       styles.Text.GetForeground(),
		Description:    styles.Text.GetForeground(),
		FlagDefault:    styles.Text.GetForeground(),
		QuotedString:   styles.Accent.GetForeground(),
		ErrorHeader: [2]color.Color{
			charmtone.Butter,
			charmtone.Cherry,
		},
	}
}
