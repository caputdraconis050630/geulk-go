package main

import "fmt"

func (cli *CLI) saveDiscussions(provider string, examCode string) {
	fmt.Printf("Saving the discussions of %s.%s to a file\n", provider, examCode)
}
