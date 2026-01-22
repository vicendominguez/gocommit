package ollama

import (
	"context"
	"fmt"
	"strings"

	"github.com/ollama/ollama/api"
)

// GenerateCommitMessage genera un mensaje de commit usando Ollama.
func GenerateCommitMessage(diff string, model string) (string, error) {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		return "", fmt.Errorf("failed to create Ollama client: %w", err)
	}

	prompt := "Generate a max 10 words commit message from this diff: " + diff

	messages := []api.Message{
		{Role: "system", Content: "You are an expert code analyzer."},
		{Role: "user", Content: prompt},
	}

	ctx := context.Background()
	req := &api.ChatRequest{
		Model:    model,
		Messages: messages,
		Stream:   new(bool),
	}

	var commitMessage string

	err = client.Chat(ctx, req, func(resp api.ChatResponse) error {
		commitMessage += resp.Message.Content
		return nil
	})

	if err != nil {
		if strings.Contains(err.Error(), "model") || strings.Contains(err.Error(), "not found") {
			return "", fmt.Errorf("model '%s' not found. Try: ollama pull %s", model, model)
		}
		return "", fmt.Errorf("failed to generate commit message: %w", err)
	}

	//cleaning up the ollama response
	commitMessage = strings.Trim(commitMessage, `"`)
	return commitMessage, nil
}
