# ☕️ Coffee

This repository is a collection of Go packages designed to help build beautiful and functional terminal user interfaces (TUIs) with [Bubble Tea](https://github.com/charmbracelet/bubbletea). The packages provide styling, UI components, and helpers, with a focus on the [Catppuccin](https://github.com/catppuccin) color palette.

## Packages

### `styles`

This is the core styling package, providing a comprehensive set of tools and themes for TUI components.

- **Colors**: A complete, adaptive Catppuccin color palette for both light and dark backgrounds.
- **Themes**: Pre-configured themes for `huh` forms, `gotable` tables, and `goeditor` text editors.
- **Component Styles**: Styled definitions for `bubbles/list` items, borders, and more.
- **Layout Helpers**: Utility functions like `Wrap` for word-wrapping text within a given width.

### `markdown`

A simple package for rendering Markdown in the terminal. It acts as a wrapper around the `glamour` library and comes pre-configured with Catppuccin themes (Mocha for dark backgrounds, Latte for light). It automatically detects the terminal's background color to apply the appropriate theme.

### `help`

A helper package for creating formatted and styled help screens. It can render help views for both key bindings (`bubbles/key`) and general command descriptions, using `lipgloss` for layout and styling.
