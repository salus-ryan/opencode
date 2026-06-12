package version

// Version is set at build time via -ldflags or defaults to the value below.
// Example: go build -ldflags "-X github.com/opencode-ai/opencode/internal/version.Version=fable-0.2.0"
var Version = "fable-0.1.0"
