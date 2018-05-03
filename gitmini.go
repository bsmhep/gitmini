package main

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/storer"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("error: %s\n", err)
	os.Exit(1)
}

func version() {
	fmt.Println("gittemp version 0.1.0")
}

func clone(args []string) {
	url := args[0]
	directory := args[1]

	_, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL: url,
	})
	CheckIfError(err)
}

func checkout(args []string) {
	path := args[0]
	ref_name := args[1]

	r, err := git.PlainOpen(path)
	CheckIfError(err)

	ref_full_name := fmt.Sprintf("refs/heads/%s", ref_name)
	_, err = r.Reference(plumbing.ReferenceName(ref_full_name), true)
	if err != nil {
		ref_full_name = fmt.Sprintf("refs/tags/%s", ref_name)
		_, err = r.Reference(plumbing.ReferenceName(ref_full_name), true)
		CheckIfError(err)
	}

	w, err := r.Worktree()
	CheckIfError(err)

	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(ref_full_name),
	})
	CheckIfError(err)
}

func ls_remote(args []string) {
	ref_type := args[0]
	url := args[1]

	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: url,
	})
	CheckIfError(err)

	var refIter storer.ReferenceIter

	switch ref_type {
	case "branches":
		refIter, err = r.Branches()
	case "tags":
		refIter, err = r.Tags()
	default:
		fmt.Printf("Unknown ref type: %s\n", ref_type)
		os.Exit(1)
	}
	CheckIfError(err)

	err = refIter.ForEach(func(c *plumbing.Reference) error {
		fmt.Println(c)
		return nil
	})
	CheckIfError(err)
}

func ls(args []string) {
	ref_type := args[0]
	path := args[1]

	r, err := git.PlainOpen(path)
	CheckIfError(err)

	var refIter storer.ReferenceIter

	switch ref_type {
	case "branches":
		refIter, err = r.Branches()
	case "tags":
		refIter, err = r.Tags()
	default:
		fmt.Printf("Unknown ref type: %s\n", ref_type)
		os.Exit(1)
	}
	CheckIfError(err)

	err = refIter.ForEach(func(c *plumbing.Reference) error {
		fmt.Println(c.Name().Short())
		return nil
	})
	CheckIfError(err)
}

func main() {
	var subCommand string
	if len(os.Args) == 1 {
		subCommand = "version"
	} else {
		subCommand = os.Args[1]
	}

	switch subCommand {
	case "version":
		version()
	case "clone":
		clone(os.Args[2:])
	case "checkout":
		checkout(os.Args[2:])
	case "ls-remote":
		ls_remote(os.Args[2:])
	case "ls":
		ls(os.Args[2:])
	default:
		fmt.Printf("Unknown command: %s\n", subCommand)
		os.Exit(1)
	}
}
