package ui

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	PrimaryColor    lipgloss.Color
	SecondaryColor  lipgloss.Color
	AccentColor     lipgloss.Color
	BackgroundColor lipgloss.Color
	TextColor       lipgloss.Color
	GrayColor       lipgloss.Color

	// derived styles
	TitleStyle    lipgloss.Style
	ContentStyle  lipgloss.Style
	CursorStyle   lipgloss.Style
	PromptStyle   lipgloss.Style
	ActiveStyle   lipgloss.Style
	InactiveStyle lipgloss.Style
	BorderStyle   lipgloss.Style
	ErrorStyle    lipgloss.Style
}

// NewTheme initializes a default theme
func NewTheme() Theme {
	primary := lipgloss.Color("#00C896")   // greenish
	secondary := lipgloss.Color("#FFDD57") // yellow
	accent := lipgloss.Color("#FF6F61")    // reddish
	background := lipgloss.Color("#1E1E1E") // dark gray
	text := lipgloss.Color("#FFFFFF")       // white
	gray := lipgloss.Color("#808080")       // gray

	return Theme{
		PrimaryColor:    primary,
		SecondaryColor:  secondary,
		AccentColor:     accent,
		BackgroundColor: background,
		TextColor:       text,
		GrayColor:       gray,

		TitleStyle:   lipgloss.NewStyle().Foreground(primary).Background(background).Bold(true).Padding(0, 1).Margin(0, 1).Underline(true),
		ContentStyle: lipgloss.NewStyle().Foreground(text).Background(background).Padding(1, 2),
		CursorStyle:  lipgloss.NewStyle().Foreground(accent).Bold(true),
		PromptStyle:  lipgloss.NewStyle().Foreground(secondary).Bold(true),
		ActiveStyle:  lipgloss.NewStyle().Foreground(primary).Bold(true),
		// inactive: faint text on background
		InactiveStyle: lipgloss.NewStyle().Foreground(gray).Faint(true).Padding(0, 0).Margin(0, 0),
		BorderStyle:   lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(primary).Padding(1, 2),
		ErrorStyle:    lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Bold(true),
	}
}
