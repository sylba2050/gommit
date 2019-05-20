package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func main() {
	items := []string{
        "feat",
        "fix",
        "update",
        "style",
        "doc",
        "add",
        "delete",
        "refactor",
        "perf",
        "disable",
    }
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    "Select Prefix",
			Items:    items,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %s\n", result)
}
