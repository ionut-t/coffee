package styles

import (
	"charm.land/lipgloss/v2"
	editor "github.com/ionut-t/goeditor/adapter-bubbletea"
)

func EditorTheme(styles Styles) editor.Theme {
	normalModeBG := styles.Info.GetForeground()
	insertModeBG := styles.Primary.GetForeground()
	visualModeBG := styles.Accent.GetForeground()
	commandModeBG := styles.Warning.GetForeground()
	searchModeBG := styles.Subtext0.GetForeground()

	return editor.Theme{
		NormalModeStyle: lipgloss.NewStyle().
			Foreground(styles.Base.GetForeground()).
			Background(normalModeBG).
			Bold(true),

		InsertModeStyle: lipgloss.NewStyle().
			Foreground(styles.Base.GetForeground()).
			Background(insertModeBG).
			Bold(true),

		VisualModeStyle: lipgloss.NewStyle().
			Foreground(styles.Base.GetForeground()).
			Background(visualModeBG).
			Bold(true),

		CommandModeStyle: lipgloss.NewStyle().
			Foreground(styles.Base.GetForeground()).
			Background(commandModeBG).
			Bold(true),

		SearchModeStyle: lipgloss.NewStyle().
			Foreground(styles.Base.GetForeground()).
			Background(searchModeBG).
			Bold(true),

		CommandLineStyle: styles.Surface0,

		StatusLineStyle: styles.Surface1.
			Foreground(styles.Subtext0.GetForeground()),

		MessageStyle: styles.Info,
		ErrorStyle:   styles.Error.Bold(true),

		LineNumberStyle: styles.Overlay0.
			Width(4).
			Align(lipgloss.Right),

		CurrentLineNumberStyle: styles.Text.
			Width(4).
			Align(lipgloss.Right),

		SelectionStyle: styles.Surface1,

		HighlightYankStyle: styles.Highlight.
			Bold(true),

		PlaceholderStyle: styles.Overlay0,

		SearchHighlightStyle: styles.Highlight.
			Bold(true),

		SearchInputTextStyle: styles.Text,

		SearchInputPromptStyle: styles.Subtext1,

		SearchInputCursorStyle: styles.Subtext1,

		CurrentLineStyle: styles.Surface0,
	}
}

func EditorLanguageTheme(isDark bool) string {
	if IsDark() {
		return "catppuccin-mocha"
	}

	return "catppuccin-latte"
}
