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

type DoubaoProvider struct {
	Endpoint string
	Keys     []string
	Model    string
	mu       sync.Mutex
	index    int
}

func (p *DoubaoProvider) getNextKey() string {
	p.mu.Lock()
	defer p.mu.Unlock()
	if len(p.Keys) == 0 {
		return ""
	}
	key := p.Keys[p.index]
	p.index = (p.index + 1) % len(p.Keys)
	return key
}

func (p *DoubaoProvider) Chat(ctx context.Context, req Request) Response {
	key := p.getNextKey()
	if key == "" {
		return Response{Error: fmt.Errorf("no doubao api keys configured")}
	}

	model := req.Model
	if model == "" {
		model = p.Model
	}

	messages := []map[string]interface{}{
		{"role": "user", "content": []interface{}{}},
	}

	content := messages[0]["content"].([]interface{})
	if len(req.Image) > 0 {
		base64Img := base64.StdEncoding.EncodeToString(req.Image)
		content = append(content, map[string]interface{}{
			"type": "image_url",
			"image_url": map[string]string{
				"url": fmt.Sprintf("data:image/jpeg;base64,%s", base64Img),
			},
		})
	}
	content = append(content, map[string]interface{}{
		"type": "text",
		"text": req.Prompt,
	})
	messages[0]["content"] = content

	body := map[string]interface{}{
		"model":    model,
		"messages": messages,
	}

	jsonBody, _ := json.Marshal(body)
	httpReq, err := http.NewRequestWithContext(ctx, "POST", p.Endpoint+"/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		return Response{Error: err}
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+key)

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return Response{Error: err}
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return Response{Error: fmt.Errorf("doubao api failed (%d): %s", resp.StatusCode, string(respBody))}
	}

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.Unmarshal(respBody, &result); err != nil {
		return Response{Error: err}
	}

	if len(result.Choices) == 0 {
		return Response{Error: fmt.Errorf("empty response from doubao")}
	}

	return Response{Content: result.Choices[0].Message.Content}
}

func (p *DoubaoProvider) Translate(ctx context.Context, text string, targetLang string) Response {
	prompt := fmt.Sprintf("Translate the following text to %s, only output the translation, no explanation:\n%s", targetLang, text)
	return p.Chat(ctx, Request{Prompt: prompt})
}

func (p *DoubaoProvider) Summarize(ctx context.Context, content string) Response {
	prompt := fmt.Sprintf("Summarize the following content in professional Chinese, keep it concise:\n%s", content)
	return p.Chat(ctx, Request{Prompt: prompt})
}
