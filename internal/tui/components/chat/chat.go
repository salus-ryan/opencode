package chat

import (
	"fmt"
	"sort"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"github.com/opencode-ai/opencode/internal/config"
	"github.com/opencode-ai/opencode/internal/message"
	"github.com/opencode-ai/opencode/internal/session"
	"github.com/opencode-ai/opencode/internal/tui/styles"
	"github.com/opencode-ai/opencode/internal/tui/theme"
	"github.com/opencode-ai/opencode/internal/version"
)

type SendMsg struct {
	Text        string
	Attachments []message.Attachment
}

type SessionSelectedMsg = session.Session

type SessionClearedMsg struct{}

type EditorFocusMsg bool

func header(width int) string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		splash(width),
		repo(width),
		"",
		cwd(width),
	)
}

// fableArt is the block-letter splash logo shown on the empty session screen.
var fableArt = []string{
	"███████ █████╗ ██████╗ ██╗     ███████",
	"██╔═════██╔══██╗██╔══██╗██║     ██╔════",
	"█████╗  ███████║██████╔╝██║     █████╗",
	"██╔══╝  ██╔══██║██╔══██╗██║     ██╔══╝",
	"██║     ██║  ██║██████╔╝███████╗███████",
	"╚═╝     ╚═╝  ╚═╝╚═════╝ ╚══════╝╚══════",
}

// splashGradient is the aurora gradient applied line-by-line to the splash.
var splashGradient = []string{
	"#c4a747", // aurora gold
	"#c39a5e",
	"#b98a8f",
	"#a87fb8",
	"#9d7bd8", // amethyst
	"#7d6fd0",
}

func splash(width int) string {
	artWidth := lipgloss.Width(fableArt[0])
	if width < artWidth+2 {
		return logo(width)
	}

	t := theme.CurrentTheme()
	baseStyle := styles.BaseStyle()

	lines := make([]string, 0, len(fableArt)+2)
	for i, line := range fableArt {
		lines = append(lines, baseStyle.
			Foreground(lipgloss.Color(splashGradient[i%len(splashGradient)])).
			Render(line))
	}

	tagline := baseStyle.
		Foreground(t.TextMuted()).
		Italic(true).
		Render("every commit tells a story — " + version.Version)
	lines = append(lines, "", tagline)

	return baseStyle.Width(width).Render(
		lipgloss.JoinVertical(lipgloss.Left, lines...),
	)
}

func lspsConfigured(width int) string {
	cfg := config.Get()
	title := "LSP Configuration"
	title = ansi.Truncate(title, width, "…")

	t := theme.CurrentTheme()
	baseStyle := styles.BaseStyle()

	lsps := baseStyle.
		Width(width).
		Foreground(t.Primary()).
		Bold(true).
		Render(title)

	// Get LSP names and sort them for consistent ordering
	var lspNames []string
	for name := range cfg.LSP {
		lspNames = append(lspNames, name)
	}
	sort.Strings(lspNames)

	var lspViews []string
	for _, name := range lspNames {
		lsp := cfg.LSP[name]
		lspName := baseStyle.
			Foreground(t.Text()).
			Render(fmt.Sprintf("• %s", name))

		cmd := lsp.Command
		cmd = ansi.Truncate(cmd, width-lipgloss.Width(lspName)-3, "…")

		lspPath := baseStyle.
			Foreground(t.TextMuted()).
			Render(fmt.Sprintf(" (%s)", cmd))

		lspViews = append(lspViews,
			baseStyle.
				Width(width).
				Render(
					lipgloss.JoinHorizontal(
						lipgloss.Left,
						lspName,
						lspPath,
					),
				),
		)
	}

	return baseStyle.
		Width(width).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				lsps,
				lipgloss.JoinVertical(
					lipgloss.Left,
					lspViews...,
				),
			),
		)
}

func logo(width int) string {
	logo := fmt.Sprintf("%s %s", styles.OpenCodeIcon, "Fable")
	t := theme.CurrentTheme()
	baseStyle := styles.BaseStyle()

	versionText := baseStyle.
		Foreground(t.TextMuted()).
		Render(version.Version)

	return baseStyle.
		Bold(true).
		Width(width).
		Render(
			lipgloss.JoinHorizontal(
				lipgloss.Left,
				logo,
				" ",
				versionText,
			),
		)
}

func repo(width int) string {
	repo := "https://github.com/salus-ryan/opencode"
	t := theme.CurrentTheme()

	return styles.BaseStyle().
		Foreground(t.TextMuted()).
		Width(width).
		Render(repo)
}

func cwd(width int) string {
	cwd := fmt.Sprintf("cwd: %s", config.WorkingDirectory())
	t := theme.CurrentTheme()

	return styles.BaseStyle().
		Foreground(t.TextMuted()).
		Width(width).
		Render(cwd)
}
