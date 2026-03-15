package styles

import (
	"charm.land/bubbles/v2/list"
	"charm.land/lipgloss/v2"
)

func ListStyles(styles Styles, isDark bool) (s list.Styles) {
	const (
		bullet   = "•"
		ellipsis = "…"
	)

	lightDark := AdaptiveColorFromString(isDark)

	verySubduedColor := lightDark("#DDDADA", "#3C3C3C")
	subduedColor := lightDark("#9B9B9B", "#5C5C5C")

	s.TitleBar = lipgloss.NewStyle().Padding(0, 0, 1, 2)

	s.Title = styles.Primary.Bold(true)

	s.Spinner = lipgloss.NewStyle().
		Foreground(lightDark("#8E8E8E", "#747373"))

	s.Filter.Focused.Prompt = styles.Primary

	s.Filter.Cursor.Color = styles.Primary.GetForeground()

	s.DefaultFilterCharacterMatch = lipgloss.NewStyle().Underline(true)

	s.StatusBar = lipgloss.NewStyle().
		Foreground(lightDark("#A49FA5", "#777777")).
		Padding(0, 0, 1, 2)

	s.StatusEmpty = lipgloss.NewStyle().Foreground(subduedColor)

	s.StatusBarActiveFilter = styles.Primary

	s.StatusBarFilterCount = lipgloss.NewStyle().Foreground(verySubduedColor)

	s.NoItems = styles.Base

	s.ArabicPagination = lipgloss.NewStyle().Foreground(subduedColor)

	s.PaginationStyle = lipgloss.NewStyle().PaddingLeft(2)

	s.HelpStyle = lipgloss.NewStyle().Padding(1, 0, 0, 2)

	s.ActivePaginationDot = lipgloss.NewStyle().
		Foreground(lightDark("#847A85", "#979797")).
		SetString(bullet)

	s.InactivePaginationDot = lipgloss.NewStyle().
		Foreground(verySubduedColor).
		SetString(bullet)

	s.DividerDot = lipgloss.NewStyle().
		Foreground(verySubduedColor).
		SetString(" " + bullet + " ")

	return s
}

func ListItemStyles(styles Styles, isDark bool) (s list.DefaultItemStyles) {
	lightDark := AdaptiveColorFromString(isDark)

	s.NormalTitle = styles.Text.Padding(0, 0, 0, 2)

	s.NormalDesc = s.NormalTitle.Foreground(styles.Overlay0.GetForeground())

	s.SelectedTitle = styles.Primary.
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(styles.Primary.GetForeground()).
		Padding(0, 0, 0, 1)

	s.SelectedDesc = s.SelectedTitle.
		Foreground(styles.Accent.GetForeground())

	s.DimmedTitle = lipgloss.NewStyle().
		Foreground(lightDark("#A49FA5", "#777777")).
		Padding(0, 0, 0, 2)

	s.DimmedDesc = s.DimmedTitle.
		Foreground(lightDark("#C2B8C2", "#4D4D4D"))

	s.FilterMatch = lipgloss.NewStyle().Underline(true)

	return s
}
