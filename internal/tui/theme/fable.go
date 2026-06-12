package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// FableTheme is the signature theme for Fable — a deep indigo night
// palette with aurora-gold accents. Designed for long sessions: low
// background luminance, warm accents, high-contrast diff colors.
type FableTheme struct {
	BaseTheme
}

// NewFableTheme creates the Fable signature theme.
func NewFableTheme() *FableTheme {
	// Dark mode — "midnight library"
	darkBackground := "#13111c" // deep indigo-black
	darkCurrentLine := "#1c1929"
	darkSelection := "#2a2440"
	darkForeground := "#e8e3f5" // soft lavender white
	darkComment := "#6f679a"    // muted violet
	darkPrimary := "#c4a747"    // aurora gold
	darkSecondary := "#9d7bd8"  // soft amethyst
	darkAccent := "#5ad4c6"     // spectral teal
	darkRed := "#f0719b"        // rose error
	darkOrange := "#f5a97f"     // ember warning
	darkGreen := "#a6e3a1"      // moss success
	darkCyan := "#89dceb"       // sky info
	darkYellow := "#e5c890"     // candlelight emphasis
	darkBorder := "#332d4d"     // violet border

	// Light mode — "morning parchment"
	lightBackground := "#faf6ee"
	lightCurrentLine := "#f0ead9"
	lightSelection := "#e3d9c0"
	lightForeground := "#3a3247"
	lightComment := "#8a8099"
	lightPrimary := "#8a6d1f" // deep gold
	lightSecondary := "#6b4bb8"
	lightAccent := "#0f8a7d"
	lightRed := "#c33764"
	lightOrange := "#c2602a"
	lightGreen := "#3a7d44"
	lightCyan := "#1e7f9e"
	lightYellow := "#a07d1c"
	lightBorder := "#d8cfba"

	theme := &FableTheme{}

	theme.PrimaryColor = lipgloss.AdaptiveColor{Dark: darkPrimary, Light: lightPrimary}
	theme.SecondaryColor = lipgloss.AdaptiveColor{Dark: darkSecondary, Light: lightSecondary}
	theme.AccentColor = lipgloss.AdaptiveColor{Dark: darkAccent, Light: lightAccent}

	theme.ErrorColor = lipgloss.AdaptiveColor{Dark: darkRed, Light: lightRed}
	theme.WarningColor = lipgloss.AdaptiveColor{Dark: darkOrange, Light: lightOrange}
	theme.SuccessColor = lipgloss.AdaptiveColor{Dark: darkGreen, Light: lightGreen}
	theme.InfoColor = lipgloss.AdaptiveColor{Dark: darkCyan, Light: lightCyan}

	theme.TextColor = lipgloss.AdaptiveColor{Dark: darkForeground, Light: lightForeground}
	theme.TextMutedColor = lipgloss.AdaptiveColor{Dark: darkComment, Light: lightComment}
	theme.TextEmphasizedColor = lipgloss.AdaptiveColor{Dark: darkYellow, Light: lightYellow}

	theme.BackgroundColor = lipgloss.AdaptiveColor{Dark: darkBackground, Light: lightBackground}
	theme.BackgroundSecondaryColor = lipgloss.AdaptiveColor{Dark: darkCurrentLine, Light: lightCurrentLine}
	theme.BackgroundDarkerColor = lipgloss.AdaptiveColor{Dark: "#0d0b14", Light: "#fffdf7"}

	theme.BorderNormalColor = lipgloss.AdaptiveColor{Dark: darkBorder, Light: lightBorder}
	theme.BorderFocusedColor = lipgloss.AdaptiveColor{Dark: darkPrimary, Light: lightPrimary}
	theme.BorderDimColor = lipgloss.AdaptiveColor{Dark: darkSelection, Light: lightSelection}

	theme.DiffAddedColor = lipgloss.AdaptiveColor{Dark: "#a6e3a1", Light: "#3a7d44"}
	theme.DiffRemovedColor = lipgloss.AdaptiveColor{Dark: "#f0719b", Light: "#c33764"}
	theme.DiffContextColor = lipgloss.AdaptiveColor{Dark: "#8a82a8", Light: "#7a7287"}
	theme.DiffHunkHeaderColor = lipgloss.AdaptiveColor{Dark: "#8a82a8", Light: "#7a7287"}
	theme.DiffHighlightAddedColor = lipgloss.AdaptiveColor{Dark: "#c8f5c0", Light: "#2c6635"}
	theme.DiffHighlightRemovedColor = lipgloss.AdaptiveColor{Dark: "#ffa3c0", Light: "#a32a52"}
	theme.DiffAddedBgColor = lipgloss.AdaptiveColor{Dark: "#1d2b1c", Light: "#e6f5e0"}
	theme.DiffRemovedBgColor = lipgloss.AdaptiveColor{Dark: "#2e1a22", Light: "#fbe6ec"}
	theme.DiffContextBgColor = lipgloss.AdaptiveColor{Dark: darkBackground, Light: lightBackground}
	theme.DiffLineNumberColor = lipgloss.AdaptiveColor{Dark: "#473f66", Light: "#c5bca5"}
	theme.DiffAddedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#16331a", Light: "#d2ecd2"}
	theme.DiffRemovedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: "#33121f", Light: "#f5d2dd"}

	theme.MarkdownTextColor = lipgloss.AdaptiveColor{Dark: darkForeground, Light: lightForeground}
	theme.MarkdownHeadingColor = lipgloss.AdaptiveColor{Dark: darkPrimary, Light: lightPrimary}
	theme.MarkdownLinkColor = lipgloss.AdaptiveColor{Dark: darkAccent, Light: lightAccent}
	theme.MarkdownLinkTextColor = lipgloss.AdaptiveColor{Dark: darkCyan, Light: lightCyan}
	theme.MarkdownCodeColor = lipgloss.AdaptiveColor{Dark: darkGreen, Light: lightGreen}
	theme.MarkdownBlockQuoteColor = lipgloss.AdaptiveColor{Dark: darkYellow, Light: lightYellow}
	theme.MarkdownEmphColor = lipgloss.AdaptiveColor{Dark: darkYellow, Light: lightYellow}
	theme.MarkdownStrongColor = lipgloss.AdaptiveColor{Dark: darkPrimary, Light: lightPrimary}
	theme.MarkdownHorizontalRuleColor = lipgloss.AdaptiveColor{Dark: darkComment, Light: lightComment}
	theme.MarkdownListItemColor = lipgloss.AdaptiveColor{Dark: darkSecondary, Light: lightSecondary}
	theme.MarkdownListEnumerationColor = lipgloss.AdaptiveColor{Dark: darkCyan, Light: lightCyan}
	theme.MarkdownImageColor = lipgloss.AdaptiveColor{Dark: darkAccent, Light: lightAccent}
	theme.MarkdownImageTextColor = lipgloss.AdaptiveColor{Dark: darkCyan, Light: lightCyan}
	theme.MarkdownCodeBlockColor = lipgloss.AdaptiveColor{Dark: darkForeground, Light: lightForeground}

	theme.SyntaxCommentColor = lipgloss.AdaptiveColor{Dark: darkComment, Light: lightComment}
	theme.SyntaxKeywordColor = lipgloss.AdaptiveColor{Dark: darkSecondary, Light: lightSecondary}
	theme.SyntaxFunctionColor = lipgloss.AdaptiveColor{Dark: darkAccent, Light: lightAccent}
	theme.SyntaxVariableColor = lipgloss.AdaptiveColor{Dark: darkRed, Light: lightRed}
	theme.SyntaxStringColor = lipgloss.AdaptiveColor{Dark: darkGreen, Light: lightGreen}
	theme.SyntaxNumberColor = lipgloss.AdaptiveColor{Dark: darkOrange, Light: lightOrange}
	theme.SyntaxTypeColor = lipgloss.AdaptiveColor{Dark: darkYellow, Light: lightYellow}
	theme.SyntaxOperatorColor = lipgloss.AdaptiveColor{Dark: darkCyan, Light: lightCyan}
	theme.SyntaxPunctuationColor = lipgloss.AdaptiveColor{Dark: darkForeground, Light: lightForeground}

	return theme
}

func init() {
	RegisterTheme("fable", NewFableTheme())
}
