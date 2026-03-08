package styles

import (
	"image/color"
	"os"

	"charm.land/lipgloss/v2"
	catppuccin "github.com/catppuccin/go"
)

func IsDark() bool {
	return lipgloss.HasDarkBackground(os.Stdout, os.Stderr)
}

func AdaptiveColorFromString(light, dark string) color.Color {
	if IsDark() {
		return lipgloss.Color(dark)
	}

	return lipgloss.Color(light)
}

type Styles struct {
	Base             lipgloss.Style
	Text             lipgloss.Style
	Primary          lipgloss.Style
	Accent           lipgloss.Style
	Success          lipgloss.Style
	Error            lipgloss.Style
	Warning          lipgloss.Style
	Info             lipgloss.Style
	Subtext0         lipgloss.Style
	Subtext1         lipgloss.Style
	Overlay0         lipgloss.Style
	Overlay1         lipgloss.Style
	Surface0         lipgloss.Style
	Surface1         lipgloss.Style
	Crust            lipgloss.Style
	AccentBackground lipgloss.Style
	Highlight        lipgloss.Style
	ActiveBorder     lipgloss.Style
	InactiveBorder   lipgloss.Style
}

func New(isDark bool) Styles {
	lightDark := lipgloss.LightDark(isDark)

	return Styles{
		Base: lipgloss.NewStyle().Foreground(lightDark(catppuccin.Latte.Base(), catppuccin.Mocha.Base())),

		Text: lipgloss.NewStyle().Foreground(lightDark(catppuccin.Latte.Text(), catppuccin.Mocha.Text())),

		Primary: lipgloss.NewStyle().Foreground(lightDark(catppuccin.Latte.Sapphire(), catppuccin.Mocha.Sapphire())),

		Accent: lipgloss.NewStyle().Foreground(lightDark(catppuccin.Latte.Teal(), catppuccin.Mocha.Teal())),

		Success: lipgloss.NewStyle().Foreground(lightDark(catppuccin.Latte.Green(), catppuccin.Mocha.Green())),

		Error: lipgloss.NewStyle().Foreground(lightDark(catppuccin.Latte.Red(), catppuccin.Mocha.Red())),

		Warning: lipgloss.NewStyle().Foreground(lightDark(catppuccin.Latte.Yellow(), catppuccin.Mocha.Yellow())),

		Info: lipgloss.NewStyle().Foreground(lightDark(catppuccin.Latte.Blue(), catppuccin.Mocha.Blue())),

		Subtext0: lipgloss.NewStyle().Foreground(lightDark(catppuccin.Latte.Subtext0(), catppuccin.Mocha.Subtext0())),

		Subtext1: lipgloss.NewStyle().Foreground(lightDark(catppuccin.Latte.Subtext1(), catppuccin.Mocha.Subtext1())),

		Overlay0: lipgloss.NewStyle().Foreground(lightDark(catppuccin.Latte.Overlay0(), catppuccin.Mocha.Overlay0())),

		Overlay1: lipgloss.NewStyle().Foreground(lightDark(catppuccin.Latte.Overlay1(), catppuccin.Mocha.Overlay1())),

		Surface0: lipgloss.NewStyle().Background(lightDark(catppuccin.Latte.Surface0(), catppuccin.Mocha.Surface0())),

		Surface1: lipgloss.NewStyle().Background(lightDark(catppuccin.Latte.Surface1(), catppuccin.Mocha.Surface1())),

		Crust: lipgloss.NewStyle().Background(lightDark(catppuccin.Latte.Crust(), catppuccin.Mocha.Crust())),

		AccentBackground: lipgloss.NewStyle().Background(lightDark(catppuccin.Latte.Teal(), catppuccin.Mocha.Teal())),

		Highlight: lipgloss.NewStyle().Foreground(lightDark(catppuccin.Latte.Base(), catppuccin.Mocha.Base())).Background(lightDark(catppuccin.Latte.Sapphire(), catppuccin.Mocha.Sapphire())),

		ActiveBorder: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lightDark(catppuccin.Latte.Sapphire(), catppuccin.Mocha.Sapphire())),

		InactiveBorder: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lightDark(catppuccin.Latte.Overlay0(), catppuccin.Mocha.Overlay0())),
	}
}
