package provider

import (
	"testing"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/opencode-ai/opencode/internal/llm/models"
)

// TestPreparedMessagesOmitsTemperatureForNoTemperatureModels locks the
// contract that models flagged NoTemperature never send the temperature
// parameter — Anthropic rejects it with 400 invalid_request_error.
func TestPreparedMessagesOmitsTemperatureForNoTemperatureModels(t *testing.T) {
	msgs := []anthropic.MessageParam{
		anthropic.NewUserMessage(anthropic.NewTextBlock("hello")),
	}

	for id, model := range models.SupportedModels {
		if model.Provider != models.ProviderAnthropic {
			continue
		}
		client := &anthropicClient{
			providerOptions: providerClientOptions{
				model:         model,
				maxTokens:     1024,
				systemMessage: "test",
			},
		}
		params := client.preparedMessages(msgs, nil)

		if model.NoTemperature && params.Temperature.Valid() {
			t.Errorf("model %s is flagged NoTemperature but temperature was set", id)
		}
		if !model.NoTemperature && !params.Temperature.Valid() {
			t.Errorf("model %s should have temperature set but it was omitted", id)
		}
	}
}

// TestFableModelRegistered ensures the fable-class default model exists with
// the capability flag set.
func TestFableModelRegistered(t *testing.T) {
	m, ok := models.SupportedModels[models.ClaudeFable5]
	if !ok {
		t.Fatal("claude-fable-5 not registered in SupportedModels")
	}
	if !m.NoTemperature {
		t.Error("claude-fable-5 must have NoTemperature=true (API rejects the param)")
	}
	if m.CostPer1MIn == 0 || m.CostPer1MOut == 0 {
		t.Error("claude-fable-5 must have non-zero pricing for cost accounting")
	}
}
