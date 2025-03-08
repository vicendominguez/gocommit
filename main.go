package main

import (
	"flag"
	"fmt"
	"os/exec"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/pterm/pterm"
	"golang.org/x/net/context"

	"github.com/ollama/ollama/api"
)

const version = "0.0.1"

func main() {

	noPrefix := flag.Bool("no-prefix", false, "Disable the prefix in the commit message")
	customPrefix := flag.String("prefix", "", "Define a custom prefix for the commit message")
	showVersion := flag.Bool("version", false, "Show the version of the application")
	showHelp := flag.Bool("help", false, "Show available flags")

	flag.Parse()

	// Mostrar la versión y salir
	if *showVersion {
		fmt.Printf("Version: %s\n", version)
		return
	}

	// Mostrar la ayuda y salir
	if *showHelp {
		flag.PrintDefaults()
		return
	}
	// Open the current repository
	repo, err := git.PlainOpen(".")
	if err != nil {
		pterm.Error.Printf("Error opening the repository: %v\n", err)
		return
	}

	// Get the current branch name
	branchName, err := getCurrentBranch(repo)
	if err != nil {
		pterm.Error.Printf("Error: %v\n", err)
		pterm.Warning.Println("Please create a branch first using 'git branch <branch-name>'.")
		return
	}

	// Get the diff of changes in the staging area
	diff, err := getGitDiffCached(repo)
	if err != nil {
		pterm.Error.Printf("Error getting the cached diff: %v\n", err)
		return
	}

	// Generate the commit message using Ollama
	commitMessage, err := generateCommitMessage(diff)
	if err != nil {
		pterm.Error.Printf("Error generating the commit message: %v\n", err)
		return
	}

	var prefixMessage string
	if *noPrefix {
		prefixMessage = ""
	} else if *customPrefix != "" {
		prefixMessage = fmt.Sprintf("[%s] ", *customPrefix)
	} else {
		prefixMessage = fmt.Sprintf("[%s] ", branchName)
	}

	// Añadir el prefijo al mensaje del commit
	fullCommitMessage := fmt.Sprintf("%s%s", prefixMessage, commitMessage)

	// Create the commit
	err = createGitCommit(repo, fullCommitMessage)
	if err != nil {
		pterm.Error.Printf("Error creating the commit: %v\n", err)
		return
	}

	pterm.Success.Printf("Commit created successfully: %s\n", fullCommitMessage)
}

// Get the current branch name
func getCurrentBranch(repo *git.Repository) (string, error) {
	ref, err := repo.Head()
	if err != nil {
		// Si el repositorio está vacío (no hay commits), devuelve un error
		if err == plumbing.ErrReferenceNotFound {
			return "", fmt.Errorf("no commits found (repository is empty)")
		}
		return "", err
	}

	// Verifica si HEAD apunta a una rama
	if !ref.Name().IsBranch() {
		return "", fmt.Errorf("HEAD is not pointing to a branch")
	}

	return ref.Name().Short(), nil
}

// I've tried it using go-git but I had all kind of problems.
// To have something working I've made this bullshit.
func getGitDiffCached(repo *git.Repository) (string, error) {
	cmd := exec.Command("git", "diff", "--cached")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to run git diff: %w", err)
	}
	return string(out), nil
}

func generateCommitMessage(diff string) (string, error) {

	client, err := api.ClientFromEnvironment()
	if err != nil {
		return "", fmt.Errorf("Error creating ollama client: %v", err)
	}

	prompt := "Analyze the following code diff and generate a summarized and concise commit message that describes the changes made. Include only the commit message itself, without any introduction or conclusion using not more than 20 words. Diff:\n" + diff

	messages := []api.Message{
		api.Message{
			Role:    "system",
			Content: "You are a expert code analyzer",
		},
		api.Message{
			Role:    "user",
			Content: prompt,
		},
	}

	ctx := context.Background()
	req := &api.ChatRequest{
		Model:    "llama3.1",
		Messages: messages,
		Stream:   new(bool),
	}

	var commitMessage string

	getcommitMessage := func(resp api.ChatResponse) error {
		commitMessage += resp.Message.Content
		return nil
	}

	err = client.Chat(ctx, req, getcommitMessage)

	if err != nil {
		pterm.Error.Println(err)
	}

	return commitMessage, nil
}

// Create the commit
func createGitCommit(repo *git.Repository, message string) error {
	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	// Add all changes to the staging area
	_, err = worktree.Add(".")
	if err != nil {
		return err
	}

	// Obtener la configuración global y local
	globalCfg, err := repo.ConfigScoped(config.GlobalScope)
	if err != nil {
		return fmt.Errorf("error reading global Git configuration: %v", err)
	}

	localCfg, err := repo.ConfigScoped(config.LocalScope)
	if err != nil {
		return fmt.Errorf("error reading local Git configuration: %v", err)
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
		return fmt.Errorf("user name or email not found in Git configuration")
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
		return err
	}

	// Print the hash of the created commit
	pterm.Info.Printf("Commit created with hash: %s\n", commit.String())
	return nil
}

