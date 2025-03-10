package git

import (
	"fmt"
	"os/exec"
	"time"

	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func OpenRepository(path string) (*git.Repository, error) {
	// Convert the path to an absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %w", err)
	}

	// Start searching from the given path and move up the directory tree
	for {
		// Try to open the repository at the current path
		repo, err := git.PlainOpen(absPath)
		if err == nil {
			return repo, nil
		}

		// If we reach the root directory and still haven't found the .git directory, return an error
		if absPath == filepath.Dir(absPath) {
			return nil, fmt.Errorf("no Git repository found in %s or any parent directory", path)
		}

		// Move up to the parent directory
		absPath = filepath.Dir(absPath)
	}
}

// GetCurrentBranch retrieves the name of the current branch.
// It returns an error if the repository is empty or HEAD is not pointing to a branch.
func GetCurrentBranch(repo *git.Repository) (string, error) {
	ref, err := repo.Head()
	if err != nil {
		if err == plumbing.ErrReferenceNotFound {
			return "", fmt.Errorf("no commits found (repository is empty)")
		}
		return "", fmt.Errorf("failed to get HEAD: %w", err)
	}

	if !ref.Name().IsBranch() {
		return "", fmt.Errorf("HEAD is not pointing to a branch")
	}

	return ref.Name().Short(), nil
}

// GetStagedDiff retrieves the diff of changes in the staging area.
// It uses the `git diff --cached` command to get the diff.
func GetStagedDiff(repo *git.Repository) (string, error) {
	cmd := exec.Command("git", "diff", "--cached")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to run git diff: %w", err)
	}
	return string(out), nil
}

// BuildCommitMessage constructs the commit message based on the provided parameters.
// - If noPrefix is true, it returns only the commit message.
// - If customPrefix is provided, it uses it directly (overriding the brackets).
// - By default, it uses the branch name inside brackets.
func BuildCommitMessage(branchName, commitMessage string, noPrefix bool, customPrefix string) string {
	// If --no-prefix is enabled, return only the commit message
	if noPrefix {
		return commitMessage
	}

	// If --prefix is provided, use it directly (overrides the brackets)
	if customPrefix != "" {
		return fmt.Sprintf("%s %s", customPrefix, commitMessage)
	}

	// Default behavior: use the branch name inside brackets
	return fmt.Sprintf("[%s] %s", branchName, commitMessage)
}

// CreateCommit creates a new commit with the specified message.
// It stages all changes, retrieves the user's Git configuration, and creates the commit.
// It returns the commit hash or an error if the commit cannot be created.
func CreateCommit(repo *git.Repository, message string) (string, error) {
	worktree, err := repo.Worktree()
	if err != nil {
		return "", fmt.Errorf("failed to get worktree: %w", err)
	}

	// Add all changes to the staging area
	_, err = worktree.Add(".")
	if err != nil {
		return "", fmt.Errorf("failed to add changes: %w", err)
	}

	// Obtener la configuración global y local
	globalCfg, err := repo.ConfigScoped(config.GlobalScope)
	if err != nil {
		return "", fmt.Errorf("error reading global Git configuration: %v", err)
	}

	localCfg, err := repo.ConfigScoped(config.LocalScope)
	if err != nil {
		return "", fmt.Errorf("error reading local Git configuration: %v", err)
	}

	// Combinar las configuraciones: priorizar la configuración local sobre la global
	userName := localCfg.User.Name
	if userName == "" {
		userName = globalCfg.User.Name
	}

	userEmail := localCfg.User.Email
	if userEmail == "" {
		userEmail = globalCfg.User.Email
	}

	if userName == "" || userEmail == "" {
		return "", fmt.Errorf("user name or email not found in Git configuration")
	}

	now := time.Now()

	// Create the commit
	commit, err := worktree.Commit(message, &git.CommitOptions{
		Author: &object.Signature{
			Name:  userName,
			Email: userEmail,
			When:  now,
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to create commit: %w", err)
	}

	return commit.String(), nil
}
