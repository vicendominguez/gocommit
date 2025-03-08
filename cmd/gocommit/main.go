package main

import (
	"fmt"
	"log"

	"github.com/pterm/pterm"

	"gocommit/internal/cli"
	"gocommit/internal/git"
	"gocommit/internal/ollama"
	"gocommit/internal/version"
)

func main() {
	// Parsear flags
	opts, err := cli.ParseFlags()
	if err != nil {
		log.Fatalf("Error parsing flags: %v", err)
	}

	// Mostrar versión y salir
	if opts.ShowVersion {
		fmt.Printf("gocommit version: %s\n", version.Version)
		return
	}

	// Mostrar ayuda y salir
	if opts.ShowHelp {
		cli.PrintHelp()
		return
	}

	// Inicializar el repositorio Git
	repo, err := git.OpenRepository(".")
	if err != nil {
		log.Fatalf("Error opening repository: %v", err)
	}

	// Obtener la rama actual
	branchName, err := git.GetCurrentBranch(repo)
	if err != nil {
		log.Fatalf("Error getting current branch: %v", err)
	}

	// Obtener el diff en el staging area
	diff, err := git.GetStagedDiff(repo)
	if err != nil {
		log.Fatalf("Error getting staged diff: %v", err)
	}

	// Generar el mensaje del commit usando Ollama
	commitMessage, err := ollama.GenerateCommitMessage(diff)
	if err != nil {
		log.Fatalf("Error generating commit message: %v", err)
	}

	// Construir el mensaje final del commit
	fullCommitMessage := git.BuildCommitMessage(branchName, commitMessage, opts.NoPrefix, opts.CustomPrefix)

	// Crear el commit
	commitHash, err := git.CreateCommit(repo, fullCommitMessage)
	if err != nil {
		pterm.Error.Printf("Error creating commit: %v", err)
	} else {
		pterm.Success.Printf("Commit created successfully: %s\n", fullCommitMessage)
		pterm.Info.Printf("Commit created with hash: %s\n", commitHash)
	}
}
