package ui

import (
	"github.com/charmbracelet/lipgloss"
)

// Theme defines all styles and colors used in the app.
type Theme struct {
	PrimaryColor    lipgloss.Color
	SecondaryColor  lipgloss.Color
	AccentColor     lipgloss.Color
	BackgroundColor lipgloss.Color
	TextColor       lipgloss.Color
	GrayColor       lipgloss.Color

	// Core styles
	TitleStyle    lipgloss.Style
	ContentStyle  lipgloss.Style
	CursorStyle   lipgloss.Style
	PromptStyle   lipgloss.Style
	ActiveStyle   lipgloss.Style
	InactiveStyle lipgloss.Style
	BorderStyle   lipgloss.Style
	StatusBar     lipgloss.Style
	Divider       lipgloss.Style

	// Message styles
	MsgStyles MsgStyles
}

// Initialize the theme
func NewTheme() Theme {
	primary := lipgloss.Color("#00E5A0")   // Mint green
	secondary := lipgloss.Color("#FFCA3A") // Yellow highlight
	accent := lipgloss.Color("#FF5F5F")    // Coral red
	background := lipgloss.Color("#101010") // Matte black
	text := lipgloss.Color("#EAEAEA")       // Soft white
	gray := lipgloss.Color("#6C757D")       // Muted gray

	return Theme{
		PrimaryColor:    primary,
		SecondaryColor:  secondary,
		AccentColor:     accent,
		BackgroundColor: background,
		TextColor:       text,
		GrayColor:       gray,

		// Core UI sections
		TitleStyle: lipgloss.NewStyle().
			Foreground(primary).
			// Background(lipgloss.Color("#181818")).
			Bold(true).
			Underline(true).
			Padding(1, 3).
			Margin(0, 0, 1, 0),

		ContentStyle: lipgloss.NewStyle().
			Foreground(text).
			// Background(background).
			Padding(1, 2),

		CursorStyle: lipgloss.NewStyle().
			Foreground(accent).
			Bold(true),

		PromptStyle: lipgloss.NewStyle().
			Foreground(secondary).
			Bold(true).
			Padding(0, 1),

		ActiveStyle: lipgloss.NewStyle().
			Foreground(primary).
			Bold(true),

		InactiveStyle: lipgloss.NewStyle().
			Foreground(gray).
			Faint(true),

		BorderStyle: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(primary).
			Padding(1, 2).
			Margin(1, 0, 0, 0),

		// UX polish: divider + status bar
		Divider: lipgloss.NewStyle().
			Foreground(gray).
			SetString("─────────────────────────────"),

		StatusBar: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#101010")).
			Background(primary).
			Bold(true).
			Padding(0, 2),

		// Message boxes (below)
		MsgStyles: DefaultMsgStyles(),
	}
}



// MsgStyles holds styles for different message types.
type MsgStyles struct {
	ErrorBox   lipgloss.Style
	InfoBox    lipgloss.Style
	SuccessBox lipgloss.Style
	WarningBox lipgloss.Style
}

// High-contrast, vivid message boxes
func DefaultMsgStyles() MsgStyles {
	return MsgStyles{
		ErrorBox: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FF4C4C")).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#FF4C4C")).
			Padding(0,5).
			Margin(0, 3).
			Bold(true),

		InfoBox: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#2D9CDB")).

			Foreground(lipgloss.Color("#EAF6FF")).
			Background(lipgloss.Color("#2D9CDB")).
			Padding(0,5).
			Margin(0, 3).
			Bold(true),

		SuccessBox: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#00E676")).
			Foreground(lipgloss.Color("#EAFBEA")).
			Background(lipgloss.Color("#00E676")).
			Padding(0,5).
			Margin(0, 3).
			Bold(true),

		WarningBox: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FFD43B")).
			Foreground(lipgloss.Color("#000000")).
			Background(lipgloss.Color("#FFD43B")).
			Padding(0,5).
			Margin(0, 3).
			Bold(true),
	}
}

// Message rendering based on type
func (t Theme) RenderMsgBox(msgType MsgType, text string) string {
	switch msgType {
	case MessageError:
		return t.MsgStyles.ErrorBox.Render("❌  " + text)
	case MessageInfo:
		return t.MsgStyles.InfoBox.Render("💡  " + text)
	case MessageSuccess:
		return t.MsgStyles.SuccessBox.Render("✅  " + text)
	case MessageWarning:
		return t.MsgStyles.WarningBox.Render("⚠️  " + text)
	default:
		return ""
	}
}

// 