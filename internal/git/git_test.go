package git

import (
	"testing"
)

func TestBuildCommitMessage(t *testing.T) {
	tests := []struct {
		name          string
		branchName    string
		commitMessage string
		noPrefix      bool
		customPrefix  string
		expected      string
	}{
		{
			name:          "Default behavior",
			branchName:    "feature/login",
			commitMessage: "Add user authentication",
			noPrefix:      false,
			customPrefix:  "",
			expected:      "[feature/login] Add user authentication",
		},
		{
			name:          "With --no-prefix",
			branchName:    "feature/login",
			commitMessage: "Add user authentication",
			noPrefix:      true,
			customPrefix:  "",
			expected:      "Add user authentication",
		},
		{
			name:          "With --prefix",
			branchName:    "feature/login",
			commitMessage: "Add user authentication",
			noPrefix:      false,
			customPrefix:  "feat",
			expected:      "feat Add user authentication",
		},
		{
			name:          "With --prefix and --no-prefix (--no-prefix should take precedence)",
			branchName:    "feature/login",
			commitMessage: "Add user authentication",
			noPrefix:      true,
			customPrefix:  "feat",
			expected:      "Add user authentication",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BuildCommitMessage(tt.branchName, tt.commitMessage, tt.noPrefix, tt.customPrefix)
			if result != tt.expected {
				t.Errorf("BuildCommitMessage() = %v, want %v", result, tt.expected)
			}
		})
	}
}
