package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// Ghost-text prediction: a small fast model (draft) predicts the rest of the
// user's message while they type; the user accepts with Tab or keeps typing
// (speculative decoding with the human as verifier).

const (
	ghostDebounce  = 500 * time.Millisecond
	ghostMinChars  = 6
	ghostMaxTokens = 40
)

type ghostTickMsg struct{ seq int }

type ghostResultMsg struct {
	input      string
	suggestion string
}

func ghostModel() string {
	if m := os.Getenv("OPENCODE_GHOST_MODEL"); m != "" {
		return m
	}
	return "claude-haiku-4-5"
}

func ghostEnabled() bool {
	return os.Getenv("OPENCODE_GHOST_DISABLE") == "" && os.Getenv("ANTHROPIC_API_KEY") != ""
}

// scheduleGhost returns a debounce tick carrying a sequence number; stale
// ticks (older seq) are dropped in the editor's Update.
func scheduleGhost(seq int) tea.Cmd {
	return tea.Tick(ghostDebounce, func(time.Time) tea.Msg {
		return ghostTickMsg{seq: seq}
	})
}

// fetchGhost calls the draft model for a continuation of the partial input.
func fetchGhost(input string) tea.Cmd {
	return func() tea.Msg {
		suggestion := requestGhostCompletion(input)
		return ghostResultMsg{input: input, suggestion: suggestion}
	}
}

func requestGhostCompletion(input string) string {
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		return ""
	}
	baseURL := os.Getenv("ANTHROPIC_BASE_URL")
	if baseURL == "" {
		baseURL = "https://api.anthropic.com"
	}

	body, err := json.Marshal(map[string]any{
		"model":      ghostModel(),
		"max_tokens": ghostMaxTokens,
		"system": "You autocomplete a user's partially-typed message to a coding agent. " +
			"Output ONLY the most likely continuation of their text - no quotes, no preamble, no repetition of what they typed. " +
			"Silently fix typos in your continuation. Keep it short (one phrase or sentence). " +
			"If the message already looks complete, output nothing.",
		"messages": []map[string]string{
			{"role": "user", "content": input},
		},
	})
	if err != nil {
		return ""
	}

	req, err := http.NewRequest("POST", strings.TrimRight(baseURL, "/")+"/v1/messages", bytes.NewReader(body))
	if err != nil {
		return ""
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return ""
	}

	var out struct {
		Content []struct {
			Text string `json:"text"`
		} `json:"content"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return ""
	}
	var sb strings.Builder
	for _, c := range out.Content {
		sb.WriteString(c.Text)
	}
	s := strings.TrimSpace(sb.String())
	s = strings.Trim(s, `"`)
	if strings.HasPrefix(strings.ToLower(s), strings.ToLower(strings.TrimSpace(input))) {
		s = strings.TrimSpace(s[len(strings.TrimSpace(input)):])
	}
	if s == "" {
		return ""
	}
	return fmt.Sprintf(" %s", strings.TrimLeft(s, " "))
}
