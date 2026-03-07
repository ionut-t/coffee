package styles

import (
	"charm.land/bubbles/v2/list"
	"charm.land/lipgloss/v2"
)

func ListStyles() (s list.Styles) {
	const (
		bullet   = "•"
		ellipsis = "…"
	)

	verySubduedColor := AdaptiveColorFromString("#DDDADA", "#3C3C3C")
	subduedColor := AdaptiveColorFromString("#9B9B9B", "#5C5C5C")

	s.TitleBar = lipgloss.NewStyle().Padding(0, 0, 1, 2)

	s.Title = Primary.Bold(true)

	s.Spinner = lipgloss.NewStyle().
		Foreground(AdaptiveColorFromString("#8E8E8E", "#747373"))

	s.Filter.Focused.Prompt = Primary

	s.Filter.Cursor.Color = Primary.GetForeground()

	s.DefaultFilterCharacterMatch = lipgloss.NewStyle().Underline(true)

	s.StatusBar = lipgloss.NewStyle().
		Foreground(AdaptiveColorFromString("#A49FA5", "#777777")).
		Padding(0, 0, 1, 2)

	s.StatusEmpty = lipgloss.NewStyle().Foreground(subduedColor)

	s.StatusBarActiveFilter = Primary

	s.StatusBarFilterCount = lipgloss.NewStyle().Foreground(verySubduedColor)

	s.NoItems = Base

	s.ArabicPagination = lipgloss.NewStyle().Foreground(subduedColor)

	s.PaginationStyle = lipgloss.NewStyle().PaddingLeft(2)

	s.HelpStyle = lipgloss.NewStyle().Padding(1, 0, 0, 2)

	s.ActivePaginationDot = lipgloss.NewStyle().
		Foreground(AdaptiveColorFromString("#847A85", "#979797")).
		SetString(bullet)

	s.InactivePaginationDot = lipgloss.NewStyle().
		Foreground(verySubduedColor).
		SetString(bullet)

	s.DividerDot = lipgloss.NewStyle().
		Foreground(verySubduedColor).
		SetString(" " + bullet + " ")

	return s
}

func ListItemStyles() (s list.DefaultItemStyles) {
	s.NormalTitle = Text.Padding(0, 0, 0, 2)

	s.NormalDesc = s.NormalTitle.Foreground(Overlay0.GetForeground())

	s.SelectedTitle = Primary.
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(Primary.GetForeground()).
		Padding(0, 0, 0, 1)

	s.SelectedDesc = s.SelectedTitle.
		Foreground(Accent.GetForeground())

	s.DimmedTitle = lipgloss.NewStyle().
		Foreground(AdaptiveColorFromString("#A49FA5", "#777777")).
		Padding(0, 0, 0, 2)

	s.DimmedDesc = s.DimmedTitle.
		Foreground(AdaptiveColorFromString("#C2B8C2", "#4D4D4D"))

	s.FilterMatch = lipgloss.NewStyle().Underline(true)

	return s
}
