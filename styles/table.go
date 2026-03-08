package styles

import (
	"charm.land/lipgloss/v2"
	table "github.com/ionut-t/gotable"
)

func TableTheme(styles Styles) table.Theme {
	return table.Theme{
		Header:       styles.AccentBackground.Bold(true),
		Cell:         styles.Subtext0,
		Border:       styles.Overlay0,
		SelectedRow:  styles.Highlight.Bold(true),
		SelectedCell: lipgloss.NewStyle().Bold(true).Background(styles.Text.GetForeground()).Foreground(styles.AccentBackground.GetBackground()),
	}
}
