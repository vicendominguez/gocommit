package ollama

import (
	"context"
	"fmt"

	"github.com/ollama/ollama/api"
	"github.com/pterm/pterm"
)

// GenerateCommitMessage genera un mensaje de commit usando Ollama.
func GenerateCommitMessage(diff string) (string, error) {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		return "", fmt.Errorf("failed to create Ollama client: %w", err)
	}

	prompt := "Analyze the following code diff and generate a summarized and concise commit message that describes the changes made. Include only the commit message itself, without any introduction or conclusion using not more than 20 words. Diff:\n" + diff

	messages := []api.Message{
		{Role: "system", Content: "You are an expert code analyzer."},
		{Role: "user", Content: prompt},
	}

	ctx := context.Background()
	req := &api.ChatRequest{
		Model:    "llama3.1",
		Messages: messages,
		Stream:   new(bool),
	}

	var commitMessage string

	err = client.Chat(ctx, req, func(resp api.ChatResponse) error {
		commitMessage += resp.Message.Content
		return nil
	})

	if err != nil {
		return "", fmt.Errorf("failed to generate commit message: %w", err)
	}
	pterm.Warning.Println(commitMessage)
	return commitMessage, nil
}
