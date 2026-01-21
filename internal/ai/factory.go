package ai

import (
	"context"
	"fmt"
	"strings"

	"github.com/OpenListTeam/OpenList/v4/internal/conf"
	"github.com/OpenListTeam/OpenList/v4/internal/op"
)

func GetProvider(name string) Provider {
	endpoint := ""
	keys := ""
	model := ""

	if name == "Gemini" {
		item, _ := op.GetSettingItemByKey(conf.AiGeminiEndpoint)
		if item != nil {
			endpoint = item.Value
		}
		item, _ = op.GetSettingItemByKey(conf.AiGeminiKeys)
		if item != nil {
			keys = item.Value
		}
		return &GeminiProvider{
			Endpoint: endpoint,
			Keys:     parseKeys(keys),
		}
	} else if name == "Doubao" {
		item, _ := op.GetSettingItemByKey(conf.AiDoubaoEndpoint)
		if item != nil {
			endpoint = item.Value
		}
		item, _ = op.GetSettingItemByKey(conf.AiDoubaoKeys)
		if item != nil {
			keys = item.Value
		}
		item, _ = op.GetSettingItemByKey(conf.AiDoubaoModel)
		if item != nil {
			model = item.Value
		}
		return &DoubaoProvider{
			Endpoint: endpoint,
			Keys:     parseKeys(keys),
			Model:    model,
		}
	}
	return nil
}

func parseKeys(raw string) []string {
	parts := strings.Split(raw, "\n")
	var keys []string
	for _, k := range parts {
		k = strings.TrimSpace(k)
		if k != "" {
			keys = append(keys, k)
		}
	}
	return keys
}

// Cooperation Mode Logic
func ProcessWithCooperation(ctx context.Context, taskType string, input interface{}) Response {
	gemini := GetProvider("Gemini")
	doubao := GetProvider("Doubao")

	if gemini == nil || doubao == nil {
		return Response{Error: fmt.Errorf("both gemini and doubao must be configured for cooperation mode")}
	}

	switch taskType {
	case "translate":
		text := input.(string)
		// 1. Gemini understands and translates raw
		res := gemini.Translate(ctx, text, "Chinese")
		if res.Error != nil {
			return res
		}
		// 2. Doubao refines for natural Chinese
		refinePrompt := fmt.Sprintf("Please refine this Chinese translation to be more natural and professional:\n%s", res.Content)
		return doubao.Chat(ctx, Request{Prompt: refinePrompt})

	case "summary":
		content := input.(string)
		// 1. Gemini handles long context summary
		res := gemini.Summarize(ctx, content)
		if res.Error != nil {
			return res
		}
		// 2. Doubao optimizes the summary tone
		refinePrompt := fmt.Sprintf("Refine this summary to be more concise and easy to read:\n%s", res.Content)
		return doubao.Chat(ctx, Request{Prompt: refinePrompt})
	}

	return Response{Error: fmt.Errorf("unknown task type: %s", taskType)}
}
