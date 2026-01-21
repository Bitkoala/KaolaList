package ai

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type GeminiProvider struct {
	Endpoint string
	Keys     []string
	Model    string
	mu       sync.Mutex
	index    int
}

func (p *GeminiProvider) getNextKey() string {
	p.mu.Lock()
	defer p.mu.Unlock()
	if len(p.Keys) == 0 {
		return ""
	}
	key := p.Keys[p.index]
	p.index = (p.index + 1) % len(p.Keys)
	return key
}

func (p *GeminiProvider) Chat(ctx context.Context, req Request) Response {
	key := p.getNextKey()
	if key == "" {
		return Response{Error: fmt.Errorf("no gemini api keys configured")}
	}

	model := req.Model
	if model == "" {
		model = p.Model
	}
	if model == "" {
		model = "gemini-1.5-flash"
	}

	// Build Gemini request body
	type Part struct {
		Text       string `json:"text,omitempty"`
		InlineData *struct {
			MimeType string `json:"mimeType"`
			Data     string `json:"data"`
		} `json:"inlineData,omitempty"`
	}
	type Content struct {
		Parts []Part `json:"parts"`
	}
	type Body struct {
		Contents []Content `json:"contents"`
	}

	parts := []Part{{Text: req.Prompt}}
	if len(req.Image) > 0 {
		parts = append(parts, Part{
			InlineData: &struct {
				MimeType string `json:"mimeType"`
				Data     string `json:"data"`
			}{
				MimeType: "image/jpeg",
				Data:     base64.StdEncoding.EncodeToString(req.Image),
			},
		})
	}

	body := Body{
		Contents: []Content{{Parts: parts}},
	}

	jsonBody, _ := json.Marshal(body)
	url := fmt.Sprintf("%s/v1beta/models/%s:generateContent?key=%s", p.Endpoint, model, key)

	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return Response{Error: err}
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return Response{Error: err}
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return Response{Error: fmt.Errorf("gemini api failed (%d): %s", resp.StatusCode, string(respBody))}
	}

	var result struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
	}

	if err := json.Unmarshal(respBody, &result); err != nil {
		return Response{Error: err}
	}

	if len(result.Candidates) == 0 || len(result.Candidates[0].Content.Parts) == 0 {
		return Response{Error: fmt.Errorf("empty response from gemini")}
	}

	return Response{Content: result.Candidates[0].Content.Parts[0].Text}
}

func (p *GeminiProvider) Translate(ctx context.Context, text string, targetLang string) Response {
	prompt := fmt.Sprintf("Translate to %s, only output the result:\n%s", targetLang, text)
	return p.Chat(ctx, Request{Prompt: prompt})
}

func (p *GeminiProvider) Summarize(ctx context.Context, content string) Response {
	prompt := fmt.Sprintf("Analyze and summarize this in professional Chinese:\n%s", content)
	return p.Chat(ctx, Request{Prompt: prompt})
}
