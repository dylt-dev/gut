package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main () {
	var err error
	var args []string = os.Args[1:]
	var cmdName string = args[0]
	switch cmdName {
	case "cat-file",
	     "hash-object":
		fmt.Println("")
		fmt.Printf("Recognized command: %s (%v)\n", cmdName, args)
		fmt.Println("We gut this!")
		fmt.Println("")
		var cmd *exec.Cmd = exec.Command("git", args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(-1)
		}
	case "clone",
	     "init",
	     "add",
	     "mv",
	     "restore",
	     "rm",
	     "bisect",
	     "diff",
	     "grep",
	     "log",
	     "show",
	     "status",
	     "branch",
	     "commit",
	     "merge",
	     "rebase",
	     "reset",
	     "switch",
	     "tag",
	     "fetch",
	     "pull",
	     "push":
		fmt.Println("")
		fmt.Println("Gut news! You can use simple commands to do what you need, and become invincible!")
		fmt.Println("Go to https://git-scm.com/book/en/v2/Git-Internals-Plumbing-and-Porcelain to learn more.")
		fmt.Println("")
		fmt.Println("Have a gut time!")
		fmt.Println("")
	case "cat": 
		cmd := exec.Command(cmdName)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		fmt.Printf("%v\n", err)
	case "ls": 
		cmd := exec.Command("ls", "-lah")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		fmt.Printf("%v\n", err)
	default: 
		fmt.Println("")
		fmt.Println("Oh my gutness ... we don't recognize that command! Sorry!")
		fmt.Println("")
	}
}
