package styles

import (
	"charm.land/huh/v2"
	"charm.land/lipgloss/v2"
)

type HuhThemeCatppuccin struct {
	Styles Styles
}

func (t HuhThemeCatppuccin) Theme(isDark bool) *huh.Styles {
	theme := huh.ThemeBase(isDark)
	styles := t.Styles

	var (
		base     = styles.Base.GetForeground()
		text     = styles.Text.GetForeground()
		subtext1 = styles.Subtext1.GetForeground()
		subtext0 = styles.Subtext0.GetForeground()
		overlay1 = styles.Overlay1.GetForeground()
		overlay0 = styles.Overlay0.GetForeground()
		green    = styles.Success.GetForeground()
		red      = styles.Error.GetForeground()
		accent   = styles.Accent.GetForeground()
		primary  = styles.Primary.GetForeground()
		cursor   = styles.Primary.GetForeground()
		surface0 = styles.Surface0.GetBackground()
	)

	theme.Focused.Base = theme.Focused.Base.BorderForeground(subtext1)
	theme.Focused.Title = theme.Focused.Title.Foreground(primary)
	theme.Focused.NoteTitle = theme.Focused.NoteTitle.Foreground(primary)
	theme.Focused.Directory = theme.Focused.Directory.Foreground(primary)
	theme.Focused.Description = theme.Focused.Description.Foreground(subtext0)
	theme.Focused.ErrorIndicator = theme.Focused.ErrorIndicator.Foreground(red)
	theme.Focused.ErrorMessage = theme.Focused.ErrorMessage.Foreground(red)
	theme.Focused.SelectSelector = theme.Focused.SelectSelector.Foreground(accent)
	theme.Focused.NextIndicator = theme.Focused.NextIndicator.Foreground(accent)
	theme.Focused.PrevIndicator = theme.Focused.PrevIndicator.Foreground(accent)
	theme.Focused.Option = theme.Focused.Option.Foreground(text)
	theme.Focused.MultiSelectSelector = theme.Focused.MultiSelectSelector.Foreground(accent)
	theme.Focused.SelectedOption = theme.Focused.SelectedOption.Foreground(green)
	theme.Focused.SelectedPrefix = theme.Focused.SelectedPrefix.Foreground(green)
	theme.Focused.UnselectedPrefix = theme.Focused.UnselectedPrefix.Foreground(text)
	theme.Focused.UnselectedOption = theme.Focused.UnselectedOption.Foreground(text)
	theme.Focused.FocusedButton = theme.Focused.FocusedButton.Foreground(base).Background(primary)
	theme.Focused.BlurredButton = theme.Focused.BlurredButton.Foreground(text).Background(surface0)

	theme.Focused.TextInput.Cursor = theme.Focused.TextInput.Cursor.Foreground(cursor)
	theme.Focused.TextInput.Placeholder = theme.Focused.TextInput.Placeholder.Foreground(overlay0)
	theme.Focused.TextInput.Prompt = theme.Focused.TextInput.Prompt.Foreground(accent)

	theme.Blurred = theme.Focused
	theme.Blurred.Base = theme.Blurred.Base.BorderStyle(lipgloss.HiddenBorder())

	theme.Help.Ellipsis = theme.Help.Ellipsis.Foreground(subtext0)
	theme.Help.ShortKey = theme.Help.ShortKey.Foreground(subtext0)
	theme.Help.ShortDesc = theme.Help.ShortDesc.Foreground(overlay1)
	theme.Help.ShortSeparator = theme.Help.ShortSeparator.Foreground(subtext0)
	theme.Help.FullKey = theme.Help.FullKey.Foreground(subtext0)
	theme.Help.FullDesc = theme.Help.FullDesc.Foreground(overlay1)
	theme.Help.FullSeparator = theme.Help.FullSeparator.Foreground(subtext0)

	return theme
}
