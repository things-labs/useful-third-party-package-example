package main

import (
	"log"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/storage/memory"
)

func main() {
	// Create the remote with repository URL
	rem := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{"https://github.com/golang/go"},
	})

	log.Print("Fetching tags...")

	// We can then use every Remote functions to retrieve wanted information
	refs, err := rem.List(&git.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// Filters the references list and only keeps tags
	var tags []string
	for _, ref := range refs {
		name := ref.Name()
		shortName := name.Short()
		if name.IsTag() && strings.HasPrefix(shortName, "go") && strings.Contains(shortName, "15") {
			tags = append(tags, shortName)
		}
	}

	if len(tags) == 0 {
		log.Println("No tags!")
		return
	}

	log.Printf("Tags found: %v", tags)
}
