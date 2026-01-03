package llm

import (
	"context"
	"fmt"
	"strings"

	"github.com/coretexos/coretex-incident-enricher/internal/types"
)

func Summarize(ctx context.Context, provider string, input types.EvidenceBundle, redactionLevel string) (types.Summary, error) {
	p := strings.ToLower(strings.TrimSpace(provider))
	switch p {
	case "", "mock":
		return SummarizeMock(input, redactionLevel), nil
	case "openai":
		return SummarizeOpenAI(ctx, input, redactionLevel)
	default:
		return types.Summary{}, fmt.Errorf("unsupported llm provider: %s", provider)
	}
}
