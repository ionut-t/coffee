package styles

import (
	"image/color"
	"os"

	"charm.land/lipgloss/v2"
	catppuccin "github.com/catppuccin/go"
)

var isDark bool

func init() {
	isDark = lipgloss.HasDarkBackground(os.Stdout, os.Stderr)
}

func IsDark() bool {
	return isDark
}

func AdaptiveColor(light, dark color.Color) color.Color {
	if isDark {
		return dark
	}

	return light
}

func AdaptiveColorFromString(light, dark string) color.Color {
	if isDark {
		return lipgloss.Color(dark)
	}

	return lipgloss.Color(light)
}

var (
	Base = lipgloss.NewStyle().
		Foreground(AdaptiveColor(catppuccin.Latte.Base(), catppuccin.Mocha.Base()))

	Text = lipgloss.NewStyle().
		Foreground(AdaptiveColor(catppuccin.Latte.Text(), catppuccin.Mocha.Text()))

	Primary = lipgloss.NewStyle().
		Foreground(AdaptiveColor(catppuccin.Latte.Sapphire(), catppuccin.Mocha.Sapphire()))

	Accent = lipgloss.NewStyle().
		Foreground(AdaptiveColor(catppuccin.Latte.Teal(), catppuccin.Mocha.Teal()))

	Success = lipgloss.NewStyle().
		Foreground(AdaptiveColor(catppuccin.Latte.Green(), catppuccin.Mocha.Green()))

	Error = lipgloss.NewStyle().
		Foreground(AdaptiveColor(catppuccin.Latte.Red(), catppuccin.Mocha.Red()))

	Warning = lipgloss.NewStyle().
		Foreground(AdaptiveColor(catppuccin.Latte.Yellow(), catppuccin.Mocha.Yellow()))

	Info = lipgloss.NewStyle().
		Foreground(AdaptiveColor(catppuccin.Latte.Blue(), catppuccin.Mocha.Blue()))

	Subtext0 = lipgloss.NewStyle().
			Foreground(AdaptiveColor(catppuccin.Latte.Subtext0(), catppuccin.Mocha.Subtext0()))

	Subtext1 = lipgloss.NewStyle().
			Foreground(AdaptiveColor(catppuccin.Latte.Subtext1(), catppuccin.Mocha.Subtext1()))

	Overlay0 = lipgloss.NewStyle().
			Foreground(AdaptiveColor(catppuccin.Latte.Overlay0(), catppuccin.Mocha.Overlay0()))

	Overlay1 = lipgloss.NewStyle().
			Foreground(AdaptiveColor(catppuccin.Latte.Overlay1(), catppuccin.Mocha.Overlay1()))

	Surface0 = lipgloss.NewStyle().
			Background(AdaptiveColor(catppuccin.Latte.Surface0(), catppuccin.Mocha.Surface0()))

	Surface1 = lipgloss.NewStyle().
			Background(AdaptiveColor(catppuccin.Latte.Surface1(), catppuccin.Mocha.Surface1()))

	Crust = lipgloss.NewStyle().
		Background(AdaptiveColor(catppuccin.Latte.Crust(), catppuccin.Mocha.Crust()))

	AccentBackground = lipgloss.NewStyle().
				Background(AdaptiveColor(catppuccin.Latte.Teal(), catppuccin.Mocha.Teal()))

	Highlight = lipgloss.NewStyle().
			Foreground(Base.GetForeground()).
			Background(Primary.GetForeground())
)
